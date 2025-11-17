package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"quantumswap-cli/contracts/core"
	"strconv"
	"strings"
	"time"

	"github.com/quantumcoinproject/quantum-coin-go/accounts/abi/bind"
	"github.com/quantumcoinproject/quantum-coin-go/accounts/keystore"
	"github.com/quantumcoinproject/quantum-coin-go/common"
	"github.com/quantumcoinproject/quantum-coin-go/console/prompt"
	"github.com/quantumcoinproject/quantum-coin-go/core/types"
	"github.com/quantumcoinproject/quantum-coin-go/crypto/cryptobase"
	"github.com/quantumcoinproject/quantum-coin-go/crypto/signaturealgorithm"
	"github.com/quantumcoinproject/quantum-coin-go/ethclient"
)

const GAS_LIMIT_ENV = "GAS_LIMIT"
const CHAIN_ID_ENV = "CHAIN_ID"
const DEFAULT_CHAIN_ID = 123123

func getChainId() (int64, error) {
	chainIdEnv := os.Getenv(CHAIN_ID_ENV)
	if len(chainIdEnv) > 0 {
		chainId, err := strconv.ParseUint(chainIdEnv, 10, 64)
		if err != nil {
			fmt.Println("Error parsing chain id, err")
			return int64(chainId), err
		}
		fmt.Println("Using CHAIN_ID passed using environment variable", chainId)
		return int64(chainId), nil
	} else {
		return DEFAULT_CHAIN_ID, nil
	}
}

func getGasLimit(defaultLimit uint64) (uint64, error) {
	gasLimitEnv := os.Getenv(GAS_LIMIT_ENV)
	if len(gasLimitEnv) > 0 {
		gasLimit, err := strconv.ParseUint(gasLimitEnv, 10, 64)
		if err != nil {
			fmt.Println("Error parsing gas limit", err)
			return gasLimit, err
		}
		fmt.Println("Using gas limit passed using environment variable", gasLimit)
		return gasLimit, nil
	} else {
		return defaultLimit, nil
	}
}

func ReadDataFile(filename string) ([]byte, error) {
	// Open our jsonFile
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	fmt.Println("Successfully Opened ", filename)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue, nil
}

func findKeyFile(keyAddress string) (string, error) {
	keyfile := os.Getenv("DP_KEY_FILE")
	if len(keyfile) > 0 {
		return keyfile, nil
	}

	keyfileDir := os.Getenv("DP_KEY_FILE_DIR")
	if len(keyfileDir) == 0 {
		return "", errors.New("Both DP_KEY_FILE and DP_KEY_FILE_DIR environment variables not set")
	}

	files, err := ioutil.ReadDir(keyfileDir)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	addr := strings.ToLower(strings.Replace(keyAddress, "0x", "", 1))
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.Contains(strings.ToLower(file.Name()), addr) {
			return filepath.Join(keyfileDir, file.Name()), nil
		}
	}

	return "", errors.New("could not find key file")
}

func GetKeyFromFile(keyFile string, accPwd string) (*signaturealgorithm.PrivateKey, error) {
	secretKey, err := ReadDataFile(keyFile)
	if err != nil {
		return nil, err
	}

	password := accPwd
	key, err := keystore.DecryptKey(secretKey, password)
	if err != nil {
		return nil, err
	}

	return key.PrivateKey, nil
}

func GetKey(address string) (*signaturealgorithm.PrivateKey, error) {
	keyFile, err := findKeyFile(address)
	if err != nil {
		return nil, err
	}

	fmt.Println("keyFile", keyFile)
	secretKey, err := ReadDataFile(keyFile)
	if err != nil {
		return nil, err
	}
	password := os.Getenv("DP_ACC_PWD")
	if len(password) == 0 {
		password, err = prompt.Stdin.PromptPassword(fmt.Sprintf("Enter the wallet password : "))
		if err != nil {
			return nil, err
		}
		if len(password) == 0 {
			return nil, errors.New("password is not correct")
		}
	}
	key, err := keystore.DecryptKey(secretKey, password)
	if err != nil {
		return nil, err
	}

	if key.Address.IsEqualTo(common.HexToAddress(address)) == false {
		return nil, errors.New("address mismatch in key file")
	}

	return key.PrivateKey, nil
}

func deployCore(from string) error {
	key, err := GetKey(from)
	if err != nil {
		return err
	}

	client, err := ethclient.Dial(rawURL)
	if err != nil {
		return err
	}

	fromAddress, err := cryptobase.SigAlg.PublicKeyToAddress(&key.PublicKey)

	if err != nil {
		return err
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	chainId, err := getChainId()
	if err != nil {
		return err
	}
	txnOpts, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(chainId))

	if err != nil {
		return err
	}

	txnOpts.From = fromAddress
	txnOpts.Nonce = big.NewInt(int64(nonce))
	txnOpts.GasLimit, err = getGasLimit(uint64(6000000))
	if err != nil {
		return err
	}

	txnOpts.Value = big.NewInt(0)

	var tx *types.Transaction
	contractAddress, tx, _, err := core.DeployCore(txnOpts, client)
	if err != nil {
		return err
	}

	fmt.Println("Your request to deploy core contracts has been added to the queue for processing. Please check your account after 10 minutes.")
	fmt.Println("The transaction hash for tracking this request is: ", tx.Hash(), "contractAddress", contractAddress)
	fmt.Println()

	time.Sleep(1000 * time.Millisecond)

	return nil
}

func enableFee(from string, contractAddr string, fee int64, tick int64) error {
	key, err := GetKey(from)
	if err != nil {
		return err
	}

	client, err := ethclient.Dial(rawURL)
	if err != nil {
		return err
	}

	fromAddress, err := cryptobase.SigAlg.PublicKeyToAddress(&key.PublicKey)

	if err != nil {
		return err
	}

	contractAddress := common.HexToAddress(contractAddr)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	chainId, err := getChainId()
	if err != nil {
		return err
	}
	txnOpts, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(chainId))

	if err != nil {
		return err
	}

	txnOpts.From = fromAddress
	txnOpts.Nonce = big.NewInt(int64(nonce))
	txnOpts.GasLimit, err = getGasLimit(uint64(6000000))
	if err != nil {
		return err
	}

	txnOpts.Value = big.NewInt(0)

	contract, err := core.NewCore(contractAddress, client)
	if err != nil {
		return err
	}

	var tx *types.Transaction
	tx, err = contract.EnableFeeAmount(txnOpts, big.NewInt(fee), big.NewInt(tick))
	if err != nil {
		return err
	}

	fmt.Println("Your request to enable fee has been added to the queue for processing. Please check your account after 10 minutes.")
	fmt.Println("The transaction hash for tracking this request is: ", tx.Hash(), "contractAddress", contractAddress)
	fmt.Println()

	time.Sleep(1000 * time.Millisecond)

	return nil
}

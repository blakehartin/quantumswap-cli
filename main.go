package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/quantumcoinproject/quantum-coin-go/common"
	"github.com/quantumcoinproject/quantum-coin-go/console/prompt"
)

// 1) git clone https://github.com/quantumswapdex/quantumswap-core.git
// 2) solc --bin --bin-runtime --abi C:\github\quantumswap\main\quantumswap-core\contracts\UniswapV3Factory.sol  -o C:\github\quantumswap\quantumswap-cli\contracts\core --optimize
// 3) abigen --bin=C:\github\quantumswap\quantumswap-cli\contracts\core\UniswapV3Factory.bin --abi=C:\github\quantumswap\quantumswap-cli\contracts\core\UniswapV3Factory.abi --pkg=core --out=C:\github\quantumswap\quantumswap-cli\contracts\core\UniswapV3Factory.go

var rawURL string

func printHelp() {
	fmt.Println("--------")
	fmt.Println(" Usage")
	fmt.Println("--------")
	fmt.Println("quantumswap-cli deploycore FROM_ADDRESS")
	fmt.Println("      Set the following environment variables:")
	fmt.Println("           DP_RAW_URL, DP_KEY_FILE_DIR or DP_KEY_FILE")
}

func main() {
	fmt.Println("===============")
	fmt.Println(" QuantumSwap")
	fmt.Println("===============")

	if len(os.Args) < 2 {
		printHelp()
		return
	}

	rawURL = os.Getenv("DP_RAW_URL")
	if len(rawURL) == 0 {
		runtimeOS := strings.ToLower(runtime.GOOS)
		if runtimeOS == "windows" {
			rawURL = "\\\\.\\pipe\\geth.ipc"
		} else {
			rawURL = "data/geth.ipc"
		}
	}

	if os.Args[1] == "balance" {
		DeployCore()
	} else {
		printHelp()
	}

}

func DeployCore() {
	from := os.Args[2]
	if common.IsHexAddress(from) == false {
		fmt.Println("Invalid address", from)
		return
	}

	ethConfirm, err := prompt.Stdin.PromptConfirm(fmt.Sprintf("Do you want to deploy the core contract from %s?", from))
	if err != nil {
		fmt.Println("error", err)
		return
	}
	if ethConfirm != true {
		fmt.Println("confirmation not made")
		return
	}

	err = deployCore(from)
	if err != nil {
		fmt.Println("error", err)
		return
	}
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pairv2

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/quantumcoinproject/quantum-coin-go"
	"github.com/quantumcoinproject/quantum-coin-go/accounts/abi"
	"github.com/quantumcoinproject/quantum-coin-go/accounts/abi/bind"
	"github.com/quantumcoinproject/quantum-coin-go/common"
	"github.com/quantumcoinproject/quantum-coin-go/core/types"
	"github.com/quantumcoinproject/quantum-coin-go/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Pairv2MetaData contains all meta data concerning the Pairv2 contract.
var Pairv2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"_blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526001600a5534801561001557600080fd5b50600046905050336003819055506128a7806100326000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c80636a627842116100c3578063ba9a7a561161007c578063ba9a7a56146105f1578063bc25cf771461060f578063c45a01551461063d578063d21220a71461065b578063dd62ed3e14610679578063fff6cae9146106c55761014d565b80636a6278421461043557806370a08231146104775780637464fc3d146104b957806389afcb44146104d757806395d89b4114610520578063a9059cbb146105a35761014d565b806318160ddd1161011557806318160ddd1461032a57806323b872dd14610348578063313ce567146103a0578063485cc955146103c15780635909c0d5146103f95780635a3d5493146104175761014d565b8063022c0d9f1461015257806306fdde03146101e95780630902f1ac1461026c578063095ea7b3146102be5780630dfe16811461030c575b600080fd5b6101e76004803603608081101561016857600080fd5b81019080803590602001909291908035906020019092919080359060200190929190803590602001906401000000008111156101a357600080fd5b8201836020820111156101b557600080fd5b803590602001918460018302840111640100000000831117156101d757600080fd5b90919293919293905050506106cf565b005b6101f1610d32565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610231578082015181840152602081019050610216565b50505050905090810190601f16801561025e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610274610d6b565b60405180846dffffffffffffffffffffffffffff168152602001836dffffffffffffffffffffffffffff1681526020018263ffffffff168152602001935050505060405180910390f35b6102f4600480360360408110156102d457600080fd5b810190808035906020019092919080359060200190929190505050610dc8565b60405180821515815260200191505060405180910390f35b610314610ddf565b6040518082815260200191505060405180910390f35b610332610de5565b6040518082815260200191505060405180910390f35b6103886004803603606081101561035e57600080fd5b81019080803590602001909291908035906020019092919080359060200190929190505050610deb565b60405180821515815260200191505060405180910390f35b6103a8610eae565b604051808260ff16815260200191505060405180910390f35b6103f7600480360360408110156103d757600080fd5b810190808035906020019092919080359060200190929190505050610eb3565b005b610401610f3c565b6040518082815260200191505060405180910390f35b61041f610f42565b6040518082815260200191505060405180910390f35b6104616004803603602081101561044b57600080fd5b8101908080359060200190929190505050610f48565b6040518082815260200191505060405180910390f35b6104a36004803603602081101561048d57600080fd5b810190808035906020019092919050505061131f565b6040518082815260200191505060405180910390f35b6104c1611337565b6040518082815260200191505060405180910390f35b610503600480360360208110156104ed57600080fd5b810190808035906020019092919050505061133d565b604051808381526020018281526020019250505060405180910390f35b61052861178a565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561056857808201518184015260208101905061054d565b50505050905090810190601f1680156105955780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6105d9600480360360408110156105b957600080fd5b8101908080359060200190929190803590602001909291905050506117c3565b60405180821515815260200191505060405180910390f35b6105f96117da565b6040518082815260200191505060405180910390f35b61063b6004803603602081101561062557600080fd5b81019080803590602001909291905050506117e0565b005b6106456119f7565b6040518082815260200191505060405180910390f35b6106636119fd565b6040518082815260200191505060405180910390f35b6106af6004803603604081101561068f57600080fd5b810190808035906020019092919080359060200190929190505050611a03565b6040518082815260200191505060405180910390f35b6106cd611a28565b005b6001600a5414610747576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f556e697377617056323a204c4f434b454400000000000000000000000000000081525060200191505060405180910390fd5b6000600a81905550600085118061075e5750600084115b6107b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806127b86025913960400191505060405180910390fd5b6000806107be610d6b565b5091509150816dffffffffffffffffffffffffffff16871080156107f15750806dffffffffffffffffffffffffffff1686105b610846576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806128016021913960400191505060405180910390fd5b60008060006004549050600060055490508189141580156108675750808914155b6108d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f556e697377617056323a20494e56414c49445f544f000000000000000000000081525060200191505060405180910390fd5b60008b11156108ee576108ed828a8d611be4565b5b60008a111561090357610902818a8c611be4565b5b60008888905011156109a957886310d1e85c338d8d8c8c6040518663ffffffff1660e01b815260040180868152602001858152602001848152602001806020018281038252848482818152602001925080828437600081840152601f19601f8201169050808301925050509650505050505050600060405180830381600087803b15801561099057600080fd5b505af11580156109a4573d6000803e3d6000fd5b505050505b816370a08231306040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156109e457600080fd5b505afa1580156109f8573d6000803e3d6000fd5b505050506040513d6020811015610a0e57600080fd5b81019080805190602001909291905050509350806370a08231306040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610a5c57600080fd5b505afa158015610a70573d6000803e3d6000fd5b505050506040513d6020811015610a8657600080fd5b810190808051906020019092919050505092505050600089856dffffffffffffffffffffffffffff16038311610abd576000610ad3565b89856dffffffffffffffffffffffffffff160383035b9050600089856dffffffffffffffffffffffffffff16038311610af7576000610b0d565b89856dffffffffffffffffffffffffffff160383035b90506000821180610b1e5750600081115b610b73576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806127dd6024913960400191505060405180910390fd5b6000610baf610b8c600385611dee90919063ffffffff16565b610ba16103e888611dee90919063ffffffff16565b611e8390919063ffffffff16565b90506000610bed610bca600385611dee90919063ffffffff16565b610bdf6103e888611dee90919063ffffffff16565b611e8390919063ffffffff16565b9050610c37620f4240610c29896dffffffffffffffffffffffffffff168b6dffffffffffffffffffffffffffff16611dee90919063ffffffff16565b611dee90919063ffffffff16565b610c4a8284611dee90919063ffffffff16565b1015610cbe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f556e697377617056323a204b000000000000000000000000000000000000000081525060200191505060405180910390fd5b5050610ccc84848888611f06565b88337fd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d82284848f8f6040518085815260200184815260200183815260200182815260200194505050505060405180910390a35050505050506001600a819055505050505050565b6040518060400160405280600a81526020017f556e69737761702056320000000000000000000000000000000000000000000081525081565b6000806000600660009054906101000a90046dffffffffffffffffffffffffffff1692506006600e9054906101000a90046dffffffffffffffffffffffffffff1691506006601c9054906101000a900463ffffffff169050909192565b6000610dd5338484612264565b6001905092915050565b60045481565b60005481565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6002600086815260200190815260200160002060003381526020019081526020016000205414610e9857610e6f8260026000878152602001908152602001600020600033815260200190815260200160002054611e8390919063ffffffff16565b600260008681526020019081526020016000206000338152602001908152602001600020819055505b610ea38484846122cb565b600190509392505050565b601281565b6003543314610f2a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f556e697377617056323a20464f5242494444454e00000000000000000000000081525060200191505060405180910390fd5b81600481905550806005819055505050565b60075481565b60085481565b60006001600a5414610fc2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f556e697377617056323a204c4f434b454400000000000000000000000000000081525060200191505060405180910390fd5b6000600a81905550600080610fd5610d6b565b509150915060006004546370a08231306040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561101957600080fd5b505afa15801561102d573d6000803e3d6000fd5b505050506040513d602081101561104357600080fd5b8101908080519060200190929190505050905060006005546370a08231306040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561109557600080fd5b505afa1580156110a9573d6000803e3d6000fd5b505050506040513d60208110156110bf57600080fd5b8101908080519060200190929190505050905060006110f7856dffffffffffffffffffffffffffff1684611e8390919063ffffffff16565b9050600061111e856dffffffffffffffffffffffffffff1684611e8390919063ffffffff16565b9050600061112c8787612383565b905060008054905060008114156111805761116c6103e861115e6111598688611dee90919063ffffffff16565b612502565b611e8390919063ffffffff16565b985061117b60006103e8612564565b6111e3565b6111e0886dffffffffffffffffffffffffffff166111a78387611dee90919063ffffffff16565b816111ae57fe5b04886dffffffffffffffffffffffffffff166111d38487611dee90919063ffffffff16565b816111da57fe5b046125fa565b98505b6000891161123c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602881526020018061284a6028913960400191505060405180910390fd5b6112468a8a612564565b61125286868a8a611f06565b81156112ca576112c36006600e9054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff16600660009054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff16611dee90919063ffffffff16565b6009819055505b337f4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f8585604051808381526020018281526020019250505060405180910390a250505050505050506001600a81905550919050565b60016020528060005260406000206000915090505481565b60095481565b6000806001600a54146113b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f556e697377617056323a204c4f434b454400000000000000000000000000000081525060200191505060405180910390fd5b6000600a819055506000806113cb610d6b565b509150915060006004549050600060055490506000826370a08231306040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561141b57600080fd5b505afa15801561142f573d6000803e3d6000fd5b505050506040513d602081101561144557600080fd5b810190808051906020019092919050505090506000826370a08231306040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561149557600080fd5b505afa1580156114a9573d6000803e3d6000fd5b505050506040513d60208110156114bf57600080fd5b8101908080519060200190929190505050905060006001600030815260200190815260200160002054905060006114f68888612383565b9050600080549050806115128685611dee90919063ffffffff16565b8161151957fe5b049a50806115308585611dee90919063ffffffff16565b8161153757fe5b04995060008b11801561154a575060008a115b61159f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806128226028913960400191505060405180910390fd5b6115a93084612613565b6115b4878d8d611be4565b6115bf868d8c611be4565b866370a08231306040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156115fa57600080fd5b505afa15801561160e573d6000803e3d6000fd5b505050506040513d602081101561162457600080fd5b81019080805190602001909291905050509450856370a08231306040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561167257600080fd5b505afa158015611686573d6000803e3d6000fd5b505050506040513d602081101561169c57600080fd5b810190808051906020019092919050505093506116bb85858b8b611f06565b81156117335761172c6006600e9054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff16600660009054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff16611dee90919063ffffffff16565b6009819055505b8b337fdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d819364968d8d604051808381526020018281526020019250505060405180910390a35050505050505050506001600a81905550915091565b6040518060400160405280600681526020017f554e492d5632000000000000000000000000000000000000000000000000000081525081565b60006117d03384846122cb565b6001905092915050565b6103e881565b6001600a5414611858576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f556e697377617056323a204c4f434b454400000000000000000000000000000081525060200191505060405180910390fd5b6000600a81905550600060045490506000600554905061192c8284611927600660009054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff16866370a08231306040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156118de57600080fd5b505afa1580156118f2573d6000803e3d6000fd5b505050506040513d602081101561190857600080fd5b8101908080519060200190929190505050611e8390919063ffffffff16565b611be4565b6119ea81846119e56006600e9054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff16856370a08231306040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561199c57600080fd5b505afa1580156119b0573d6000803e3d6000fd5b505050506040513d60208110156119c657600080fd5b8101908080519060200190929190505050611e8390919063ffffffff16565b611be4565b50506001600a8190555050565b60035481565b60055481565b6002602052816000526040600020602052806000526040600020600091509150505481565b6001600a5414611aa0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f556e697377617056323a204c4f434b454400000000000000000000000000000081525060200191505060405180910390fd5b6000600a81905550611bda6004546370a08231306040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015611ae857600080fd5b505afa158015611afc573d6000803e3d6000fd5b505050506040513d6020811015611b1257600080fd5b81019080805190602001909291905050506005546370a08231306040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015611b6057600080fd5b505afa158015611b74573d6000803e3d6000fd5b505050506040513d6020811015611b8a57600080fd5b8101908080519060200190929190505050600660009054906101000a90046dffffffffffffffffffffffffffff166006600e9054906101000a90046dffffffffffffffffffffffffffff16611f06565b6001600a81905550565b600080846040518060400160405280601981526020017f7472616e7366657228616464726573732c75696e7432353629000000000000008152508051906020012085856040516024018083815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040518082805190602001908083835b60208310611cce5780518252602082019150602081019050602083039250611cab565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114611d30576040519150601f19603f3d011682016040523d82523d6000602084013e611d35565b606091505b5091509150818015611d755750600081511480611d745750808060200190516020811015611d6257600080fd5b81019080805190602001909291905050505b5b611de7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f556e697377617056323a205452414e534645525f4641494c454400000000000081525060200191505060405180910390fd5b5050505050565b600080821480611e0b5750828283850292508281611e0857fe5b04145b611e7d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f64732d6d6174682d6d756c2d6f766572666c6f7700000000000000000000000081525060200191505060405180910390fd5b92915050565b6000828284039150811115611f00576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f64732d6d6174682d7375622d756e646572666c6f77000000000000000000000081525060200191505060405180910390fd5b92915050565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6dffffffffffffffffffffffffffff168411158015611f7657507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6dffffffffffffffffffffffffffff168311155b611fe8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f556e697377617056323a204f564552464c4f570000000000000000000000000081525060200191505060405180910390fd5b60006401000000004281611ff857fe5b06905060006006601c9054906101000a900463ffffffff168203905060008163ffffffff1611801561203b57506000846dffffffffffffffffffffffffffff1614155b801561205857506000836dffffffffffffffffffffffffffff1614155b1561213a578063ffffffff1661209d85612071866126a9565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166126d490919063ffffffff16565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16026007600082825401925050819055508063ffffffff1661210b846120df876126a9565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166126d490919063ffffffff16565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16026008600082825401925050819055505b85600660006101000a8154816dffffffffffffffffffffffffffff02191690836dffffffffffffffffffffffffffff160217905550846006600e6101000a8154816dffffffffffffffffffffffffffff02191690836dffffffffffffffffffffffffffff160217905550816006601c6101000a81548163ffffffff021916908363ffffffff1602179055507f1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1600660009054906101000a90046dffffffffffffffffffffffffffff166006600e9054906101000a90046dffffffffffffffffffffffffffff1660405180836dffffffffffffffffffffffffffff168152602001826dffffffffffffffffffffffffffff1681526020019250505060405180910390a1505050505050565b806002600085815260200190815260200160002060008481526020019081526020016000208190555081837f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b6122f1816001600086815260200190815260200160002054611e8390919063ffffffff16565b600160008581526020019081526020016000208190555061232e81600160008581526020019081526020016000205461273490919063ffffffff16565b600160008481526020019081526020016000208190555081837fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b60008060035463017e7e586040518163ffffffff1660e01b815260040160206040518083038186803b1580156123b857600080fd5b505afa1580156123cc573d6000803e3d6000fd5b505050506040513d60208110156123e257600080fd5b81019080805190602001909291905050509050600081141591506000600954905082156124e857600081146124e357600061244e612449866dffffffffffffffffffffffffffff16886dffffffffffffffffffffffffffff16611dee90919063ffffffff16565b612502565b9050600061245b83612502565b9050808211156124e057600061248e61247d8385611e8390919063ffffffff16565b600054611dee90919063ffffffff16565b905060006124b8836124aa600587611dee90919063ffffffff16565b61273490919063ffffffff16565b905060008183816124c557fe5b04905060008111156124dc576124db8782612564565b5b5050505b50505b6124fa565b600081146124f95760006009819055505b5b505092915050565b6000600382111561255157819050600060016002848161251e57fe5b040190505b8181101561254b5780915060028182858161253a57fe5b04018161254357fe5b049050612523565b5061255f565b6000821461255e57600190505b5b919050565b6125798160005461273490919063ffffffff16565b6000819055506125a581600160008581526020019081526020016000205461273490919063ffffffff16565b60016000848152602001908152602001600020819055508160007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b6000818310612609578161260b565b825b905092915050565b612639816001600085815260200190815260200160002054611e8390919063ffffffff16565b600160008481526020019081526020016000208190555061266581600054611e8390919063ffffffff16565b6000819055506000827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b60006e010000000000000000000000000000826dffffffffffffffffffffffffffff16029050919050565b6000816dffffffffffffffffffffffffffff167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff168161272b57fe5b04905092915050565b60008282840191508110156127b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f64732d6d6174682d6164642d6f766572666c6f7700000000000000000000000081525060200191505060405180910390fd5b9291505056fe556e697377617056323a20494e53554646494349454e545f4f55545055545f414d4f554e54556e697377617056323a20494e53554646494349454e545f494e5055545f414d4f554e54556e697377617056323a20494e53554646494349454e545f4c4951554944495459556e697377617056323a20494e53554646494349454e545f4c49515549444954595f4255524e4544556e697377617056323a20494e53554646494349454e545f4c49515549444954595f4d494e544544a26469706673582212208806423f0561522375d4d15ba3785948c0b250d0817eaf87b4f8edb4afe4b90864736f6c63430007060033",
}

// Pairv2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Pairv2MetaData.ABI instead.
var Pairv2ABI = Pairv2MetaData.ABI

// Pairv2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Pairv2MetaData.Bin instead.
var Pairv2Bin = Pairv2MetaData.Bin

// DeployPairv2 deploys a new Ethereum contract, binding an instance of Pairv2 to it.
func DeployPairv2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pairv2, error) {
	parsed, err := Pairv2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Pairv2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pairv2{Pairv2Caller: Pairv2Caller{contract: contract}, Pairv2Transactor: Pairv2Transactor{contract: contract}, Pairv2Filterer: Pairv2Filterer{contract: contract}}, nil
}

// Pairv2 is an auto generated Go binding around an Ethereum contract.
type Pairv2 struct {
	Pairv2Caller     // Read-only binding to the contract
	Pairv2Transactor // Write-only binding to the contract
	Pairv2Filterer   // Log filterer for contract events
}

// Pairv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Pairv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pairv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Pairv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pairv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Pairv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pairv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Pairv2Session struct {
	Contract     *Pairv2           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Pairv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Pairv2CallerSession struct {
	Contract *Pairv2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Pairv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Pairv2TransactorSession struct {
	Contract     *Pairv2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Pairv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Pairv2Raw struct {
	Contract *Pairv2 // Generic contract binding to access the raw methods on
}

// Pairv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Pairv2CallerRaw struct {
	Contract *Pairv2Caller // Generic read-only contract binding to access the raw methods on
}

// Pairv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Pairv2TransactorRaw struct {
	Contract *Pairv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPairv2 creates a new instance of Pairv2, bound to a specific deployed contract.
func NewPairv2(address common.Address, backend bind.ContractBackend) (*Pairv2, error) {
	contract, err := bindPairv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pairv2{Pairv2Caller: Pairv2Caller{contract: contract}, Pairv2Transactor: Pairv2Transactor{contract: contract}, Pairv2Filterer: Pairv2Filterer{contract: contract}}, nil
}

// NewPairv2Caller creates a new read-only instance of Pairv2, bound to a specific deployed contract.
func NewPairv2Caller(address common.Address, caller bind.ContractCaller) (*Pairv2Caller, error) {
	contract, err := bindPairv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Pairv2Caller{contract: contract}, nil
}

// NewPairv2Transactor creates a new write-only instance of Pairv2, bound to a specific deployed contract.
func NewPairv2Transactor(address common.Address, transactor bind.ContractTransactor) (*Pairv2Transactor, error) {
	contract, err := bindPairv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Pairv2Transactor{contract: contract}, nil
}

// NewPairv2Filterer creates a new log filterer instance of Pairv2, bound to a specific deployed contract.
func NewPairv2Filterer(address common.Address, filterer bind.ContractFilterer) (*Pairv2Filterer, error) {
	contract, err := bindPairv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Pairv2Filterer{contract: contract}, nil
}

// bindPairv2 binds a generic wrapper to an already deployed contract.
func bindPairv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Pairv2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairv2 *Pairv2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pairv2.Contract.Pairv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairv2 *Pairv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairv2.Contract.Pairv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairv2 *Pairv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairv2.Contract.Pairv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairv2 *Pairv2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pairv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairv2 *Pairv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairv2 *Pairv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairv2.Contract.contract.Transact(opts, method, params...)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Pairv2 *Pairv2Caller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pairv2.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Pairv2 *Pairv2Session) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _Pairv2.Contract.MINIMUMLIQUIDITY(&_Pairv2.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Pairv2 *Pairv2CallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _Pairv2.Contract.MINIMUMLIQUIDITY(&_Pairv2.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Pairv2 *Pairv2Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pairv2.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Pairv2 *Pairv2Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Pairv2.Contract.Allowance(&_Pairv2.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Pairv2 *Pairv2CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Pairv2.Contract.Allowance(&_Pairv2.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Pairv2 *Pairv2Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pairv2.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Pairv2 *Pairv2Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Pairv2.Contract.BalanceOf(&_Pairv2.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Pairv2 *Pairv2CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Pairv2.Contract.BalanceOf(&_Pairv2.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pairv2 *Pairv2Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Pairv2.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pairv2 *Pairv2Session) Decimals() (uint8, error) {
	return _Pairv2.Contract.Decimals(&_Pairv2.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pairv2 *Pairv2CallerSession) Decimals() (uint8, error) {
	return _Pairv2.Contract.Decimals(&_Pairv2.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pairv2 *Pairv2Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pairv2.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pairv2 *Pairv2Session) Factory() (common.Address, error) {
	return _Pairv2.Contract.Factory(&_Pairv2.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pairv2 *Pairv2CallerSession) Factory() (common.Address, error) {
	return _Pairv2.Contract.Factory(&_Pairv2.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_Pairv2 *Pairv2Caller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _Pairv2.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_Pairv2 *Pairv2Session) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _Pairv2.Contract.GetReserves(&_Pairv2.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_Pairv2 *Pairv2CallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _Pairv2.Contract.GetReserves(&_Pairv2.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Pairv2 *Pairv2Caller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pairv2.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Pairv2 *Pairv2Session) KLast() (*big.Int, error) {
	return _Pairv2.Contract.KLast(&_Pairv2.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Pairv2 *Pairv2CallerSession) KLast() (*big.Int, error) {
	return _Pairv2.Contract.KLast(&_Pairv2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pairv2 *Pairv2Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pairv2.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pairv2 *Pairv2Session) Name() (string, error) {
	return _Pairv2.Contract.Name(&_Pairv2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pairv2 *Pairv2CallerSession) Name() (string, error) {
	return _Pairv2.Contract.Name(&_Pairv2.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Pairv2 *Pairv2Caller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pairv2.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Pairv2 *Pairv2Session) Price0CumulativeLast() (*big.Int, error) {
	return _Pairv2.Contract.Price0CumulativeLast(&_Pairv2.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Pairv2 *Pairv2CallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _Pairv2.Contract.Price0CumulativeLast(&_Pairv2.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Pairv2 *Pairv2Caller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pairv2.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Pairv2 *Pairv2Session) Price1CumulativeLast() (*big.Int, error) {
	return _Pairv2.Contract.Price1CumulativeLast(&_Pairv2.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Pairv2 *Pairv2CallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _Pairv2.Contract.Price1CumulativeLast(&_Pairv2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pairv2 *Pairv2Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pairv2.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pairv2 *Pairv2Session) Symbol() (string, error) {
	return _Pairv2.Contract.Symbol(&_Pairv2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pairv2 *Pairv2CallerSession) Symbol() (string, error) {
	return _Pairv2.Contract.Symbol(&_Pairv2.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pairv2 *Pairv2Caller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pairv2.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pairv2 *Pairv2Session) Token0() (common.Address, error) {
	return _Pairv2.Contract.Token0(&_Pairv2.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pairv2 *Pairv2CallerSession) Token0() (common.Address, error) {
	return _Pairv2.Contract.Token0(&_Pairv2.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pairv2 *Pairv2Caller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pairv2.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pairv2 *Pairv2Session) Token1() (common.Address, error) {
	return _Pairv2.Contract.Token1(&_Pairv2.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pairv2 *Pairv2CallerSession) Token1() (common.Address, error) {
	return _Pairv2.Contract.Token1(&_Pairv2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pairv2 *Pairv2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pairv2.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pairv2 *Pairv2Session) TotalSupply() (*big.Int, error) {
	return _Pairv2.Contract.TotalSupply(&_Pairv2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pairv2 *Pairv2CallerSession) TotalSupply() (*big.Int, error) {
	return _Pairv2.Contract.TotalSupply(&_Pairv2.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Pairv2 *Pairv2Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pairv2.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Pairv2 *Pairv2Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pairv2.Contract.Approve(&_Pairv2.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Pairv2 *Pairv2TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pairv2.Contract.Approve(&_Pairv2.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Pairv2 *Pairv2Transactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Pairv2.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Pairv2 *Pairv2Session) Burn(to common.Address) (*types.Transaction, error) {
	return _Pairv2.Contract.Burn(&_Pairv2.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Pairv2 *Pairv2TransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Pairv2.Contract.Burn(&_Pairv2.TransactOpts, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_Pairv2 *Pairv2Transactor) Initialize(opts *bind.TransactOpts, _token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _Pairv2.contract.Transact(opts, "initialize", _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_Pairv2 *Pairv2Session) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _Pairv2.Contract.Initialize(&_Pairv2.TransactOpts, _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_Pairv2 *Pairv2TransactorSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _Pairv2.Contract.Initialize(&_Pairv2.TransactOpts, _token0, _token1)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pairv2 *Pairv2Transactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Pairv2.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pairv2 *Pairv2Session) Mint(to common.Address) (*types.Transaction, error) {
	return _Pairv2.Contract.Mint(&_Pairv2.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pairv2 *Pairv2TransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Pairv2.Contract.Mint(&_Pairv2.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Pairv2 *Pairv2Transactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Pairv2.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Pairv2 *Pairv2Session) Skim(to common.Address) (*types.Transaction, error) {
	return _Pairv2.Contract.Skim(&_Pairv2.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Pairv2 *Pairv2TransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _Pairv2.Contract.Skim(&_Pairv2.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Pairv2 *Pairv2Transactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Pairv2.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Pairv2 *Pairv2Session) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Pairv2.Contract.Swap(&_Pairv2.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Pairv2 *Pairv2TransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Pairv2.Contract.Swap(&_Pairv2.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Pairv2 *Pairv2Transactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairv2.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Pairv2 *Pairv2Session) Sync() (*types.Transaction, error) {
	return _Pairv2.Contract.Sync(&_Pairv2.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Pairv2 *Pairv2TransactorSession) Sync() (*types.Transaction, error) {
	return _Pairv2.Contract.Sync(&_Pairv2.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Pairv2 *Pairv2Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pairv2.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Pairv2 *Pairv2Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pairv2.Contract.Transfer(&_Pairv2.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Pairv2 *Pairv2TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pairv2.Contract.Transfer(&_Pairv2.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Pairv2 *Pairv2Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pairv2.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Pairv2 *Pairv2Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pairv2.Contract.TransferFrom(&_Pairv2.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Pairv2 *Pairv2TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pairv2.Contract.TransferFrom(&_Pairv2.TransactOpts, from, to, value)
}

// Pairv2ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Pairv2 contract.
type Pairv2ApprovalIterator struct {
	Event *Pairv2Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pairv2ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pairv2Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pairv2Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pairv2ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pairv2ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pairv2Approval represents a Approval event raised by the Pairv2 contract.
type Pairv2Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pairv2 *Pairv2Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Pairv2ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pairv2.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Pairv2ApprovalIterator{contract: _Pairv2.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pairv2 *Pairv2Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Pairv2Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pairv2.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pairv2Approval)
				if err := _Pairv2.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pairv2 *Pairv2Filterer) ParseApproval(log types.Log) (*Pairv2Approval, error) {
	event := new(Pairv2Approval)
	if err := _Pairv2.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pairv2BurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Pairv2 contract.
type Pairv2BurnIterator struct {
	Event *Pairv2Burn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pairv2BurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pairv2Burn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pairv2Burn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pairv2BurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pairv2BurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pairv2Burn represents a Burn event raised by the Pairv2 contract.
type Pairv2Burn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Pairv2 *Pairv2Filterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*Pairv2BurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pairv2.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Pairv2BurnIterator{contract: _Pairv2.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Pairv2 *Pairv2Filterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *Pairv2Burn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pairv2.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pairv2Burn)
				if err := _Pairv2.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Pairv2 *Pairv2Filterer) ParseBurn(log types.Log) (*Pairv2Burn, error) {
	event := new(Pairv2Burn)
	if err := _Pairv2.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pairv2MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Pairv2 contract.
type Pairv2MintIterator struct {
	Event *Pairv2Mint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pairv2MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pairv2Mint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pairv2Mint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pairv2MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pairv2MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pairv2Mint represents a Mint event raised by the Pairv2 contract.
type Pairv2Mint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Pairv2 *Pairv2Filterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*Pairv2MintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pairv2.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &Pairv2MintIterator{contract: _Pairv2.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Pairv2 *Pairv2Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *Pairv2Mint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pairv2.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pairv2Mint)
				if err := _Pairv2.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Pairv2 *Pairv2Filterer) ParseMint(log types.Log) (*Pairv2Mint, error) {
	event := new(Pairv2Mint)
	if err := _Pairv2.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pairv2SwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Pairv2 contract.
type Pairv2SwapIterator struct {
	Event *Pairv2Swap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pairv2SwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pairv2Swap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pairv2Swap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pairv2SwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pairv2SwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pairv2Swap represents a Swap event raised by the Pairv2 contract.
type Pairv2Swap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Pairv2 *Pairv2Filterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*Pairv2SwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pairv2.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Pairv2SwapIterator{contract: _Pairv2.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Pairv2 *Pairv2Filterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *Pairv2Swap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pairv2.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pairv2Swap)
				if err := _Pairv2.contract.UnpackLog(event, "Swap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Pairv2 *Pairv2Filterer) ParseSwap(log types.Log) (*Pairv2Swap, error) {
	event := new(Pairv2Swap)
	if err := _Pairv2.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pairv2SyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the Pairv2 contract.
type Pairv2SyncIterator struct {
	Event *Pairv2Sync // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pairv2SyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pairv2Sync)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pairv2Sync)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pairv2SyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pairv2SyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pairv2Sync represents a Sync event raised by the Pairv2 contract.
type Pairv2Sync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Pairv2 *Pairv2Filterer) FilterSync(opts *bind.FilterOpts) (*Pairv2SyncIterator, error) {

	logs, sub, err := _Pairv2.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &Pairv2SyncIterator{contract: _Pairv2.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Pairv2 *Pairv2Filterer) WatchSync(opts *bind.WatchOpts, sink chan<- *Pairv2Sync) (event.Subscription, error) {

	logs, sub, err := _Pairv2.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pairv2Sync)
				if err := _Pairv2.contract.UnpackLog(event, "Sync", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Pairv2 *Pairv2Filterer) ParseSync(log types.Log) (*Pairv2Sync, error) {
	event := new(Pairv2Sync)
	if err := _Pairv2.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Pairv2TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Pairv2 contract.
type Pairv2TransferIterator struct {
	Event *Pairv2Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pairv2TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pairv2Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pairv2Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pairv2TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pairv2TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pairv2Transfer represents a Transfer event raised by the Pairv2 contract.
type Pairv2Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pairv2 *Pairv2Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Pairv2TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pairv2.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Pairv2TransferIterator{contract: _Pairv2.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pairv2 *Pairv2Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Pairv2Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pairv2.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pairv2Transfer)
				if err := _Pairv2.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pairv2 *Pairv2Filterer) ParseTransfer(log types.Log) (*Pairv2Transfer, error) {
	event := new(Pairv2Transfer)
	if err := _Pairv2.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

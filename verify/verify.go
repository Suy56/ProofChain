// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verify

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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

// VerifyMetaData contains all meta data concerning the Verify contract.
var VerifyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_EncryptedIPFSHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_institute\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"addDocument\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"approveVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDocuments\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"requester\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"verifer\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"institute\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"ipfs\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"name\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"desc\",\"type\":\"string[]\"},{\"internalType\":\"enumVerification.DocStatus[]\",\"name\":\"stats\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getInstituePublicKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"pubKey\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"institutions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"publicAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"}],\"name\":\"registerAsUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_publicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"registerInstitution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_institute\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ipfs\",\"type\":\"string\"},{\"internalType\":\"enumVerification.DocStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"verifyDocument\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600d5534801561001557600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611d72806100656000396000f3fe60806040526004361061007b5760003560e01c8063ba7857251161004e578063ba7857251461012b578063c6b5c5cc14610154578063ea6509f81461017d578063ef2d8700146101bd5761007b565b80630e72ca7a1461008057806320857a9f146100a957806390752bb2146100d2578063ba1f3dc7146100ee575b600080fd5b34801561008c57600080fd5b506100a760048036038101906100a291906114b4565b6101ee565b005b3480156100b557600080fd5b506100d060048036038101906100cb9190611313565b6104fc565b005b6100ec60048036038101906100e79190611435565b610673565b005b3480156100fa57600080fd5b5061011560048036038101906101109190611388565b610870565b604051610122919061197b565b60405180910390f35b34801561013757600080fd5b50610152600480360381019061014d9190611388565b610923565b005b34801561016057600080fd5b5061017b600480360381019061017691906113c9565b6109ed565b005b34801561018957600080fd5b506101a4600480360381019061019f9190611388565b610aed565b6040516101b49493929190611888565b60405180910390f35b3480156101c957600080fd5b506101d2610c70565b6040516101e597969594939291906118db565b60405180910390f35b6000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1690506001151581151514610286576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027d906119bd565b60405180910390fd5b600d546003866040516102999190611871565b9081526020016040518091039020819055506006339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600760009080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506009849080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906103ad9291906111a9565b506008839080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906103e99291906111a9565b50600a829080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906104259291906111a9565b50600c600290806001815401808255809150506001900390600052602060002090602091828204019190069091909190916101000a81548160ff0219169083600281111561049c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b0217905550600b859080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906104dc9291906111a9565b50600d60008154809291906104f090611c02565b91905055505050505050565b604051806040016040528085858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200183838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815250600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000190805190602001906105f49291906111a9565b5060208201518160010190805190602001906106119291906111a9565b509050506001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050505050565b600115156001846040516106879190611871565b908152602001604051809103902060030160009054906101000a900460ff16151514801561072157503373ffffffffffffffffffffffffffffffffffffffff166001846040516106d79190611871565b908152602001604051809103902060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b61072a57600080fd5b600060038360405161073c9190611871565b908152602001604051809103902054905081600c8281548110610788577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002090602091828204019190066101000a81548160ff021916908360028111156107e2577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055503360078281548110610822577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b60606001826040516108829190611871565b9081526020016040518091039020600101805461089e90611bd0565b80601f01602080910402602001604051908101604052809291908181526020018280546108ca90611bd0565b80156109175780601f106108ec57610100808354040283529160200191610917565b820191906000526020600020905b8154815290600101906020018083116108fa57829003601f168201915b50505050509050919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109a89061199d565b60405180910390fd5b600180826040516109c29190611871565b908152602001604051809103902060030160006101000a81548160ff02191690831515021790555050565b60405180608001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182815260200160001515815250600182604051610a379190611871565b908152602001604051809103902060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019080519060200190610aa89291906111a9565b506040820151816002019080519060200190610ac59291906111a9565b5060608201518160030160006101000a81548160ff0219169083151502179055509050505050565b6001818051602081018201805184825260208301602085012081835280955050505050506000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054610b4c90611bd0565b80601f0160208091040260200160405190810160405280929190818152602001828054610b7890611bd0565b8015610bc55780601f10610b9a57610100808354040283529160200191610bc5565b820191906000526020600020905b815481529060010190602001808311610ba857829003601f168201915b505050505090806002018054610bda90611bd0565b80601f0160208091040260200160405190810160405280929190818152602001828054610c0690611bd0565b8015610c535780601f10610c2857610100808354040283529160200191610c53565b820191906000526020600020905b815481529060010190602001808311610c3657829003601f168201915b5050505050908060030160009054906101000a900460ff16905084565b6060806060806060806060600660076009600b6008600a600c86805480602002602001604051908101604052809291908181526020018280548015610d0a57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610cc0575b5050505050965085805480602002602001604051908101604052809291908181526020018280548015610d9257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610d48575b5050505050955084805480602002602001604051908101604052809291908181526020016000905b82821015610e66578382906000526020600020018054610dd990611bd0565b80601f0160208091040260200160405190810160405280929190818152602001828054610e0590611bd0565b8015610e525780601f10610e2757610100808354040283529160200191610e52565b820191906000526020600020905b815481529060010190602001808311610e3557829003601f168201915b505050505081526020019060010190610dba565b50505050945083805480602002602001604051908101604052809291908181526020016000905b82821015610f39578382906000526020600020018054610eac90611bd0565b80601f0160208091040260200160405190810160405280929190818152602001828054610ed890611bd0565b8015610f255780601f10610efa57610100808354040283529160200191610f25565b820191906000526020600020905b815481529060010190602001808311610f0857829003601f168201915b505050505081526020019060010190610e8d565b50505050935082805480602002602001604051908101604052809291908181526020016000905b8282101561100c578382906000526020600020018054610f7f90611bd0565b80601f0160208091040260200160405190810160405280929190818152602001828054610fab90611bd0565b8015610ff85780601f10610fcd57610100808354040283529160200191610ff8565b820191906000526020600020905b815481529060010190602001808311610fdb57829003601f168201915b505050505081526020019060010190610f60565b50505050925081805480602002602001604051908101604052809291908181526020016000905b828210156110df57838290600052602060002001805461105290611bd0565b80601f016020809104026020016040519081016040528092919081815260200182805461107e90611bd0565b80156110cb5780601f106110a0576101008083540402835291602001916110cb565b820191906000526020600020905b8154815290600101906020018083116110ae57829003601f168201915b505050505081526020019060010190611033565b5050505091508080548060200260200160405190810160405280929190818152602001828054801561118b57602002820191906000526020600020906000905b82829054906101000a900460ff166002811115611165577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001906001019060208260000104928301926001038202915080841161111f5790505b50505050509050965096509650965096509650965090919293949596565b8280546111b590611bd0565b90600052602060002090601f0160209004810192826111d7576000855561121e565b82601f106111f057805160ff191683800117855561121e565b8280016001018555821561121e579182015b8281111561121d578251825591602001919060010190611202565b5b50905061122b919061122f565b5090565b5b80821115611248576000816000905550600101611230565b5090565b600061125f61125a84611a0e565b6119dd565b90508281526020810184848401111561127757600080fd5b611282848285611b8e565b509392505050565b60008135905061129981611d2c565b92915050565b60008083601f8401126112b157600080fd5b8235905067ffffffffffffffff8111156112ca57600080fd5b6020830191508360018202830111156112e257600080fd5b9250929050565b600082601f8301126112fa57600080fd5b813561130a84826020860161124c565b91505092915050565b6000806000806040858703121561132957600080fd5b600085013567ffffffffffffffff81111561134357600080fd5b61134f8782880161129f565b9450945050602085013567ffffffffffffffff81111561136e57600080fd5b61137a8782880161129f565b925092505092959194509250565b60006020828403121561139a57600080fd5b600082013567ffffffffffffffff8111156113b457600080fd5b6113c0848285016112e9565b91505092915050565b600080604083850312156113dc57600080fd5b600083013567ffffffffffffffff8111156113f657600080fd5b611402858286016112e9565b925050602083013567ffffffffffffffff81111561141f57600080fd5b61142b858286016112e9565b9150509250929050565b60008060006060848603121561144a57600080fd5b600084013567ffffffffffffffff81111561146457600080fd5b611470868287016112e9565b935050602084013567ffffffffffffffff81111561148d57600080fd5b611499868287016112e9565b92505060406114aa8682870161128a565b9150509250925092565b600080600080608085870312156114ca57600080fd5b600085013567ffffffffffffffff8111156114e457600080fd5b6114f0878288016112e9565b945050602085013567ffffffffffffffff81111561150d57600080fd5b611519878288016112e9565b935050604085013567ffffffffffffffff81111561153657600080fd5b611542878288016112e9565b925050606085013567ffffffffffffffff81111561155f57600080fd5b61156b878288016112e9565b91505092959194509250565b600061158383836115bb565b60208301905092915050565b600061159b8383611719565b60208301905092915050565b60006115b38383611728565b905092915050565b6115c481611b21565b82525050565b6115d381611b21565b82525050565b60006115e482611a6e565b6115ee8185611ac1565b93506115f983611a3e565b8060005b8381101561162a5781516116118882611577565b975061161c83611a9a565b9250506001810190506115fd565b5085935050505092915050565b600061164282611a79565b61164c8185611ae3565b935061165783611a4e565b8060005b8381101561168857815161166f888261158f565b975061167a83611aa7565b92505060018101905061165b565b5085935050505092915050565b60006116a082611a84565b6116aa8185611ad2565b9350836020820285016116bc85611a5e565b8060005b858110156116f857848403895281516116d985826115a7565b94506116e483611ab4565b925060208a019950506001810190506116c0565b50829750879550505050505092915050565b61171381611b33565b82525050565b61172281611b7c565b82525050565b600061173382611a8f565b61173d8185611af4565b935061174d818560208601611b9d565b61175681611d07565b840191505092915050565b600061176c82611a8f565b6117768185611b05565b9350611786818560208601611b9d565b61178f81611d07565b840191505092915050565b60006117a582611a8f565b6117af8185611b16565b93506117bf818560208601611b9d565b80840191505092915050565b60006117d8602183611b05565b91507f4f6e6c792061646d696e2063616e20706572666f6d207468697320616374696f60008301527f6e000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b600061183e601883611b05565b91507f726567697374657220666972737420746f2076657269667900000000000000006000830152602082019050919050565b600061187d828461179a565b915081905092915050565b600060808201905061189d60008301876115ca565b81810360208301526118af8186611761565b905081810360408301526118c38185611761565b90506118d2606083018461170a565b95945050505050565b600060e08201905081810360008301526118f5818a6115d9565b9050818103602083015261190981896115d9565b9050818103604083015261191d8188611695565b905081810360608301526119318187611695565b905081810360808301526119458186611695565b905081810360a08301526119598185611695565b905081810360c083015261196d8184611637565b905098975050505050505050565b600060208201905081810360008301526119958184611761565b905092915050565b600060208201905081810360008301526119b6816117cb565b9050919050565b600060208201905081810360008301526119d681611831565b9050919050565b6000604051905081810181811067ffffffffffffffff82111715611a0457611a03611cd8565b5b8060405250919050565b600067ffffffffffffffff821115611a2957611a28611cd8565b5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000611b2c82611b52565b9050919050565b60008115159050919050565b6000819050611b4d82611d18565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611b8782611b3f565b9050919050565b82818337600083830152505050565b60005b83811015611bbb578082015181840152602081019050611ba0565b83811115611bca576000848401525b50505050565b60006002820490506001821680611be857607f821691505b60208210811415611bfc57611bfb611ca9565b5b50919050565b6000611c0d82611b72565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611c4057611c3f611c4b565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60038110611d2957611d28611c7a565b5b50565b60038110611d3957600080fd5b5056fea26469706673582212209b7d8c657f28b0fc4ab8e175c39886d2b5b47196df5fdaae02b6f0f60b66969364736f6c63430008000033",
}

// VerifyABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifyMetaData.ABI instead.
var VerifyABI = VerifyMetaData.ABI

// VerifyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerifyMetaData.Bin instead.
var VerifyBin = VerifyMetaData.Bin

// DeployVerify deploys a new Ethereum contract, binding an instance of Verify to it.
func DeployVerify(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verify, error) {
	parsed, err := VerifyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verify{VerifyCaller: VerifyCaller{contract: contract}, VerifyTransactor: VerifyTransactor{contract: contract}, VerifyFilterer: VerifyFilterer{contract: contract}}, nil
}

// Verify is an auto generated Go binding around an Ethereum contract.
type Verify struct {
	VerifyCaller     // Read-only binding to the contract
	VerifyTransactor // Write-only binding to the contract
	VerifyFilterer   // Log filterer for contract events
}

// VerifyCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifySession struct {
	Contract     *Verify           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifyCallerSession struct {
	Contract *VerifyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VerifyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifyTransactorSession struct {
	Contract     *VerifyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifyRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifyRaw struct {
	Contract *Verify // Generic contract binding to access the raw methods on
}

// VerifyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifyCallerRaw struct {
	Contract *VerifyCaller // Generic read-only contract binding to access the raw methods on
}

// VerifyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifyTransactorRaw struct {
	Contract *VerifyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerify creates a new instance of Verify, bound to a specific deployed contract.
func NewVerify(address common.Address, backend bind.ContractBackend) (*Verify, error) {
	contract, err := bindVerify(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verify{VerifyCaller: VerifyCaller{contract: contract}, VerifyTransactor: VerifyTransactor{contract: contract}, VerifyFilterer: VerifyFilterer{contract: contract}}, nil
}

// NewVerifyCaller creates a new read-only instance of Verify, bound to a specific deployed contract.
func NewVerifyCaller(address common.Address, caller bind.ContractCaller) (*VerifyCaller, error) {
	contract, err := bindVerify(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyCaller{contract: contract}, nil
}

// NewVerifyTransactor creates a new write-only instance of Verify, bound to a specific deployed contract.
func NewVerifyTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifyTransactor, error) {
	contract, err := bindVerify(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyTransactor{contract: contract}, nil
}

// NewVerifyFilterer creates a new log filterer instance of Verify, bound to a specific deployed contract.
func NewVerifyFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifyFilterer, error) {
	contract, err := bindVerify(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifyFilterer{contract: contract}, nil
}

// bindVerify binds a generic wrapper to an already deployed contract.
func bindVerify(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verify *VerifyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verify.Contract.VerifyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verify *VerifyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verify.Contract.VerifyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verify *VerifyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verify.Contract.VerifyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verify *VerifyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verify.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verify *VerifyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verify.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verify *VerifyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verify.Contract.contract.Transact(opts, method, params...)
}

// GetDocuments is a free data retrieval call binding the contract method 0xef2d8700.
//
// Solidity: function getDocuments() view returns(address[] requester, address[] verifer, string[] institute, string[] ipfs, string[] name, string[] desc, uint8[] stats)
func (_Verify *VerifyCaller) GetDocuments(opts *bind.CallOpts) (struct {
	Requester []common.Address
	Verifer   []common.Address
	Institute []string
	Ipfs      []string
	Name      []string
	Desc      []string
	Stats     []uint8
}, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "getDocuments")

	outstruct := new(struct {
		Requester []common.Address
		Verifer   []common.Address
		Institute []string
		Ipfs      []string
		Name      []string
		Desc      []string
		Stats     []uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Requester = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Verifer = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Institute = *abi.ConvertType(out[2], new([]string)).(*[]string)
	outstruct.Ipfs = *abi.ConvertType(out[3], new([]string)).(*[]string)
	outstruct.Name = *abi.ConvertType(out[4], new([]string)).(*[]string)
	outstruct.Desc = *abi.ConvertType(out[5], new([]string)).(*[]string)
	outstruct.Stats = *abi.ConvertType(out[6], new([]uint8)).(*[]uint8)

	return *outstruct, err

}

// GetDocuments is a free data retrieval call binding the contract method 0xef2d8700.
//
// Solidity: function getDocuments() view returns(address[] requester, address[] verifer, string[] institute, string[] ipfs, string[] name, string[] desc, uint8[] stats)
func (_Verify *VerifySession) GetDocuments() (struct {
	Requester []common.Address
	Verifer   []common.Address
	Institute []string
	Ipfs      []string
	Name      []string
	Desc      []string
	Stats     []uint8
}, error) {
	return _Verify.Contract.GetDocuments(&_Verify.CallOpts)
}

// GetDocuments is a free data retrieval call binding the contract method 0xef2d8700.
//
// Solidity: function getDocuments() view returns(address[] requester, address[] verifer, string[] institute, string[] ipfs, string[] name, string[] desc, uint8[] stats)
func (_Verify *VerifyCallerSession) GetDocuments() (struct {
	Requester []common.Address
	Verifer   []common.Address
	Institute []string
	Ipfs      []string
	Name      []string
	Desc      []string
	Stats     []uint8
}, error) {
	return _Verify.Contract.GetDocuments(&_Verify.CallOpts)
}

// GetInstituePublicKey is a free data retrieval call binding the contract method 0xba1f3dc7.
//
// Solidity: function getInstituePublicKey(string _name) view returns(string pubKey)
func (_Verify *VerifyCaller) GetInstituePublicKey(opts *bind.CallOpts, _name string) (string, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "getInstituePublicKey", _name)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetInstituePublicKey is a free data retrieval call binding the contract method 0xba1f3dc7.
//
// Solidity: function getInstituePublicKey(string _name) view returns(string pubKey)
func (_Verify *VerifySession) GetInstituePublicKey(_name string) (string, error) {
	return _Verify.Contract.GetInstituePublicKey(&_Verify.CallOpts, _name)
}

// GetInstituePublicKey is a free data retrieval call binding the contract method 0xba1f3dc7.
//
// Solidity: function getInstituePublicKey(string _name) view returns(string pubKey)
func (_Verify *VerifyCallerSession) GetInstituePublicKey(_name string) (string, error) {
	return _Verify.Contract.GetInstituePublicKey(&_Verify.CallOpts, _name)
}

// Institutions is a free data retrieval call binding the contract method 0xea6509f8.
//
// Solidity: function institutions(string ) view returns(address publicAddr, string publicKey, string name, bool approved)
func (_Verify *VerifyCaller) Institutions(opts *bind.CallOpts, arg0 string) (struct {
	PublicAddr common.Address
	PublicKey  string
	Name       string
	Approved   bool
}, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "institutions", arg0)

	outstruct := new(struct {
		PublicAddr common.Address
		PublicKey  string
		Name       string
		Approved   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PublicAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PublicKey = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Name = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Approved = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Institutions is a free data retrieval call binding the contract method 0xea6509f8.
//
// Solidity: function institutions(string ) view returns(address publicAddr, string publicKey, string name, bool approved)
func (_Verify *VerifySession) Institutions(arg0 string) (struct {
	PublicAddr common.Address
	PublicKey  string
	Name       string
	Approved   bool
}, error) {
	return _Verify.Contract.Institutions(&_Verify.CallOpts, arg0)
}

// Institutions is a free data retrieval call binding the contract method 0xea6509f8.
//
// Solidity: function institutions(string ) view returns(address publicAddr, string publicKey, string name, bool approved)
func (_Verify *VerifyCallerSession) Institutions(arg0 string) (struct {
	PublicAddr common.Address
	PublicKey  string
	Name       string
	Approved   bool
}, error) {
	return _Verify.Contract.Institutions(&_Verify.CallOpts, arg0)
}

// AddDocument is a paid mutator transaction binding the contract method 0x0e72ca7a.
//
// Solidity: function addDocument(string _EncryptedIPFSHash, string _institute, string _name, string _description) returns()
func (_Verify *VerifyTransactor) AddDocument(opts *bind.TransactOpts, _EncryptedIPFSHash string, _institute string, _name string, _description string) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "addDocument", _EncryptedIPFSHash, _institute, _name, _description)
}

// AddDocument is a paid mutator transaction binding the contract method 0x0e72ca7a.
//
// Solidity: function addDocument(string _EncryptedIPFSHash, string _institute, string _name, string _description) returns()
func (_Verify *VerifySession) AddDocument(_EncryptedIPFSHash string, _institute string, _name string, _description string) (*types.Transaction, error) {
	return _Verify.Contract.AddDocument(&_Verify.TransactOpts, _EncryptedIPFSHash, _institute, _name, _description)
}

// AddDocument is a paid mutator transaction binding the contract method 0x0e72ca7a.
//
// Solidity: function addDocument(string _EncryptedIPFSHash, string _institute, string _name, string _description) returns()
func (_Verify *VerifyTransactorSession) AddDocument(_EncryptedIPFSHash string, _institute string, _name string, _description string) (*types.Transaction, error) {
	return _Verify.Contract.AddDocument(&_Verify.TransactOpts, _EncryptedIPFSHash, _institute, _name, _description)
}

// ApproveVerifier is a paid mutator transaction binding the contract method 0xba785725.
//
// Solidity: function approveVerifier(string _name) returns()
func (_Verify *VerifyTransactor) ApproveVerifier(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "approveVerifier", _name)
}

// ApproveVerifier is a paid mutator transaction binding the contract method 0xba785725.
//
// Solidity: function approveVerifier(string _name) returns()
func (_Verify *VerifySession) ApproveVerifier(_name string) (*types.Transaction, error) {
	return _Verify.Contract.ApproveVerifier(&_Verify.TransactOpts, _name)
}

// ApproveVerifier is a paid mutator transaction binding the contract method 0xba785725.
//
// Solidity: function approveVerifier(string _name) returns()
func (_Verify *VerifyTransactorSession) ApproveVerifier(_name string) (*types.Transaction, error) {
	return _Verify.Contract.ApproveVerifier(&_Verify.TransactOpts, _name)
}

// RegisterAsUser is a paid mutator transaction binding the contract method 0x20857a9f.
//
// Solidity: function registerAsUser(string _name, string _email) returns()
func (_Verify *VerifyTransactor) RegisterAsUser(opts *bind.TransactOpts, _name string, _email string) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "registerAsUser", _name, _email)
}

// RegisterAsUser is a paid mutator transaction binding the contract method 0x20857a9f.
//
// Solidity: function registerAsUser(string _name, string _email) returns()
func (_Verify *VerifySession) RegisterAsUser(_name string, _email string) (*types.Transaction, error) {
	return _Verify.Contract.RegisterAsUser(&_Verify.TransactOpts, _name, _email)
}

// RegisterAsUser is a paid mutator transaction binding the contract method 0x20857a9f.
//
// Solidity: function registerAsUser(string _name, string _email) returns()
func (_Verify *VerifyTransactorSession) RegisterAsUser(_name string, _email string) (*types.Transaction, error) {
	return _Verify.Contract.RegisterAsUser(&_Verify.TransactOpts, _name, _email)
}

// RegisterInstitution is a paid mutator transaction binding the contract method 0xc6b5c5cc.
//
// Solidity: function registerInstitution(string _publicKey, string _name) returns()
func (_Verify *VerifyTransactor) RegisterInstitution(opts *bind.TransactOpts, _publicKey string, _name string) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "registerInstitution", _publicKey, _name)
}

// RegisterInstitution is a paid mutator transaction binding the contract method 0xc6b5c5cc.
//
// Solidity: function registerInstitution(string _publicKey, string _name) returns()
func (_Verify *VerifySession) RegisterInstitution(_publicKey string, _name string) (*types.Transaction, error) {
	return _Verify.Contract.RegisterInstitution(&_Verify.TransactOpts, _publicKey, _name)
}

// RegisterInstitution is a paid mutator transaction binding the contract method 0xc6b5c5cc.
//
// Solidity: function registerInstitution(string _publicKey, string _name) returns()
func (_Verify *VerifyTransactorSession) RegisterInstitution(_publicKey string, _name string) (*types.Transaction, error) {
	return _Verify.Contract.RegisterInstitution(&_Verify.TransactOpts, _publicKey, _name)
}

// VerifyDocument is a paid mutator transaction binding the contract method 0x90752bb2.
//
// Solidity: function verifyDocument(string _institute, string _ipfs, uint8 _status) payable returns()
func (_Verify *VerifyTransactor) VerifyDocument(opts *bind.TransactOpts, _institute string, _ipfs string, _status uint8) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "verifyDocument", _institute, _ipfs, _status)
}

// VerifyDocument is a paid mutator transaction binding the contract method 0x90752bb2.
//
// Solidity: function verifyDocument(string _institute, string _ipfs, uint8 _status) payable returns()
func (_Verify *VerifySession) VerifyDocument(_institute string, _ipfs string, _status uint8) (*types.Transaction, error) {
	return _Verify.Contract.VerifyDocument(&_Verify.TransactOpts, _institute, _ipfs, _status)
}

// VerifyDocument is a paid mutator transaction binding the contract method 0x90752bb2.
//
// Solidity: function verifyDocument(string _institute, string _ipfs, uint8 _status) payable returns()
func (_Verify *VerifyTransactorSession) VerifyDocument(_institute string, _ipfs string, _status uint8) (*types.Transaction, error) {
	return _Verify.Contract.VerifyDocument(&_Verify.TransactOpts, _institute, _ipfs, _status)
}

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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_EncryptedIPFSHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_institute\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"addDocument\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"approveVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDocuments\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"requester\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"verifer\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"institute\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"ipfs\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"name\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"desc\",\"type\":\"string[]\"},{\"internalType\":\"enumVerification.DocStatus[]\",\"name\":\"stats\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getInstituePublicKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"pubKey\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"getUserPublicKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"institutions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"publicAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_publicKey\",\"type\":\"string\"}],\"name\":\"registerAsUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_publicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"registerInstitution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_institute\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ipfs\",\"type\":\"string\"},{\"internalType\":\"enumVerification.DocStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"verifyDocument\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600d5534801561001557600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611e4a806100656000396000f3fe6080604052600436106100865760003560e01c8063ba1f3dc711610059578063ba1f3dc714610136578063ba78572514610173578063c6b5c5cc1461019c578063ea6509f8146101c5578063ef2d87001461020557610086565b80630e72ca7a1461008b57806311231fe0146100b4578063515ef106146100f157806390752bb21461011a575b600080fd5b34801561009757600080fd5b506100b260048036038101906100ad9190611575565b610236565b005b3480156100c057600080fd5b506100db60048036038101906100d691906113db565b610544565b6040516100e89190611a3c565b60405180910390f35b3480156100fd57600080fd5b5061011860048036038101906101139190611404565b610618565b005b610134600480360381019061012f91906114f6565b610726565b005b34801561014257600080fd5b5061015d60048036038101906101589190611449565b610923565b60405161016a9190611a3c565b60405180910390f35b34801561017f57600080fd5b5061019a60048036038101906101959190611449565b6109d6565b005b3480156101a857600080fd5b506101c360048036038101906101be919061148a565b610aa0565b005b3480156101d157600080fd5b506101ec60048036038101906101e79190611449565b610ba0565b6040516101fc9493929190611949565b60405180910390f35b34801561021157600080fd5b5061021a610d23565b60405161022d979695949392919061199c565b60405180910390f35b6000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905060011515811515146102ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c590611a7e565b60405180910390fd5b600d546003866040516102e19190611932565b9081526020016040518091039020819055506006339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600760009080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506009849080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906103f592919061125c565b5060088390806001815401808255809150506001900390600052602060002001600090919091909150908051906020019061043192919061125c565b50600a8290806001815401808255809150506001900390600052602060002001600090919091909150908051906020019061046d92919061125c565b50600c600290806001815401808255809150506001900390600052602060002090602091828204019190069091909190916101000a81548160ff021916908360028111156104e4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b0217905550600b8590806001815401808255809150506001900390600052602060002001600090919091909150908051906020019061052492919061125c565b50600d600081548092919061053890611cc3565b91905055505050505050565b6060600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001805461059390611c91565b80601f01602080910402602001604051908101604052809291908181526020018280546105bf90611c91565b801561060c5780601f106105e15761010080835404028352916020019161060c565b820191906000526020600020905b8154815290600101906020018083116105ef57829003601f168201915b50505050509050919050565b604051806020016040528083838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815250600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000190805190602001906106c692919061125c565b509050506001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6001151560018460405161073a9190611932565b908152602001604051809103902060030160009054906101000a900460ff1615151480156107d457503373ffffffffffffffffffffffffffffffffffffffff1660018460405161078a9190611932565b908152602001604051809103902060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b6107dd57600080fd5b60006003836040516107ef9190611932565b908152602001604051809103902054905081600c828154811061083b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002090602091828204019190066101000a81548160ff02191690836002811115610895577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555033600782815481106108d5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b60606001826040516109359190611932565b9081526020016040518091039020600101805461095190611c91565b80601f016020809104026020016040519081016040528092919081815260200182805461097d90611c91565b80156109ca5780601f1061099f576101008083540402835291602001916109ca565b820191906000526020600020905b8154815290600101906020018083116109ad57829003601f168201915b50505050509050919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a64576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5b90611a5e565b60405180910390fd5b60018082604051610a759190611932565b908152602001604051809103902060030160006101000a81548160ff02191690831515021790555050565b60405180608001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182815260200160001515815250600182604051610aea9190611932565b908152602001604051809103902060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019080519060200190610b5b92919061125c565b506040820151816002019080519060200190610b7892919061125c565b5060608201518160030160006101000a81548160ff0219169083151502179055509050505050565b6001818051602081018201805184825260208301602085012081835280955050505050506000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054610bff90611c91565b80601f0160208091040260200160405190810160405280929190818152602001828054610c2b90611c91565b8015610c785780601f10610c4d57610100808354040283529160200191610c78565b820191906000526020600020905b815481529060010190602001808311610c5b57829003601f168201915b505050505090806002018054610c8d90611c91565b80601f0160208091040260200160405190810160405280929190818152602001828054610cb990611c91565b8015610d065780601f10610cdb57610100808354040283529160200191610d06565b820191906000526020600020905b815481529060010190602001808311610ce957829003601f168201915b5050505050908060030160009054906101000a900460ff16905084565b6060806060806060806060600660076009600b6008600a600c86805480602002602001604051908101604052809291908181526020018280548015610dbd57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610d73575b5050505050965085805480602002602001604051908101604052809291908181526020018280548015610e4557602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610dfb575b5050505050955084805480602002602001604051908101604052809291908181526020016000905b82821015610f19578382906000526020600020018054610e8c90611c91565b80601f0160208091040260200160405190810160405280929190818152602001828054610eb890611c91565b8015610f055780601f10610eda57610100808354040283529160200191610f05565b820191906000526020600020905b815481529060010190602001808311610ee857829003601f168201915b505050505081526020019060010190610e6d565b50505050945083805480602002602001604051908101604052809291908181526020016000905b82821015610fec578382906000526020600020018054610f5f90611c91565b80601f0160208091040260200160405190810160405280929190818152602001828054610f8b90611c91565b8015610fd85780601f10610fad57610100808354040283529160200191610fd8565b820191906000526020600020905b815481529060010190602001808311610fbb57829003601f168201915b505050505081526020019060010190610f40565b50505050935082805480602002602001604051908101604052809291908181526020016000905b828210156110bf57838290600052602060002001805461103290611c91565b80601f016020809104026020016040519081016040528092919081815260200182805461105e90611c91565b80156110ab5780601f10611080576101008083540402835291602001916110ab565b820191906000526020600020905b81548152906001019060200180831161108e57829003601f168201915b505050505081526020019060010190611013565b50505050925081805480602002602001604051908101604052809291908181526020016000905b8282101561119257838290600052602060002001805461110590611c91565b80601f016020809104026020016040519081016040528092919081815260200182805461113190611c91565b801561117e5780601f106111535761010080835404028352916020019161117e565b820191906000526020600020905b81548152906001019060200180831161116157829003601f168201915b5050505050815260200190600101906110e6565b5050505091508080548060200260200160405190810160405280929190818152602001828054801561123e57602002820191906000526020600020906000905b82829054906101000a900460ff166002811115611218577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200190600101906020826000010492830192600103820291508084116111d25790505b50505050509050965096509650965096509650965090919293949596565b82805461126890611c91565b90600052602060002090601f01602090048101928261128a57600085556112d1565b82601f106112a357805160ff19168380011785556112d1565b828001600101855582156112d1579182015b828111156112d05782518255916020019190600101906112b5565b5b5090506112de91906112e2565b5090565b5b808211156112fb5760008160009055506001016112e3565b5090565b600061131261130d84611acf565b611a9e565b90508281526020810184848401111561132a57600080fd5b611335848285611c4f565b509392505050565b60008135905061134c81611ded565b92915050565b60008135905061136181611e04565b92915050565b60008083601f84011261137957600080fd5b8235905067ffffffffffffffff81111561139257600080fd5b6020830191508360018202830111156113aa57600080fd5b9250929050565b600082601f8301126113c257600080fd5b81356113d28482602086016112ff565b91505092915050565b6000602082840312156113ed57600080fd5b60006113fb8482850161133d565b91505092915050565b6000806020838503121561141757600080fd5b600083013567ffffffffffffffff81111561143157600080fd5b61143d85828601611367565b92509250509250929050565b60006020828403121561145b57600080fd5b600082013567ffffffffffffffff81111561147557600080fd5b611481848285016113b1565b91505092915050565b6000806040838503121561149d57600080fd5b600083013567ffffffffffffffff8111156114b757600080fd5b6114c3858286016113b1565b925050602083013567ffffffffffffffff8111156114e057600080fd5b6114ec858286016113b1565b9150509250929050565b60008060006060848603121561150b57600080fd5b600084013567ffffffffffffffff81111561152557600080fd5b611531868287016113b1565b935050602084013567ffffffffffffffff81111561154e57600080fd5b61155a868287016113b1565b925050604061156b86828701611352565b9150509250925092565b6000806000806080858703121561158b57600080fd5b600085013567ffffffffffffffff8111156115a557600080fd5b6115b1878288016113b1565b945050602085013567ffffffffffffffff8111156115ce57600080fd5b6115da878288016113b1565b935050604085013567ffffffffffffffff8111156115f757600080fd5b611603878288016113b1565b925050606085013567ffffffffffffffff81111561162057600080fd5b61162c878288016113b1565b91505092959194509250565b6000611644838361167c565b60208301905092915050565b600061165c83836117da565b60208301905092915050565b600061167483836117e9565b905092915050565b61168581611be2565b82525050565b61169481611be2565b82525050565b60006116a582611b2f565b6116af8185611b82565b93506116ba83611aff565b8060005b838110156116eb5781516116d28882611638565b97506116dd83611b5b565b9250506001810190506116be565b5085935050505092915050565b600061170382611b3a565b61170d8185611ba4565b935061171883611b0f565b8060005b838110156117495781516117308882611650565b975061173b83611b68565b92505060018101905061171c565b5085935050505092915050565b600061176182611b45565b61176b8185611b93565b93508360208202850161177d85611b1f565b8060005b858110156117b9578484038952815161179a8582611668565b94506117a583611b75565b925060208a01995050600181019050611781565b50829750879550505050505092915050565b6117d481611bf4565b82525050565b6117e381611c3d565b82525050565b60006117f482611b50565b6117fe8185611bb5565b935061180e818560208601611c5e565b61181781611dc8565b840191505092915050565b600061182d82611b50565b6118378185611bc6565b9350611847818560208601611c5e565b61185081611dc8565b840191505092915050565b600061186682611b50565b6118708185611bd7565b9350611880818560208601611c5e565b80840191505092915050565b6000611899602183611bc6565b91507f4f6e6c792061646d696e2063616e20706572666f6d207468697320616374696f60008301527f6e000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006118ff601883611bc6565b91507f726567697374657220666972737420746f2076657269667900000000000000006000830152602082019050919050565b600061193e828461185b565b915081905092915050565b600060808201905061195e600083018761168b565b81810360208301526119708186611822565b905081810360408301526119848185611822565b905061199360608301846117cb565b95945050505050565b600060e08201905081810360008301526119b6818a61169a565b905081810360208301526119ca818961169a565b905081810360408301526119de8188611756565b905081810360608301526119f28187611756565b90508181036080830152611a068186611756565b905081810360a0830152611a1a8185611756565b905081810360c0830152611a2e81846116f8565b905098975050505050505050565b60006020820190508181036000830152611a568184611822565b905092915050565b60006020820190508181036000830152611a778161188c565b9050919050565b60006020820190508181036000830152611a97816118f2565b9050919050565b6000604051905081810181811067ffffffffffffffff82111715611ac557611ac4611d99565b5b8060405250919050565b600067ffffffffffffffff821115611aea57611ae9611d99565b5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000611bed82611c13565b9050919050565b60008115159050919050565b6000819050611c0e82611dd9565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611c4882611c00565b9050919050565b82818337600083830152505050565b60005b83811015611c7c578082015181840152602081019050611c61565b83811115611c8b576000848401525b50505050565b60006002820490506001821680611ca957607f821691505b60208210811415611cbd57611cbc611d6a565b5b50919050565b6000611cce82611c33565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611d0157611d00611d0c565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60038110611dea57611de9611d3b565b5b50565b611df681611be2565b8114611e0157600080fd5b50565b60038110611e1157600080fd5b5056fea2646970667358221220a6b98a5882383cc99fc48fcdfd72538d504f28e32efb34c8c1710045754913c964736f6c63430008000033",
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

// GetUserPublicKey is a free data retrieval call binding the contract method 0x11231fe0.
//
// Solidity: function getUserPublicKey(address userAddr) view returns(string)
func (_Verify *VerifyCaller) GetUserPublicKey(opts *bind.CallOpts, userAddr common.Address) (string, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "getUserPublicKey", userAddr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUserPublicKey is a free data retrieval call binding the contract method 0x11231fe0.
//
// Solidity: function getUserPublicKey(address userAddr) view returns(string)
func (_Verify *VerifySession) GetUserPublicKey(userAddr common.Address) (string, error) {
	return _Verify.Contract.GetUserPublicKey(&_Verify.CallOpts, userAddr)
}

// GetUserPublicKey is a free data retrieval call binding the contract method 0x11231fe0.
//
// Solidity: function getUserPublicKey(address userAddr) view returns(string)
func (_Verify *VerifyCallerSession) GetUserPublicKey(userAddr common.Address) (string, error) {
	return _Verify.Contract.GetUserPublicKey(&_Verify.CallOpts, userAddr)
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

// RegisterAsUser is a paid mutator transaction binding the contract method 0x515ef106.
//
// Solidity: function registerAsUser(string _publicKey) returns()
func (_Verify *VerifyTransactor) RegisterAsUser(opts *bind.TransactOpts, _publicKey string) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "registerAsUser", _publicKey)
}

// RegisterAsUser is a paid mutator transaction binding the contract method 0x515ef106.
//
// Solidity: function registerAsUser(string _publicKey) returns()
func (_Verify *VerifySession) RegisterAsUser(_publicKey string) (*types.Transaction, error) {
	return _Verify.Contract.RegisterAsUser(&_Verify.TransactOpts, _publicKey)
}

// RegisterAsUser is a paid mutator transaction binding the contract method 0x515ef106.
//
// Solidity: function registerAsUser(string _publicKey) returns()
func (_Verify *VerifyTransactorSession) RegisterAsUser(_publicKey string) (*types.Transaction, error) {
	return _Verify.Contract.RegisterAsUser(&_Verify.TransactOpts, _publicKey)
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

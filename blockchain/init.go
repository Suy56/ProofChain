package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ClientConnection struct {
	Client    *ethclient.Client
	ClientURL string
	TxOpts    *bind.TransactOpts
	CallOpts  *bind.CallOpts
	ChainId   *big.Int
	ctx       context.Context
}

func (conn *ClientConnection) New(privateKey string) error {
	conn.ClientURL = "http://localhost:7545"
	log.Println("Nooo")

	client, err := ethclient.Dial(conn.ClientURL)
	log.Println("YOO")
	if err != nil {
		log.Fatal("Error connecting to the client : ", err)
	}
	conn.Client=client
	conn.ctx=context.Background()
	conn.ChainId,err=client.ChainID(conn.ctx)
	if err!=nil{
		log.Fatal("Error getting chainID : ",err)
	}
	log.Print("moooo")

	err=conn.setTxOpts(privateKey[2:])
	log.Print("goooo")

	if err!=nil{
		return err
	}
	return nil
}

func (conn *ClientConnection) setTxOpts(privateKeyString string) error {
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		return err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	if !ok {
		return err
	}
	nonce, err := conn.Client.PendingNonceAt(conn.ctx, fromAddress)
	if err != nil {
		log.Fatal("nonce error : ", err)
	}

	fmt.Println("nonce : ", nonce)

	gasPrice, err := conn.Client.SuggestGasPrice(conn.ctx)
	if err != nil {
		log.Fatal("Error getting gas price : ", err)
	}
	chainID, err := conn.Client.ChainID(conn.ctx)
	if err != nil {
		log.Fatal("Error getting chainID : ", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal("Error new key transaction : ", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice
	conn.TxOpts = auth
	return nil
}

package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/Suy56/ProofChain/blockchain"
	"github.com/Suy56/ProofChain/keyUtils"
	"github.com/Suy56/ProofChain/nodeData"
	"github.com/Suy56/ProofChain/wallet"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)

// App struct
type App struct {
	ctx      		context.Context
	conn		  	*blockchain.ClientConnection
	instance 		*blockchain.ContractVerifyOperations
	keys			*keyUtils.ECKeys
	ipfs 			*nodeData.IPFSManager
	user			string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		conn: 		&blockchain.ClientConnection{},
		instance: 	&blockchain.ContractVerifyOperations{},
		keys: 		&keyUtils.ECKeys{},
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (app *App) startup(ctx context.Context) {
	app.ctx = ctx
}

func(app *App)tryDecrypt(encryptedIPFS string, publicAddr common.Address)string{
		pub,err:=app.instance.Instance.GetUserPublicKey(app.conn.CallOpts,publicAddr);if err!=nil{
			fmt.Println("error : ",err)
			return ""
		}
		fmt.Println("user pub : ",pub)
		if err:=app.keys.SetMultiSigKey(pub);err!=nil{
			fmt.Println("error : ",err)
			return ""
		}
		sec,err:=app.keys.GenerateSecret();if err!=nil{
			fmt.Println("error : ",err)
			return ""
		}
		ipfs,err:=keyUtils.DecryptIPFSHash(sec,[]byte(encryptedIPFS));if err!=nil{
			fmt.Println("error : ",err,ipfs)
			return ""
		}
		return ipfs
}

func (app *App) Login(username string, password string) (error) {
	envMap,err:=godotenv.Read(".env");if err!=nil{
		log.Println("Error loading .env while loggin in : ",err)
		return fmt.Errorf("error loading configuration environment; files may've been corrupted.\n Please report the issue to the admin")
	}
	var wg sync.WaitGroup
	errchan:=make(chan error)
	if err:=godotenv.Load(); err!=nil{
		return err
	}
	wg.Add(1)
	go func(){
		defer wg.Done()
		app.keys.OnLogin(username,password,errchan)
	}()
	go func(){
		wg.Wait()
		defer close(errchan)
	}()
	for err:=range errchan{
		if err!=nil{
			log.Println("Error retrieving user's keys  : ",err)
			return fmt.Errorf("error retrieving account. Make sure the credentials are correct")
		}
	}

	privateKey,err:=wallet.RetriveAccount(username,password);if err!=nil{
		log.Println("Error retrieving user's wallet in : ",err)
		return fmt.Errorf("error retrieving account. Make sure the credentials are correct")

	}

	if err:=blockchain.Init(app.conn,app.instance,privateKey,envMap["CONTRACT_ADDR"]);err!=nil{
		log.Println("Error connecting to the blockchain : ",err)
		return fmt.Errorf("error connecting to the smart contract")
	}
	return nil
}


func (app *App)Register(privateKeyString, institute, password string) error {
	envMap,err:=godotenv.Read(".env");if err!=nil{
		log.Println("Error reading .env for contractAddr ",err)
		return err
	}
	fmt.Println("Private key length : ",len(privateKeyString))
	if len(privateKeyString)<64{
		fmt.Println("private key error")
		return fmt.Errorf("invalid private key")
	}
	if err:=blockchain.Init(app.conn,app.instance,privateKeyString[2:],envMap["CONTRACT_ADDR"]);err!=nil{
		return err
	}

	var wg sync.WaitGroup
	errchan:=make(chan error)
	publicKeychan:=make(chan string)	

	wg.Add(2)
	go func() {
		defer wg.Done()
		app.keys.OnRegister(institute,password,publicKeychan,errchan)
		
	}()
	go func(){
		defer wg.Done()
		wallet.NewWallet(privateKeyString[2:],institute,password,errchan)
	}()
	go func(){
		wg.Wait()
		close(publicKeychan)
		close(errchan)
	}()

	for{
		select{
		case pub,ok:=<-publicKeychan:
			if !ok{
				continue
			}else{
				if err:=app.instance.RegisterInstitution(app.conn.TxOpts,pub,institute);err!=nil{
					return err
				}
			}
		case err,ok:=<-errchan:
			if err!=nil{
				return err
			}
			if !ok{
				return nil
			}
		}
	}
}

func (app *App)RegisterVerifier(privateKeyString, username, password string) error {
	envMap,err:=godotenv.Read(".env");if err!=nil{
		log.Println("Error reading .env for contractAddr ",err)
		return err
	}
	fmt.Println("Private key length : ",len(privateKeyString))
	if len(privateKeyString)<64{
		fmt.Println("private key error")
		return fmt.Errorf("invalid private key")
	}
	if err:=blockchain.Init(app.conn,app.instance,privateKeyString[2:],envMap["CONTRACT_ADDR"]);err!=nil{
		return err
	}

	var wg sync.WaitGroup
	errchan:=make(chan error)
	publicKeychan:=make(chan string)	

	wg.Add(2)
	go func() {
		defer wg.Done()
		app.keys.OnRegister(username,password,publicKeychan,errchan)
		
	}()
	go func(){
		defer wg.Done()
		wallet.NewWallet(privateKeyString[2:],username,password,errchan)
	}()
	go func(){
		wg.Wait()
		close(publicKeychan)
		close(errchan)
	}()

	for{
		select{
		case pub,ok:=<-publicKeychan:
			if !ok{
				continue
			}else{
				if err:=app.instance.RegisterUser(app.conn.TxOpts,pub,username,password);err!=nil{
					return err
				}
			}
		case err,ok:=<-errchan:
			if err!=nil{
				return err
			}
			if !ok{
				return nil
			}
		}
	}
}


func (app *App)GetAllDocs()([]blockchain.VerificationDocument,error){
	docs, err := app.instance.GetDocuments(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	for i:=0;i<len(docs);i++{
		docs[i].IpfsAddress=app.tryDecrypt(docs[i].IpfsAddress,common.HexToAddress(docs[i].Requester))
	}
	return docs,nil
}

func (app *App) GetAcceptedDocs() ([]blockchain.VerificationDocument, error) {
	docs, err := app.instance.GetDocuments(app.conn.CallOpts)
	if err != nil {
		return nil, err
	}
	for i:=0;i<len(docs);i++{
		docs[i].IpfsAddress=app.tryDecrypt(docs[i].IpfsAddress,common.HexToAddress(docs[i].Requester))
	}
	verifiedDocs := blockchain.FilterDocument(docs, func(doc blockchain.VerificationDocument, requester string) bool {
		return doc.Requester == requester && doc.Stats == 0
	}, app.conn.CallOpts.From.Hex())
	fmt.Println("Verified docs : ", verifiedDocs)
	return verifiedDocs, nil
}

func (app *App) GetRejectedDocuments() ([]blockchain.VerificationDocument, error) {
	docs, err := app.instance.GetDocuments(app.conn.CallOpts)
	if err != nil {
		return nil, err
	}
	for i:=0;i<len(docs);i++{
		docs[i].IpfsAddress=app.tryDecrypt(docs[i].IpfsAddress,common.HexToAddress(docs[i].Requester))
	}
	verifiedDocs := blockchain.FilterDocument(docs, func(doc blockchain.VerificationDocument, requester string) bool {
		return doc.Requester == requester && doc.Stats == 1
	}, app.conn.CallOpts.From.Hex())
	fmt.Println("Verified docs : ", verifiedDocs)
	return verifiedDocs, nil

}

func (app *App) GetPendingDocuments() ([]blockchain.VerificationDocument, error) {
	docs, err := app.instance.GetDocuments(app.conn.CallOpts)
	if err != nil {
		return nil, err
	}
	verifiedDocs := blockchain.FilterDocument(docs, func(doc blockchain.VerificationDocument, requester string) bool {
		return doc.Requester == requester && doc.Stats == 3
	}, app.conn.CallOpts.From.Hex())
	fmt.Println("Verified docs : ", verifiedDocs)
		return verifiedDocs, nil
}


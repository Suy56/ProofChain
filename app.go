package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/Suy56/ProofChain/blockchain"
	"github.com/Suy56/ProofChain/keyUtils"
	"github.com/Suy56/ProofChain/nodeData"
	"github.com/Suy56/ProofChain/wallet"
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
	contractAddr	string
	user			string
}
type Login interface{
	OnLogin()
	OnRegister()
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

func (app *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (app *App) Login(username string, password string) (error) {
	var wg sync.WaitGroup
	errchan:=make(chan error)
	if err:=godotenv.Load(); err!=nil{
		return err
	}
	contractAddr:=os.Getenv("CONTRACT_ADDR")

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
			return err
		}
	}

	privateKey,err:=wallet.RetriveAccount(username,password);if err!=nil{
		return err
	}

	if err:=blockchain.Init(app.conn,app.instance,privateKey,contractAddr);err!=nil{
		return err
	}

	return nil
}

func (app *App) LoginUserTest() bool {
	return true
}

func (app *App) LoginVerifierTest() bool {
	return true
}

func (app *App) Register(privateKeyString, username, password string) error {
	if err:=blockchain.Init(app.conn,app.instance,privateKeyString,app.contractAddr);err!=nil{
		return err
	}

	var wg sync.WaitGroup
	errchan:=make(chan error)
	publicKeychan:=make(chan string)	

	wg.Add(2)
	fmt.Println(username,password)
	go func() {
		defer wg.Done()
		app.keys.OnRegister(username,password,publicKeychan,errchan)
		
	}()
	go func(){
		defer wg.Done()
		wallet.NewWallet(privateKeyString,username,password,errchan)
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
			if !ok{
				return nil
			}
			if err!=nil{
				return err
			}
		}
	}
}

func (app *App) GetAcceptedDocs() ([]blockchain.VerificationDocument, error) {
	docs, err := app.instance.GetDocuments(app.conn.CallOpts)
	if err != nil {
		return nil, err
	}
	verifiedDocs := blockchain.FilterDocument(docs, func(doc blockchain.VerificationDocument, requester common.Address) bool {
		return doc.Requester == requester && doc.Stats == 0
	}, app.conn.CallOpts.From)
	fmt.Println("Verified docs : ", verifiedDocs)
	return verifiedDocs, nil
}

func (app *App) GetRejectedDocuments() ([]blockchain.VerificationDocument, error) {
	docs, err := app.instance.GetDocuments(app.conn.CallOpts)
	if err != nil {
		return nil, err
	}
	a:=func(encryptedIPFS string)string{
		pub,err:=app.instance.Instance.GetUserPublicKey(app.conn.CallOpts,app.conn.TxOpts.From);if err!=nil{
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
		return "ipfs"
	}
	for i:=0;i<len(docs);i++{
		docs[i].IpfsAddress=a(docs[i].IpfsAddress)
	}
	verifiedDocs := blockchain.FilterDocument(docs, func(doc blockchain.VerificationDocument, requester common.Address) bool {
		return doc.Requester == requester && doc.Stats == 1
	}, app.conn.CallOpts.From)
	fmt.Println("Verified docs : ", verifiedDocs)
	return verifiedDocs, nil

}

func (app *App) GetPendingDocuments() ([]blockchain.VerificationDocument, error) {
	docs, err := app.instance.GetDocuments(app.conn.CallOpts)
	if err != nil {
		return nil, err
	}
	verifiedDocs := blockchain.FilterDocument(docs, func(doc blockchain.VerificationDocument, requester common.Address) bool {
		return doc.Requester == requester && doc.Stats == 3
	}, app.conn.CallOpts.From)
	fmt.Println("Verified docs : ", verifiedDocs)
		return verifiedDocs, nil
}


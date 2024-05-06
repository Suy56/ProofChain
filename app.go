package main

import (
	"context"
	"fmt"

	"github.com/Suy56/ProofChain/blockchain"
	// "github.com/Suy56/ProofChain/verify"
	"github.com/Suy56/ProofChain/wallet"
	"github.com/ethereum/go-ethereum/common"
)

// App struct
type App struct {
	ctx      		context.Context
	conn		  	*blockchain.ClientConnection
	instance 		*blockchain.ContractVerifyOperations
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		conn: 		&blockchain.ClientConnection{},
		instance: 	&blockchain.ContractVerifyOperations{},
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (app *App) startup(ctx context.Context) {
	app.ctx = ctx
}

// Greet returns a greeting for the given name
func (app *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (app *App) Login(username string, password string) (error) {
	privateKey, err := wallet.RetriveAccount(username, password)
	if err != nil {
		return err
	}
	err = app.conn.New(privateKey)
	if err != nil {
		return  err
	}
	app.instance.Client = app.conn.Client
	err = app.instance.New(blockchain.TempContractAddress, true)
	if err != nil {
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
	go func() error{
		err := wallet.NewWallet(privateKeyString, username, password)
			if err != nil {
			return err
		}
		return nil
	}()
	
	err := app.conn.New(privateKeyString)
	if err != nil {
		return err
	}
	app.instance.Client = app.conn.Client
	err = app.instance.New(blockchain.TempContractAddress, true)
	if err != nil {
		return err
	}
	app.instance.RegisterUser(app.conn.TxOpts,username,password)
	return nil
}

func (app *App) GetAcceptedDocs() ([]blockchain.VerificationDocument, error) {
	docs, err := app.instance.GetDocuments(app.conn.CallOpts)
	if err != nil {
		return nil, err
	}
	verifiedDocs := blockchain.GetDocuments(docs, func(doc blockchain.VerificationDocument, requester common.Address) bool {
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
	verifiedDocs := blockchain.GetDocuments(docs, func(doc blockchain.VerificationDocument, requester common.Address) bool {
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
	verifiedDocs := blockchain.GetDocuments(docs, func(doc blockchain.VerificationDocument, requester common.Address) bool {
		return doc.Requester == requester && doc.Stats == 3
	}, app.conn.CallOpts.From)
	fmt.Println("Verified docs : ", verifiedDocs)
	return verifiedDocs, nil
}

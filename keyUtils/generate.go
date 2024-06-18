package keyUtils

import (
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)
type ECKeys struct{
	Private 	*ecdh.PrivateKey
	Public		*ecdh.PublicKey
	MultiSig	*ecdh.PublicKey 
}

func(k *ECKeys)OnLogin(user,passphrase string)error{
	pemKey,err:=DecryptPrivateKeyFile(user,passphrase);if err!=nil{
		return err
	}
	k.Private,err=GetECDSAPrivateKeyFromPEM(string(pemKey)); if err!=nil{
		return err
	}
	k.Public=k.Private.PublicKey()
	return nil
}


func (k *ECKeys)OnRegister(username,password string)error{
	privECDSA,err:=ecdsa.GenerateKey(elliptic.P256(),rand.Reader);if err!=nil{
		return err
	}
	k.Private,err=privECDSA.ECDH();if err!=nil{
		return err
	}
	k.Public=k.Private.PublicKey()
	pemPrivateKey,err:=k.MarshalECDHPrivateKey();if err!=nil{
		return err
	}
	EncryptPrivateKeyFile(pemPrivateKey,username,password)
	return nil
}

func(k *ECKeys)SetMultiSigKey(multiSigKey string)error{
	var err error
	k.MultiSig,err=GetECDSAPublicKeyFromPEM(multiSigKey); if err!=nil{
		return err
	}
	return nil
}

func(k *ECKeys)GenerateSecret()([]byte,error){
	if k.MultiSig==nil{
		return nil,fmt.Errorf("Multi-Sig-Public-Key not provided")
	}
	secret,err:=k.Private.ECDH(k.MultiSig); if err!=nil{
		return nil,err
	}
	return secret,nil
}

func(k * ECKeys)MarshalECDHPublicKey()(string,error){
	ecdhSKBytes, err := x509.MarshalPKIXPublicKey(k.Public)
	if err != nil {
			return "", fmt.Errorf("failed to marshal public key into PKIX format")
	}
	ecdhSKPEMBlock := pem.EncodeToMemory(
			&pem.Block{
					Type:  "PUBLIC KEY",
					Bytes: ecdhSKBytes,
			},
	)
	fmt.Println(string(ecdhSKPEMBlock))
	return string(ecdhSKPEMBlock), nil
}

func(k *ECKeys)MarshalECDHPrivateKey()(string,error){
	ecdhSKBytes, err := x509.MarshalPKCS8PrivateKey(k.Private)
    if err != nil {
        return "", fmt.Errorf("failed to marshal private key into PKIX format")
    }

    ecdhSKPEMBlock := pem.EncodeToMemory(
        &pem.Block{
            Type:  "PRIVATE KEY",
            Bytes: ecdhSKBytes,
        },
    )
	fmt.Println(string(ecdhSKPEMBlock))
    return string(ecdhSKPEMBlock), nil
}



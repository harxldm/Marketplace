package authorization

import (
	"crypto/rsa"
	"io/ioutil"
	"log"

	"github.com/dgrijalva/jwt-go"
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
)

func LoadFiles(privateFile, publicFile string) error {
	privateBytes, err := ioutil.ReadFile(privateFile)
	if err != nil {
		return err
	}

	publicBytes, err := ioutil.ReadFile(publicFile)
	if err != nil {
		return err
	}

	return parseRSA(privateBytes, publicBytes)
}

func parseRSA(privateBytes, publicBytes []byte) error {
	var err error

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil && signKey == nil {
		log.Fatalf("signKey is nil")
		return err
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return err
	}

	return nil
}

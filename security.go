package main

import (
	jwt "github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"net/http"
)

// location of the files used for signing and verification
const (
	privKeyPath = "keys/app.rsa"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

// keys are held in global variables
// i havn't seen a memory corruption/info leakage in go yet
// but maybe it's a better idea, just to store the public key in ram?
// and load the signKey on every signing request? depends on your usage i guess
var (
	verifyKey, signKey []byte
)

// read the JWT key files before starting http handlers
func initJWT() {
	var err error

	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}

	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
}

// Validate the received security token
// If good, return the UserID
func securityCheck(passedToken string) (int, string) {
	// validate the token

	//log.Println("Security Check:", passedToken)
	token, err := jwt.Parse(passedToken, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return signKey, nil
	})

	// branch out into the possible error from signing
	switch err.(type) {

	case nil: // no error

		if !token.Valid { // but may still be invalid
			log.Println("Invalid Token", passedToken)
			return http.StatusUnauthorized, err.Error()
		}

		log.Printf("Token OK:%+v\n", token.Claims)
		return http.StatusOK, "Token Valid"

	case *jwt.ValidationError: // something was wrong during the validation
		vErr := err.(*jwt.ValidationError)

		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			return http.StatusUnauthorized, err.Error()

		default:
			return http.StatusUnauthorized, "Invalid Token!"
		}

	default: // something else went wrong
		log.Printf("Token parse error: %v\n", err)
		return http.StatusUnauthorized, "Invalid Token!"
	}
}

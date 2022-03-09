package main

import (
	crand "crypto/rand"
	"encoding/hex"
	"fmt"
	eddsabn254 "github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"
	"github.com/consensys/gnark-crypto/signature"
)

type Keys struct {
	Sk string `json:"sk"'`
	Pk string `json:"pk"'`
}

type Signature struct {
	Sig string `json:"sig"`
}

func generateKeys() *Keys {
	signature.Register(signature.EDDSA_BN254, eddsabn254.GenerateKeyInterfaces)
	privateKey, _ := signature.EDDSA_BN254.New(crand.Reader)
	sk := hex.EncodeToString(privateKey.Bytes())
	pk := hex.EncodeToString(privateKey.Public().Bytes())
	keys := &Keys{
		sk,
		pk,
	}
	return keys
}

func signMessage(skBytes []byte, msg []byte) string {
	signature.Register(signature.EDDSA_BN254, eddsabn254.GenerateKeyInterfaces)
	privateKey, _ := signature.EDDSA_BN254.New(crand.Reader)
	privateKey.SetBytes(skBytes)
	sig, _ := privateKey.Sign(msg, hFunc)

	return hex.EncodeToString(sig)
}

func verifySig(sig string, pkBytes []byte, msg []byte) bool {
	signature.Register(signature.EDDSA_BN254, eddsabn254.GenerateKeyInterfaces)
	privateKey, _ := signature.EDDSA_BN254.New(crand.Reader)
	publicKey := privateKey.Public()
	publicKey.SetBytes(pkBytes)

	sigbin, _ := hex.DecodeString(sig)
	isValid, _ := publicKey.Verify(sigbin, msg, hFunc)
	if !isValid {
		fmt.Println("\n1. invalid signature")
		return false
	} else {
		fmt.Println("\n1. valid signature")
		return true
	}
}

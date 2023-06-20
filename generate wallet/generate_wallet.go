package main

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	fmt.Printf("here is the private key: %v \n", hexutil.Encode(privateKeyBytes))

	publicKey := privateKey.Public()

	publicKeyHe, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyHe)
	fmt.Printf("\n\n here's your public key :%s ", hexutil.Encode(publicKeyBytes))

	address := crypto.PubkeyToAddress(*publicKeyHe).Hex()

	fmt.Printf("here is the generated address: %v\n", address)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes)
	fmt.Printf("here is the hash of the wallet: %v\n", hexutil.Encode(hash.Sum(nil)))
}
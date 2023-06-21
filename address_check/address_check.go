package main

import (
	"context"
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Check if Address is Valid
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	fmt.Printf("the address is valid address? %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d"))
	fmt.Printf("the address is valid address? %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d"))

	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err!= nil {
		panic(err)
	}

	// Check if Address is an Account or a Smart Contract
	address1 := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	bytecode1, err := client.CodeAt(context.Background(), address1, nil)
	if err!= nil {
		panic(err)
	}
	isCorrect1 := len(bytecode1) > 0
	fmt.Printf("the address is a smart contract? %v\n", isCorrect1)

	address2 := common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
	bytecode2, err := client.CodeAt(context.Background(), address2, nil)
	if err!= nil {
		panic(err)
	}
	isCorrect2 := len(bytecode2) > 0
	fmt.Printf("the address is a smart contract? %v\n", isCorrect2)
}
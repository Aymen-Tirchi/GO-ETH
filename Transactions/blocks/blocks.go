package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		panic(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err!=nil{
		panic(err)
	}

	blockNumber := new(big.Int)
	blockNumber.SetString(header.Number.String(), 10)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err!= nil {
		panic(err)
	}
	fmt.Printf("the block number is : %v\n", block.Number().Uint64())
	fmt.Printf("the timestamp of the block is : %v\n", block.Time())
	fmt.Printf("the block difficulty is : %v\n", block.Difficulty().Uint64())
	fmt.Printf("the block hash is : %v\n", block.Hash().Hex())
	fmt.Printf("the block transactions : %v\n", len(block.Transactions()))
}
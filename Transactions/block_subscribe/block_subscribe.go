package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://eth-mainnet.g.alchemy.com/v2/-lUiiErMbM2Y2Qs_Q8187mwupUVH8qSC")
	if err != nil {
		panic(err)
	}
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		panic(err)
	}
	for {
		select {
		case err := <-sub.Err():
			panic(err)
		case header := <-headers:
			fmt.Printf("the hash of the new block : %v \n", header.Hash().Hex())
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err!=nil{
				panic(err)
			}
			fmt.Printf("the block hash %v\n", block.Hash().Hex())
			fmt.Printf("the block number %v\n", block.Number().Uint64())
			fmt.Printf("the block timestamp %v\n", block.Time())
			fmt.Printf("the block nonce %v\n", block.Nonce())
			fmt.Printf("the block transactions %v\n", block.Transactions())
		}
	}

}

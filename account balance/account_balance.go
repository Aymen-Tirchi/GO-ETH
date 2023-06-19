package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err!= nil{
		panic(err)
	}
	fmt.Printf("the balance of the account: %v is: %v wei\n", account, balance)

	blockNumber := big.NewInt(5555454)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		panic(err)
	}
	fmt.Printf("here is the balance and the block number of the account %v:\n", account)
	fmt.Printf("balance %v wei:\n", balanceAt)
	fmt.Printf("block number %v:\n", blockNumber)

	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Printf("the balance converted to eth : %v ETH\n", ethValue)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	
	fmt.Printf("here is the pending balance : %v wei\n", pendingBalance)
}
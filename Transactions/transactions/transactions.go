package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		panic(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	blockNumber := new(big.Int)
	blockNumber.SetString(header.Number.String(), 10)

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		panic(err)
	}

	contractAddress := common.HexToAddress("0x1413cD9B3026213B9C4e77621E17023E53235e23")

	for _, tx := range block.Transactions() {
		fmt.Println("here is the transactions hash:", tx.Hash().Hex())
		fmt.Println("here is the transactions value:", tx.Value().String())
		fmt.Println("here is the transactions gas:", tx.Gas())
		fmt.Println("here is the transactions gas price:", tx.GasPrice().Uint64())
		fmt.Println("here is the transactions nonce:", tx.Nonce())
		fmt.Println("here is the transactions data:", tx.Data())
		fmt.Println("here is the transactions to address:", tx.To().Hex())

		if tx.To() == nil {
			fmt.Println("Skipping contract creation transaction")
			continue
		}

		if len(tx.Data()) == 0 {
			fmt.Println("Skipping empty data transaction")
			continue
		}

		if tx.To().Hex() == contractAddress.Hex() {
			fmt.Println("Skipping contract address transaction")
			continue
		}

		sender, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
		if err != nil {
			fmt.Println("Error getting sender address:", err)
			continue
		}

		fmt.Println("here is the transaction from address:", sender.Hex())

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			fmt.Println("Error getting transaction receipt:", err)
			continue
		}

		fmt.Println("here is the transaction receipt:", receipt)
	}

	txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Is the transaction hash %v pending? %v\n", tx.Hash().Hex(), isPending)
}

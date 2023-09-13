package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    client, err := ethclient.Dial("https://cloudflare-eth.com")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("successfully connected")
    _ = client
}
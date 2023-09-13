package main

import (
	"fmt"

	store "github.com/aymen-tirchi/go-ethereum/cmd/smart_contracts/smart_contract_read/build"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main () {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		panic(err)
	}
	address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	instance, err := store.NewStore(address, client)
	if err != nil {
		panic(err)
	}
	version, err := instance.Version(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("version : ", version)
}
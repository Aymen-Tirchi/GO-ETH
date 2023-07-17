package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	store "github.com/aymen-tirchi/go-ethereum/smart_contracts/smart_contract_load/build"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		panic(err)
	}
	address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")

	instance, err := store.NewStore(address, client)
	if err != nil {
		panic(err)
	}
	_ = instance
}
package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		panic(err)
	}
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	value := big.NewInt(0)
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Printf("the method ID is %v\n", methodID)
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Printf("then padded address : %v\n", hexutil.Encode(paddedAddress))
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Printf("the padded amount : %v\n", hexutil.Encode(paddedAmount))
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		panic(err)
	}
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}
	signer := types.NewEIP155Signer(chainID)
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		panic(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("transaction sent %v\n", signedTx.Hash().Hex())
}

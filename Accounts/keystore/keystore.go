package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func createKs(){
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		panic("Error creating account")
	}
	fmt.Printf("here is your account: %v\n", account.Address.Hex())
}
func importks() {
	file := "./tmp/UTC--2023-06-20T13-10-00.338095600Z--e240d509744c275ceea87b1f34315af81a854978"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := os.ReadFile(file)
	if err != nil {
		panic("Error reading file")
	}
	password := "secret"

	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		panic("Error importing wallet: " + err.Error())
	}

	fmt.Printf("Here is the account address: %v\n", account.Address.Hex())

	if err := os.Remove(file); err != nil {
		panic(err)
	}
}
func main() {
	var choice string
	fmt.Println("do u want to create a keystore ? then type create and if u want to import a keystore then type import")

	fmt.Scan(&choice)

	switch choice {
	case "create" : createKs()
	case "import" : importks()
	default : fmt.Println("you entered invalid choice")
	}
}
## Smart Contract Compilation & ABI:
To compile the Store.sol smart contract and generate the ABI and EVM bytecode, run the following command:

```bash
solc --bin Store.sol -o build
```
Make sure you have the solc compiler installed before running this command.

To install the abigen tool, which is used to generate the Go contract file, execute the following command:
```bash
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```
Once the installation is complete, you can use abigen to generate the Go contract file by running the following command:
```bash
abigen --bin=./build/Store.bin --abi=./build/Store.abi --pkg=store --out=Store.go
```
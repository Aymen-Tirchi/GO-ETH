## Smart Contracts:

Make sure you have the solc compiler installed

- Smart Contract Compilation & ABI

To compile a Solidity contract and generate the EVM bytecode and ABI file, use the following command:

```bash
solc --bin Store.sol -o build
```

To generate the Go contract file using the abigen tool, run the following command:

```bash
abigen --bin=build/Store.bin --abi=build/Store.abi --pkg=store --out=build/Store.go
```

- Deploying a Smart Contract

```bash
solc --bin Store.sol -o build
```

```bash
abigen --bin=build/Store.bin --abi=build/Store.abi --pkg=store --out=build/Store.go
```
To deploy the contract using `contract_deploy.go`, you can run the following command:

```bash
go run contract_deploy.go
```

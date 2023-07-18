## Smart Contracts:

Make sure you have the solc compiler installed

- Smart Contract Compilation & ABI

```bash
cd smart_contract_compile
```

To compile a Solidity contract and generate the EVM bytecode and ABI file, use the following command:

```bash
solc --abi --bin store.sol -o build
```

To generate the Go contract file using the abigen tool, run the following command:

```bash
abigen --bin=build/store.bin --abi=build/store.abi --pkg=store --out=build/Store.go
```

- Deploying a Smart Contract

```bash
cd smart_contract_deploy
```

```bash
solc --abi --bin store.sol -o build
```

```bash
abigen --bin=build/store.bin --abi=build/store.abi --pkg=store --out=build/Store.go
```
To deploy the contract using `contract_deploy.go`, you can run the following command:

```bash
go run contract_deploy.go
```

- Loading a Smart Contract

```bash
cd smart_contract_load
```

```bash
solc --abi --bin store.sol -o build
```

```bash
abigen --bin=build/store.bin --abi=build/store.abi --pkg=store --out=build/Store.go
```

```bash
go run contract_load.go
```

- Querying a Smart Contract

```bash
cd smart_contract_read
```

```bash
solc --abi --bin store.sol -o build
```

```bash
abigen --bin=build/store.bin --abi=build/store.abi --pkg=store --out=build/Store.go
```

```bash
go run contract_read.go
```
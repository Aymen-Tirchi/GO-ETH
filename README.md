# GO-ETH

GO-ETH is a project for practicing Go programming with the go-ethereum library. It includes various functionalities related to Ethereum account management and interaction.

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/Aymen-Tirchi/GO-ETH.git && cd GO-ETH
```

2. Install the required dependencies:

```bash
go mod download
```

3. Run the desired functionality by executing the corresponding Go file:

- Setting up the Client:

```bash
go run ./client/client.go
```

- Accounts

```bash
go run ./address/address.go
```

- Account Balances

```bash
go run ./account_balance/account_balance.go
```

- Generating New Wallets

```bash
go run ./generate_wallet/generate_wallet.go
```

- Keystores

```bash
go run ./keystore/keystore.go
```

- Address Check

```bash
go run ./address_check/address_check.go
```

- Querying Blocks

```bash
go run ./blocks/blocks.go
```

- Querying Transactions

```bash
go run ./transactions/transactions.go
```

- Transferring ETH

```bash
go run ./transfer_eth/transfer_eth.go
```

- Transferring Tokens (ERC-20)

```bash
go run ./transfer_tokens/transfer_tokens.go
```

- Subscribing to New Blocks

```bash
go run block_subscribe/block_subscribe.go
```

- Create Raw Transaction

```bash
go run transaction_raw_create/transaction_raw_create.go
```

- Send Raw Transaction:

```bash
go run transaction_raw_sendreate/transaction_raw_sendreate.go
```

- Smart Contract Compilation & ABI:
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

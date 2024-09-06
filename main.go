package main

import (
	"fmt"

	"github.com/alexdconf/backend-chain-golang/blockchain"
	"github.com/alexdconf/backend-chain-golang/wallet"
)

func main() {
	// Initialize the blockchain
	bc := blockchain.InitializeBlockchain()

	// Create two wallets
	wallet1 := wallet.NewWallet()
	wallet2 := wallet.NewWallet()

	// Display wallet addresses
	fmt.Println("Wallet 1 Address:", wallet1.Address())
	fmt.Println("Wallet 2 Address:", wallet2.Address())

	// Create a transaction from wallet1 to wallet2
	tx := blockchain.CreateTransaction(wallet1, wallet2, 10.0)

	// Add the transaction to the blockchain
	if bc.AddTransaction(tx) {
		fmt.Println("Transaction added successfully.")
	} else {
		fmt.Println("Failed to add transaction.")
	}

	// Mine a block
	bc.MinePendingTransactions(wallet1.Address())

	// Display the blockchain
	fmt.Printf("Blockchain: %+v\n", bc.Blocks)
}

package tests

import (
    "testing"
    "github.com/alexdconf/IOchain/blockchain"
)

func TestBlockchainInitialization(t *testing.T) {
    bc := blockchain.InitializeBlockchain()
    if len(bc.Blocks) != 1 {
        t.Errorf("Expected blockchain to have 1 block, got %d", len(bc.Blocks))
    }
}

func TestBlockCreation(t *testing.T) {
    bc := blockchain.InitializeBlockchain()

    transactions := []blockchain.Transaction{
        {"Alice", "Bob", 50.0},
    }

    newBlock := blockchain.CreateBlock(bc.Blocks[0], transactions)
    if newBlock.Index != 1 {
        t.Errorf("Expected new block index to be 1, got %d", newBlock.Index)
    }

    if newBlock.Transactions[0].Sender != "Alice" {
        t.Errorf("Expected sender to be Alice, got %s", newBlock.Transactions[0].Sender)
    }
}

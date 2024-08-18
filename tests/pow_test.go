package tests

import (
    "testing"
    "github.com/alexdconf/IOchain/blockchain"
)

func TestMining(t *testing.T) {
    genesisBlock := blockchain.CreateGenesisBlock()
    genesisBlock.MineBlock()

    if genesisBlock.Hash[:blockchain.difficulty] != strings.Repeat("0", blockchain.difficulty) {
        t.Errorf("Block mining failed, got hash: %s", genesisBlock.Hash)
    }
}

func TestIsBlockValid(t *testing.T) {
    genesisBlock := blockchain.CreateGenesisBlock()
    genesisBlock.MineBlock()

    newBlock := blockchain.CreateBlock(genesisBlock, []blockchain.Transaction{
        {"Alice", "Bob", 50.0},
    })

    if !blockchain.IsBlockValid(newBlock, genesisBlock) {
        t.Errorf("Block should be valid")
    }
}

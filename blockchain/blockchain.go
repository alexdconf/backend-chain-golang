package blockchain

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "github.com/alexdconf/IOchain/wallet"
)

type Transaction struct {
    Sender    string
    Receiver  string
    Amount    float64
    Signature string
}

type Block struct {
    Index        int
    Timestamp    string
    Transactions []Transaction
    PreviousHash string
    Hash         string
    Nonce        int
}

type Blockchain struct {
    Blocks             []Block
    PendingTransactions []Transaction
}

// CreateTransaction creates a new transaction, signs it with the sender's wallet, and returns it.
func CreateTransaction(senderWallet *wallet.Wallet, receiver string, amount float64) Transaction {
    tx := Transaction{
        Sender:    senderWallet.Address(),
        Receiver:  receiver,
        Amount:    amount,
    }
    tx.SignTransaction(senderWallet)
    return tx
}

func (tx *Transaction) HashTransaction() [32]byte {
    data := tx.Sender + tx.Receiver + fmt.Sprintf("%f", tx.Amount)
    return sha256.Sum256([]byte(data))
}

func (tx *Transaction) SignTransaction(w *wallet.Wallet) {
    txHash := tx.HashTransaction()
    r, s, err := wallet.SignData(w.PrivateKey, txHash[:])
    if err != nil {
        fmt.Println("Failed to sign transaction:", err)
    }
    signature := append(r.Bytes(), s.Bytes()...)
    tx.Signature = hex.EncodeToString(signature)
}

func (tx *Transaction) VerifySignature() bool {
    r, s, err := wallet.DecodeSignature(tx.Signature)
    if err != nil {
        return false
    }
    txHash := tx.HashTransaction()
    return wallet.VerifySignature(tx.Sender, r, s, txHash[:])
}

func (bc *Blockchain) AddTransaction(tx Transaction) bool {
    if tx.VerifySignature() {
        bc.PendingTransactions = append(bc.PendingTransactions, tx)
        return true
    }
    return false
}

func InitializeBlockchain() *Blockchain {
    genesisBlock := createGenesisBlock()
    return &Blockchain{[]Block{genesisBlock}, []Transaction{}}
}

func createGenesisBlock() Block {
    return Block{
        Index:        0,
        Timestamp:    "2023-01-01",
        Transactions: []Transaction{},
        PreviousHash: "0",
        Hash:         calculateHash(0, "2023-01-01", []Transaction{}, "0", 0),
    }
}

func (bc *Blockchain) MinePendingTransactions(minerAddress string) {
    block := Block{
        Index:        len(bc.Blocks),
        Timestamp:    "2023-01-01",
        Transactions: bc.PendingTransactions,
        PreviousHash: bc.Blocks[len(bc.Blocks)-1].Hash,
    }
    block.Hash = block.mineBlock()
    bc.Blocks = append(bc.Blocks, block)

    // Reset the list of pending transactions
    bc.PendingTransactions = []Transaction{
        {Sender: "System", Receiver: minerAddress, Amount: 1}, // Mining reward
    }
}

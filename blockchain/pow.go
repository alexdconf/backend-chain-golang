package blockchain

import (
    "crypto/sha256"
    "encoding/hex"
    "strconv"
    "strings"
    "fmt"
)

const difficulty = 3

func calculateHash(index int, timestamp string, transactions []Transaction, previousHash string, nonce int) string {
    record := strconv.Itoa(index) + timestamp + fmt.Sprintf("%v", transactions) + previousHash + strconv.Itoa(nonce)
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}

func (block *Block) mineBlock() string {
    target := strings.Repeat("0", difficulty)
    for {
        block.Hash = calculateHash(block.Index, block.Timestamp, block.Transactions, block.PreviousHash, block.Nonce)
        if block.Hash[:difficulty] == target {
            return block.Hash
        }
        block.Nonce++
    }
}

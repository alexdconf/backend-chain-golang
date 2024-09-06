package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func calculateHash(index int, timestamp time.Time, transactions []Transaction, previousHash string, nonce int) string {
	record := strconv.Itoa(index) + timestamp.String() + fmt.Sprintf("%v", transactions) + previousHash + strconv.Itoa(nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// Adjust difficulty based on the time it took to mine the last 'adjustmentInterval' blocks
func (blockchain *Blockchain) adjustDifficulty() {
	const targetBlockTime int = 10    // Target block time in seconds
	const adjustmentInterval int = 10 // Adjust difficulty every 10 blocks

	if len(blockchain.Blocks)%adjustmentInterval != 0 {
		return
	}

	latestBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	firstBlockInInterval := blockchain.Blocks[len(blockchain.Blocks)-adjustmentInterval]

	actualTimeForInterval := int(math.Floor(latestBlock.Timestamp.Sub(firstBlockInInterval.Timestamp).Seconds()))
	expectedTimeForInterval := targetBlockTime * adjustmentInterval

	// Adjust the difficulty based on the time taken to mine the last interval
	if actualTimeForInterval < expectedTimeForInterval {
		blockchain.Difficulty++
	} else if actualTimeForInterval > expectedTimeForInterval {
		if blockchain.Difficulty > 1 {
			blockchain.Difficulty--
		}
	}
}

// Mine the block with dynamic difficulty adjustment
func (block *Block) mineBlock(blockchain *Blockchain) string {
	blockchain.adjustDifficulty()
	target := strings.Repeat("0", blockchain.Difficulty)
	for {
		block.Hash = calculateHash(block.Index, block.Timestamp, block.Transactions, block.PreviousHash, block.Nonce)
		if block.Hash[:blockchain.Difficulty] == target {
			return block.Hash
		}
		block.Nonce++
	}
}

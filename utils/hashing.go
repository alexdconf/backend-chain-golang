package utils

import (
    "crypto/sha256"
    "encoding/hex"
)

func CalculateHash(input string) string {
    h := sha256.New()
    h.Write([]byte(input))
    return hex.EncodeToString(h.Sum(nil))
}

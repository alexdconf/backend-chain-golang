package tests

import (
    "testing"
    "github.com/alexdconf/backend-chain-golang/wallet"
)

func TestWalletCreation(t *testing.T) {
    w := wallet.NewWallet()
    if w.Address() == "" {
        t.Errorf("Expected wallet address to be non-empty")
    }
}

func TestTransactionSigning(t *testing.T) {
    w := wallet.NewWallet()
    data := []byte("test data")
    r, s, err := wallet.SignData(w.PrivateKey, data)
    if err != nil {
        t.Errorf("Failed to sign data: %v", err)
    }

    if !wallet.VerifySignature(w.Address(), r, s, data) {
        t.Errorf("Failed to verify signature")
    }
}

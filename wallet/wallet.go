package wallet

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "math/big"
    "encoding/hex"
    "log"
)

type Wallet struct {
    PrivateKey *ecdsa.PrivateKey
    PublicKey  *ecdsa.PublicKey
}

func NewWallet() *Wallet {
    private, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        log.Fatal(err)
    }
    public := &private.PublicKey
    return &Wallet{PrivateKey: private, PublicKey: public}
}

func (w *Wallet) Address() string {
    pubKeyBytes := append(w.PublicKey.X.Bytes(), w.PublicKey.Y.Bytes()...)
    hash := sha256.Sum256(pubKeyBytes)
    return hex.EncodeToString(hash[:])
}

func SignData(privateKey *ecdsa.PrivateKey, data []byte) (r, s *big.Int, err error) {
    r, s, err = ecdsa.Sign(rand.Reader, privateKey, data)
    return r, s, err
}

func VerifySignature(address string, r, s *big.Int, data []byte) bool {
    x, y := elliptic.Unmarshal(elliptic.P256(), []byte(address))
    if x == nil {
        return false
    }
    pubKey := ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}
    return ecdsa.Verify(&pubKey, data, r, s)
}

func DecodeSignature(signature string) (r, s *big.Int, err error) {
    sigBytes, err := hex.DecodeString(signature)
    if err != nil {
        return nil, nil, err
    }
    r = new(big.Int).SetBytes(sigBytes[:len(sigBytes)/2])
    s = new(big.Int).SetBytes(sigBytes[len(sigBytes)/2:])
    return r, s, nil
}

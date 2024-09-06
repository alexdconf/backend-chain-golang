package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"log"
	"math/big"
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
	return hex.EncodeToString(pubKeyBytes)
}

func SignData(privateKey *ecdsa.PrivateKey, data []byte) (r, s *big.Int, err error) {
	r, s, err = ecdsa.Sign(rand.Reader, privateKey, data)
	return r, s, err
}

func VerifySignature(address string, r, s *big.Int, data []byte) bool {
	// Check for the correct public key length (for P-256 curve, it's 64 bytes)
	if len(address) != 128 {
		return false
	}

	// Extract the X and Y coordinates from the public key bytes
	tmp_x, err := hex.DecodeString(address[:len(address)/2])
	if err != nil {

	}
	x := new(big.Int).SetBytes(tmp_x)

	tmp_y, err := hex.DecodeString(address[len(address)/2:])
	if err != nil {

	}
	y := new(big.Int).SetBytes(tmp_y)

	// Create the ECDSA public key
	pubKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	// Verify the signature using the ECDSA public key
	ret := ecdsa.Verify(&pubKey, data, r, s)
	return ret
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

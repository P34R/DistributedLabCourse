package cryptography

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
)

func SignData(message string, keys *KeyPair) []byte {
	ms := sha256.Sum256([]byte(message))
	sig, err := ecdsa.SignASN1(rand.Reader, keys.keys, ms[:])
	if err != nil {
		panic(err)
	}
	return sig
}

func VerifySignature(sig []byte, publicKey *ecdsa.PublicKey, message string) bool {
	ms := sha256.Sum256([]byte(message))
	return ecdsa.VerifyASN1(publicKey, ms[:], sig)
}

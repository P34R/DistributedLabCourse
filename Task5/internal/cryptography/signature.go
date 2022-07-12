package cryptography

import (
	"crypto/ecdsa"
	"crypto/rand"
)

func SignData(message string, keys *KeyPair) []byte {
	ms := ToSHA256(message)
	sig, err := ecdsa.SignASN1(rand.Reader, keys.keys, ms)
	if err != nil {
		panic(err)
	}
	return sig
}

func VerifySignature(sig []byte, publicKey *ecdsa.PublicKey, message string) bool {
	ms := ToSHA256(message)
	return ecdsa.VerifyASN1(publicKey, ms, sig)
}

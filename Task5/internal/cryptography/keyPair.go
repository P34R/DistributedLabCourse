package cryptography

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

type KeyPair struct {
	keys *ecdsa.PrivateKey //there is public key inside private. You can get public key with PublicKey() func
}

func GenKeyPair() *KeyPair {
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return &KeyPair{
		keys: key,
	}
}

// PublicKey Same as p.keys.Public()
func (pair *KeyPair) PublicKey() *ecdsa.PublicKey {
	return &pair.keys.PublicKey
}
func (pair *KeyPair) ToString() string {
	return "private: " + pair.keys.D.String() + " X: " + pair.keys.PublicKey.X.String() + " Y: " + pair.keys.PublicKey.Y.String()
}

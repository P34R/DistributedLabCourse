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

func (p *KeyPair) PublicKey() *ecdsa.PublicKey {
	return &p.keys.PublicKey
}
func (p *KeyPair) ToString() string {
	return "private: " + p.keys.D.String() + " X: " + p.keys.PublicKey.X.String() + " Y: " + p.keys.PublicKey.Y.String()
}

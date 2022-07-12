package cryptography

import "crypto/sha256"

func ToSHA256(message string) []byte {
	ret := sha256.Sum256([]byte(message))
	return ret[:]
}

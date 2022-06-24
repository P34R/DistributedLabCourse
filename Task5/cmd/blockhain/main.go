package main

import (
	"fmt"
	"github.com/P34R/DistributedLabCourse/Task5/internal/cryptography"
)

func main() {
	a := cryptography.GenKeyPair()
	message := "Test message"
	b := cryptography.SignData(message, a)
	fmt.Println(cryptography.VerifySignature(b, a.PublicKey(), message))
}

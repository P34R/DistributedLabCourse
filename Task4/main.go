package main

import (
	"fmt"
	"github.com/P34R/DistributedLabCourse/Task4/RSA"
)

// Implementation inside RSA Folder
func main() {
	s := "I love Distributed Lab"
	e, d, n := RSA.KeyGen()
	r := RSA.SendMessage(s, e, n)
	k := RSA.GetMessage(r, d, n)
	fmt.Println("start message:   ", s)
	fmt.Println("ciphertext ([]*Int):    ", r)
	fmt.Println("recovered original message:    ", k)
}

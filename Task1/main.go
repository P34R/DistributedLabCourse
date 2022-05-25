package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

// Task 1
func task1() {
	two := big.NewInt(2)
	counter := big.NewInt(8)
	counter.Exp(two, counter, nil)
	new(big.Int).Set(counter)

	for i := big.NewInt(8); i.Int64() <= 4096; i.Mul(i, big.NewInt(2)) {
		fmt.Println(i.String(), ": ", counter.String())
		randKey := generate_random_key(counter)
		fmt.Println("random key (hex): ", randKey.Text(16))

		//Task 3
		//fmt.Println("brute force done in: ", brute_force(randKey),"ms")

		fmt.Println()
		counter.Exp(counter, two, nil)
	}
}

// Task 2
func generate_random_key(max *big.Int) *big.Int {
	m := new(big.Int).Set(max) //copying, so max won't be changed
	m.Sub(m, big.NewInt(1))    //2^n - 1
	n, err := rand.Int(rand.Reader, m)
	if err != nil {
		panic(err)
	}
	return n
}

// Task 3
func brute_force(key *big.Int) int64 {
	start := time.Now().UnixMilli()
	startKey := big.NewInt(0)
	for key.Cmp(startKey) != 0 {
		startKey.Add(startKey, big.NewInt(1))
	}
	return time.Now().UnixMilli() - start
}

func main() {
	task1()
}

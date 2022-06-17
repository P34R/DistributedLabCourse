package RSA

import (
	"crypto/rand"
	"math/big"
)

const (
	BITS_SIZE = 512
)

func KeyGen() (*big.Int, *big.Int, *big.Int) {
	n := big.NewInt(1)
	fi := big.NewInt(1)

	p, _ := rand.Prime(rand.Reader, BITS_SIZE)
	q, _ := rand.Prime(rand.Reader, BITS_SIZE)

	n.Mul(p, q)
	fi.Add(fi, n) // n+1
	fi.Sub(fi, p)
	fi.Sub(fi, q) // = (p-1)(q-1)
	e, x, _ := getCoprime(fi)

	d := big.NewInt(0)
	d.Mod(x, fi)
	d.Add(d, fi)
	d.Mod(d, fi)

	return e, d, n
}

// Actually returns prime, but whatever?
func getCoprime(a *big.Int) (*big.Int, *big.Int, *big.Int) {
	r, _ := rand.Prime(rand.Reader, BITS_SIZE/2)
	x := big.NewInt(0)
	y := big.NewInt(0)
	temp := big.NewInt(1)
	for true {
		temp.GCD(x, y, r, a)
		if temp.Cmp(big.NewInt(1)) == 0 {
			break
		} else {
			r, _ = rand.Prime(rand.Reader, BITS_SIZE/2)
		}
	}

	return r, x, y
}

func SendMessage(s string, e, n *big.Int) []*big.Int {
	ret := make([]*big.Int, len(s))
	for i, ele := range s {
		ret[i] = big.NewInt(0)
		ret[i].SetInt64(int64(ele))
		ret[i].Exp(ret[i], e, n)

	}
	return ret
}
func GetMessage(s []*big.Int, d, n *big.Int) string {

	ret := ""
	temp := big.NewInt(0)
	for i := range s {
		temp.Exp(s[i], d, n)
		ret += string(rune(temp.Int64()))
	}
	return ret
}

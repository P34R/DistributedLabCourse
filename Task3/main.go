package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
)

//FourXOR probably useless function, cause we have XOR, but it was implemented earlier
func FourXOR(a, b, c, d []byte) (ret []byte) {
	for i := range a {
		ret = append(ret, a[i]^b[i]^c[i]^d[i])
	}
	return ret
}

func XOR(a, b []byte) (ret []byte) {
	for i := range a {
		ret = append(ret, a[i]^b[i])
	}
	//	fmt.Println(a,"XOR ",b," = ",ret)
	return ret
}

func NOT(a []byte) (ret []byte) {
	for i := range a {
		ret = append(ret, 255-a[i])
	}
	//	fmt.Println(a," NOT ",ret)
	return ret
}

func AND(a, b []byte) (ret []byte) {
	size := 0
	if len(a) > len(b) {
		size = len(b)
	} else {
		size = len(a)
	}
	for i := 0; i < size; i++ {
		ret = append(ret, a[i]&b[i])
	}
	//	fmt.Println(a,"AND ",b," = ",ret)
	return ret
}

func OR(a, b []byte) (ret []byte) {
	for i := range a {
		ret = append(ret, a[i]|b[i])
	}
	//	fmt.Println(a,"OR ",b," = ",ret)
	return ret
}

func SHA1MessagePreprocessing(message string, ml int) []byte {
	chunks := bytes.NewBufferString(message).Bytes()
	chunks = append(chunks, 128)
	for len(chunks)%64 != 56 {
		chunks = append(chunks, 0)
	}
	ml_chunk := make([]byte, 8)
	binary.BigEndian.PutUint64(ml_chunk, uint64(ml*8))
	chunks = append(chunks, ml_chunk...)
	return chunks
}

func plus(a, b []byte) (ret []byte) {
	aUint32 := binary.BigEndian.Uint32(a)
	bUint32 := binary.BigEndian.Uint32(b)

	ret = make([]byte, len(a))
	binary.BigEndian.PutUint32(ret, aUint32+bUint32)

	return ret
}

func ZerosSwitcher(input byte) (str string) {
	if input == 0 {
		str += "00000000"
	} else if input == 1 {
		str += "0000000"
	} else if input >= 2 && input < 4 {
		str += "000000"
	} else if input >= 4 && input < 8 {
		str += "00000"
	} else if input >= 8 && input < 16 {
		str += "0000"
	} else if input >= 16 && input < 32 {
		str += "000"
	} else if input >= 32 && input < 64 {
		str += "00"
	} else if input >= 64 && input < 128 {
		str += "0"
	}
	return str
}

func StrToByteArrUint32(s string) []byte {
	ret := make([]byte, 4)
	counter := 0
	for i := range s {
		if i > 0 && i%8 == 0 {
			counter++
		}
		if s[i] == '1' {
			ret[counter] += byte(math.Pow(2, float64(7-i%8)))
		}
	}

	return ret
}

//leftRotate Naive implementation of Left Rotate
func leftRotate(input []byte, num int) []byte {
	str := ""
	str += ZerosSwitcher(input[0])
	k := 0
	for input[k] == 0 {
		if k == len(input)-1 {
			return StrToByteArrUint32(str)
		}
		str += ZerosSwitcher(input[k+1])
		k++
	}
	str += strconv.FormatUint(uint64(binary.BigEndian.Uint32(input)), 2)
	//	fmt.Println("A            ",str)
	str = str[num:] + str[0:num]
	//	fmt.Println("B            ",str)
	return StrToByteArrUint32(str)
}

func GetSHA1Hash(message string) string {
	h0, _ := hex.DecodeString("67452301")
	h1, _ := hex.DecodeString("efcdab89")
	h2, _ := hex.DecodeString("98badcfe")
	h3, _ := hex.DecodeString("10325476")
	h4, _ := hex.DecodeString("c3d2e1f0") //variables initialization
	ml := len(message)

	bits512 := 64
	processed := SHA1MessagePreprocessing(message, ml)
	num_of_blocks := len(processed) / bits512
	chunks := make([][]byte, num_of_blocks)
	for i := 0; i < num_of_blocks; i++ {
		chunks[i] = processed[bits512*(i) : bits512*(i+1)]
	}
	for ch := range chunks {
		if ch == 0 {
		}
		word := make([][]byte, 80)
		for i := 0; i < 16; i++ {
			word[i] = append(word[i], chunks[ch][4*i:4*(i+1)]...)
		}
		for i := 16; i < 80; i++ {
			word[i] = leftRotate(FourXOR(word[i-3], word[i-8], word[i-14], word[i-16]), 1)

		}
		//		fmt.Println("\t\t\ta\t\t\t", "b\t\t\t", "c\t\t\t", "d\t\t\t", "e\t\t\t")
		a := h0
		b := h1
		c := h2
		d := h3
		e := h4
		k := make([]byte, 8)
		var f []byte
		for i := 0; i < 80; i++ {
			if i >= 0 && i <= 19 {
				f = OR(AND(b, c), AND(NOT(b), d))
				k, _ = hex.DecodeString("5a827999")
			} else if i >= 20 && i <= 39 {
				f = XOR(XOR(b, c), d)
				k, _ = hex.DecodeString("6ed9eba1")
			} else if i >= 40 && i <= 59 {
				f = OR(OR(AND(b, c), AND(b, d)), AND(c, d))
				k, _ = hex.DecodeString("8f1bbcdc")
			} else if i >= 60 && i <= 79 {
				f = XOR(XOR(b, c), d)
				k, _ = hex.DecodeString("ca62c1d6")
			}
			temp := plus(plus(plus(plus(leftRotate(a, 5), f), e), k), word[i])
			e = d
			d = c
			c = leftRotate(b, 30)
			b = a
			a = temp
			//			fmt.Println("i =", i, "\t", hex.EncodeToString(a), "\t", hex.EncodeToString(b), "\t", hex.EncodeToString(c), "\t", hex.EncodeToString(d), "\t", hex.EncodeToString(e))
		}
		h0 = plus(h0, a)
		h1 = plus(h1, b)
		h2 = plus(h2, c)
		h3 = plus(h3, d)
		h4 = plus(h4, e)
	}
	hash := hex.EncodeToString(h0) + hex.EncodeToString(h1) + hex.EncodeToString(h2) + hex.EncodeToString(h3) + hex.EncodeToString(h4)
	return hash
}

func main() {

	fmt.Println(GetSHA1Hash("This is a text message!"))
}

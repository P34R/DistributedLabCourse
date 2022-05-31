package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
)

//As mentioned in task, hex number may be given in string format
func hexToLittle(input string) *big.Int {
	inputStr := removeHexPart(input)
	inputBytes, _ := hex.DecodeString(inputStr)
	inputBytes = reverseBytes(inputBytes)
	bigInt := big.Int{}
	bigInt.SetBytes(inputBytes)
	return &bigInt
}

//As mentioned in task, hex number may be given in string format
func hexToBig(input string) *big.Int {
	inputStr := removeHexPart(input)
	inputBytes, _ := hex.DecodeString(inputStr)
	bigInt := big.Int{}
	bigInt.SetBytes(inputBytes)
	return &bigInt
}

//littleEndStrToHex assumes that input is Little-endian decimal number in string format (example "255")
func littleEndStrToHex(input string) string {
	bigInt := big.Int{}
	bigInt.SetString(input, 10)
	ret := bigInt.Text(16)
	ret = "0x" + ret
	return ret
}

//bigEndStrToHex assumes that input is Big-endian decimal number in string format (example "1829371892391823798178923789127893")
func bigEndStrToHex(input string) string {
	bigInt := big.Int{}
	bigInt.SetString(input, 10)
	ret := bigInt.Text(16)
	ret = "0x" + ret
	return ret
}

//littleEndIntToHex assumes that input is Little-endian decimal number in big.Int format (function is the same as bigEndIntToHexNaive)
func littleEndIntToHex(input *big.Int) string {
	ret := input.Text(16)
	return "0x" + ret
}

//bigEndIntToHex assumes that input is Big-endian decimal number in big.Int format (function is the same as littleEndIntToHexNaive)
func bigEndIntToHex(input *big.Int) string {
	ret := input.Text(16)
	return "0x" + ret
}

//bigToHex assumes that input is byte array in classic (Big-endian) format
func bigToHex(input []byte) string {
	bigInt := big.Int{}
	bigInt.SetBytes(input)
	ret := bigInt.Text(16)
	return "0x" + ret
}

//littleToHex assumes that input is byte array in classic (Big-endian) format
func littleToHex(input []byte) string {
	bigInt := big.Int{}
	inputWork := reverseBytes(input)
	bigInt.SetBytes(inputWork)
	ret := bigInt.Text(16)
	return "0x" + ret
}

func reverseBytes(input []byte) []byte {
	ret := input

	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return ret
}

func removeHexPart(input string) string {
	if len(input) > 2 && input[0] == '0' && input[1] == 'x' {
		return input[2:]
	}
	return input
}

func countBytes(input []byte) int {
	return len(input)
}

func printAll(value string) (*big.Int, *big.Int) {
	little := hexToLittle(value)
	Big := hexToBig(value)
	fmt.Println("Value: ", value)
	fmt.Println("Number of bytes: ", countBytes(Big.Bytes()))
	fmt.Println("Little-Endian: ", little)
	fmt.Println("Big-Endian: ", Big)
	return little, Big
}

func main() {
	//	value1:="0xff00000000000000000000000000000000000000000000000000000000000000"
	//	value2:="0xaaaa000000000000000000000000000000000000000000000000000000000000"
	//	value3:="0xFFFFFFFF"
	value4 := "0xF000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"

	a, b := printAll(value4)
	fmt.Println()
	fmt.Println(littleEndIntToHex(a))
	fmt.Println(bigEndIntToHex(b))
	//	printAll(value2); fmt.Println()
	//	printAll(value3); fmt.Println()
	//	printAll(value4); fmt.Println()
}

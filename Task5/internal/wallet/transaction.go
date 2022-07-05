package wallet

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type Transaction struct {
	TxID            string
	SetOfOperations []*Operation
	Nonce           uint64
}

func CreateTransaction(setOfOperations []*Operation, nonce uint64) *Transaction {

	var id string
	for _, ele := range setOfOperations {
		id += ele.sender.AccountID + ele.receiver.AccountID + strconv.Itoa(int(ele.amount)) + string(ele.signature)
	}
	id += strconv.Itoa(int(nonce))

	txID := sha256.Sum256([]byte(id))
	return &Transaction{
		TxID:            hex.EncodeToString(txID[:]),
		SetOfOperations: setOfOperations,
		Nonce:           nonce,
	}
}

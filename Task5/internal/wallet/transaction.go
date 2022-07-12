package wallet

import (
	"encoding/hex"
	"github.com/P34R/DistributedLabCourse/Task5/internal/cryptography"
	"strconv"
)

type Transaction struct {
	TxID            string
	SetOfOperations []*Operation
	Nonce           uint64
}

func CreateTransaction(setOfOperations []*Operation, nonce uint64) *Transaction {
	// Assumed that operations are sorted by timestamp
	var id string
	for _, ele := range setOfOperations {
		id += ele.Sender.AccountID + ele.Receiver.AccountID + strconv.Itoa(int(ele.Amount)) + string(ele.Signature)
	}
	id += strconv.Itoa(int(nonce))
	txID := cryptography.ToSHA256(id)
	return &Transaction{
		TxID:            hex.EncodeToString(txID),
		SetOfOperations: setOfOperations,
		Nonce:           nonce,
	}
}

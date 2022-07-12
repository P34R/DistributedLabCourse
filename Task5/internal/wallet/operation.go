package wallet

import (
	"github.com/P34R/DistributedLabCourse/Task5/internal/cryptography"
	"strconv"
)

type Operation struct {
	timestamp string // Used to prevent multiple use of the same operation, this may be changed in future
	Sender    *Account
	Receiver  *Account
	Amount    uint64
	Signature []byte
}

func CreateOperation(timestamp string, sender, receiver *Account, amount uint64, signature []byte) *Operation {
	return &Operation{
		timestamp: timestamp,
		Sender:    sender,
		Receiver:  receiver,
		Amount:    amount,
		Signature: signature,
	}
}

func VerifyOperation(operation *Operation) bool {
	if operation.Sender == nil || operation.Receiver == nil {
		return false
	}
	if operation.Amount <= operation.Sender.GetBalance() {
		message := operation.timestamp + operation.Sender.AccountID + operation.Receiver.AccountID + strconv.Itoa(int(operation.Amount))
		for _, ele := range operation.Sender.wallet {
			if cryptography.VerifySignature(operation.Signature, ele.PublicKey(), message) {
				return true
			}
		}
	}
	return false
}

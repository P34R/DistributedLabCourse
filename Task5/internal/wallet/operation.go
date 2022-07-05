package wallet

import (
	"github.com/P34R/DistributedLabCourse/Task5/internal/cryptography"
	"strconv"
)

type Operation struct {
	timestamp string // Used to prevent multiple use of the same operation, this may be changed in future
	sender    *Account
	receiver  *Account
	amount    uint64
	signature []byte
}

func CreateOperation(timestamp string, sender, receiver *Account, amount uint64, signature []byte) *Operation {
	return &Operation{
		timestamp: timestamp,
		sender:    sender,
		receiver:  receiver,
		amount:    amount,
		signature: signature,
	}
}

func VerifyOperation(operation *Operation) bool {
	if operation.sender == nil || operation.receiver == nil {
		return false
	}
	if operation.amount <= operation.sender.GetBalance() {
		message := operation.timestamp + operation.sender.AccountID + operation.receiver.AccountID + strconv.Itoa(int(operation.amount))
		for _, ele := range operation.sender.wallet {
			if cryptography.VerifySignature(operation.signature, ele.PublicKey(), message) {
				return true
			}
		}
	}
	return false
}

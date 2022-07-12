package wallet

import (
	"encoding/hex"
	"fmt"
	"github.com/P34R/DistributedLabCourse/Task5/internal/cryptography"
	"strconv"
	"time"
)

type Account struct {
	AccountID string
	wallet    []*cryptography.KeyPair
	balance   uint64
}

func GenAccount() *Account {

	var wallet []*cryptography.KeyPair
	wallet = append(wallet, cryptography.GenKeyPair())
	str := wallet[0].PublicKey().X.Text(16) + wallet[0].PublicKey().Y.Text(16)
	id := cryptography.ToSHA256(str)
	return &Account{
		AccountID: hex.EncodeToString(id), //SHA256 hash of first key pair generated in account
		wallet:    wallet,
		balance:   0,
	}
}

func (acc *Account) AddKeyPairToWallet(keyPair *cryptography.KeyPair) {
	acc.wallet = append(acc.wallet, keyPair)
}

func (acc *Account) UpdateBalance(balance uint64) {
	//Uint is used, so the "if" statement is probably useless
	//if balance >= 0 {
	acc.balance = balance
	//}
}

func (acc *Account) GetBalance() uint64 {
	return acc.balance
}

// PrintBalance does not print '\n' at the end
func (acc *Account) PrintBalance() {
	fmt.Print(acc.balance)
}

func (acc *Account) SignData(message string, index int) []byte {
	return cryptography.SignData(message, acc.wallet[index])
}

func (acc *Account) CreateOperation(receiver *Account, amount uint64, index int) *Operation {
	if index > len(acc.wallet) {
		panic("Index out of range")
	}
	timestamp := time.Now().UnixMicro()
	message :=
		strconv.Itoa(int(timestamp)) +
			acc.AccountID +
			receiver.AccountID +
			strconv.Itoa(int(amount))
	return CreateOperation(strconv.Itoa(int(timestamp)), acc, receiver, amount, acc.SignData(message, index))
}

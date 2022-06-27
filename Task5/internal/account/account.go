package account

import (
	"crypto/sha256"
	"fmt"
	"github.com/P34R/DistributedLabCourse/Task5/internal/cryptography"
)

type Account struct {
	AccountID []byte
	Wallet    []*cryptography.KeyPair
	balance   uint64
}

func GenAccount() *Account {

	var wallet []*cryptography.KeyPair
	wallet = append(wallet, cryptography.GenKeyPair())
	str := wallet[0].PublicKey().X.Text(16) + wallet[0].PublicKey().Y.Text(16)
	id := sha256.Sum256([]byte(str))
	return &Account{
		AccountID: id[:], //SHA256 hash of first key pair generated in account (may be changed in future)
		Wallet:    wallet,
		balance:   0,
	}
}

func (a *Account) AddKeyPairToWallet(keyPair *cryptography.KeyPair) {
	a.Wallet = append(a.Wallet, keyPair)
}

func (a *Account) UpdateBalance(balance uint64) {
	if balance >= 0 {
		a.balance = balance
	}
}

func (a *Account) GetBalance() uint64 {
	return a.balance
}

// PrintBalance does not print '\n'
func (a *Account) PrintBalance() {
	fmt.Print(a.balance)
}

func (a *Account) SignData(message string, index int) []byte {
	return cryptography.SignData(message, a.Wallet[index])
}

// CreateOperation currently unavailable cause no "Operation" Class
func (a *Account) CreateOperation(recepient Account, amount int, index int) {
	// will be added later
}

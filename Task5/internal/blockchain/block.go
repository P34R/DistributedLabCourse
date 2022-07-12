package blockchain

import (
	"encoding/hex"
	"github.com/P34R/DistributedLabCourse/Task5/internal/cryptography"
	"github.com/P34R/DistributedLabCourse/Task5/internal/wallet"
)

type Block struct {
	BlockID           string
	PrevHash          string
	SetOfTransactions []*wallet.Transaction
}

func CreateBlock(SetOfTransactions []*wallet.Transaction, PrevHash string) *Block {
	id := PrevHash
	for _, ele := range SetOfTransactions {
		id += ele.TxID //taking only TxID cause it consists of other tx fields (hashed ofc)
	}
	return &Block{
		BlockID:           hex.EncodeToString(cryptography.ToSHA256(id)),
		PrevHash:          PrevHash,
		SetOfTransactions: SetOfTransactions,
	}
}

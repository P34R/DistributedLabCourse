package blockchain

import (
	"fmt"
	"github.com/P34R/DistributedLabCourse/Task5/internal/wallet"
)

type Blockchain struct {
	CoinDatabase map[string]uint64
	BlockHistory []*Block
	TxDatabase   []*wallet.Transaction
	FaucetCoins  uint64
}

func InitBlockchain() *Blockchain {
	genesis := CreateBlock(nil, "0000000000000000000000000000000000000000000000000000000000000000") //64 hex characters
	var history []*Block
	var TxDB []*wallet.Transaction
	history = append(history, genesis)
	CoinDB := make(map[string]uint64)
	TestCoins := uint64(50000)
	return &Blockchain{
		CoinDatabase: CoinDB,
		BlockHistory: history,
		TxDatabase:   TxDB,
		FaucetCoins:  TestCoins,
	}

}

func (chain *Blockchain) GetTokenFromFaucet(Account *wallet.Account, amount uint64) {
	if chain.FaucetCoins >= amount {
		Account.UpdateBalance(amount + Account.GetBalance()) // adds coins
		chain.FaucetCoins -= amount
		chain.CoinDatabase[Account.AccountID] = Account.GetBalance()
	}

}
func (chain *Blockchain) validateBlock(block *Block) bool {

	// Making coins DB copy. Will be adding here TXes in future to check double-spending
	CoinDatabaseCopy := make(map[string]uint64)
	for i, ele := range chain.CoinDatabase {
		CoinDatabaseCopy[i] = ele
	}

	// Checking that our block is unique (new)
	if block.PrevHash != chain.BlockHistory[len(chain.BlockHistory)-1].BlockID {
		return false
	}
	fail := false
	for _, tx := range block.SetOfTransactions {

		// Checking that our transactions are unique (new)
		for _, DBtx := range chain.TxDatabase {
			if tx.TxID == DBtx.TxID {
				fail = true
				break
			}
		}

		if fail {
			break
		}

		// Operations validation (and conflicting transactions too)
		for _, oper := range tx.SetOfOperations {
			if !wallet.VerifyOperation(oper) {
				fail = true
				break
			}
			//applying operation
			CoinDatabaseCopy[oper.Sender.AccountID] = oper.Sender.GetBalance() - oper.Amount
			CoinDatabaseCopy[oper.Receiver.AccountID] = oper.Receiver.GetBalance() + oper.Amount

			// updating (real) balances, but that will be updated once more in the end from the original Coins DB
			oper.Sender.UpdateBalance(CoinDatabaseCopy[oper.Sender.AccountID])
			oper.Receiver.UpdateBalance(CoinDatabaseCopy[oper.Receiver.AccountID])
		}

		if fail {
			break
		}

	}
	if fail {
		// I'm still wondering what would be the best way to implement this (Reupdating to real balances after failure)
		for _, tx := range block.SetOfTransactions {
			for _, oper := range tx.SetOfOperations {
				oper.Sender.UpdateBalance(chain.CoinDatabase[oper.Sender.AccountID])
				oper.Receiver.UpdateBalance(chain.CoinDatabase[oper.Receiver.AccountID])
			}
		}
		return false
	}
	for _, tx := range block.SetOfTransactions {
		chain.TxDatabase = append(chain.TxDatabase, tx)
	}
	chain.CoinDatabase = CoinDatabaseCopy
	chain.BlockHistory = append(chain.BlockHistory, block)
	return true
}

func (chain *Blockchain) ShowCoinDatabase() {
	fmt.Println(chain.CoinDatabase)
}

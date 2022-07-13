package main

import (
	"fmt"
	"github.com/P34R/DistributedLabCourse/Task5/internal/blockchain"
	"github.com/P34R/DistributedLabCourse/Task5/internal/wallet"
)

func main() {

	bc := blockchain.InitBlockchain()

	acc1 := wallet.GenAccount()
	acc2 := wallet.GenAccount()
	acc3 := wallet.GenAccount()
	bc.GetTokenFromFaucet(acc1, 1000)
	bc.GetTokenFromFaucet(acc2, 10)
	bc.GetTokenFromFaucet(acc3, 0) //just adding account to CoinsDatabase (not needed actually, will be added after first TX)

	opp1 := acc1.CreateOperation(acc2, 50, 0)
	opp2 := acc2.CreateOperation(acc3, 30, 0)

	opps := make([]*wallet.Operation, 2)
	opps[0] = opp1
	opps[1] = opp2

	tx := wallet.CreateTransaction(opps, 15)

	opp2_1 := acc3.CreateOperation(acc1, 5, 0)
	opp2_2 := acc2.CreateOperation(acc3, 20, 0)

	opps2 := make([]*wallet.Operation, 2)
	opps2[0] = opp2_1
	opps2[1] = opp2_2

	tx2 := wallet.CreateTransaction(opps2, 6)

	txes := make([]*wallet.Transaction, 2)
	txes[0] = tx
	txes[1] = tx2
	block := blockchain.CreateBlock(txes, bc.GetLastBlockHash())
	bc.ShowCoinDatabase()
	if bc.ValidateBlock(block) {
		fmt.Println("block is OK")

	}
	bc.ShowCoinDatabase()
}

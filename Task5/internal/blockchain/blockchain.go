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

func (chain *Blockchain) ValidateBlock(block *Block) bool {

	/*
		Функция validateBlock должна содержать набор следующих проверок:
		Проверка что блок содержит ссылку на последний актуальный блок в истории;
		Проверка что транзакции в блоке еще не были добавлены в историю;
		Проверка того что блок не содержит конфликтующих транзакций.
		Проверка каждой операции в транзакции:
		Проверка подписи;
		Проверка того что операция платит не больше монет чем хранится на балансе аккаунта отправителя.

	*/

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
	for i, v := range chain.CoinDatabase {
		fmt.Println(i, v)
	}
}

func (chain *Blockchain) GetLastBlockHash() string {
	return chain.BlockHistory[len(chain.BlockHistory)-1].BlockID
}

/*

Объекты:

coinDatabase
таблица отражающая текущее состояние балансов в системе. В качестве ключа используется идентификатор аккаунт, в качестве значения - баланс пользователя.
blockHistory
массив хранящий все блоки добавленные в историю.
txDatabase
массив хранящий все транзакции в истории. Будет использоваться для более быстрого доступа при проверке существования транзакции в истории (защита от дублирования).
faucetCoins
целочисленное значение определяющее количество монет доступных в кране для тестирования.

Методы:

initBlockchain()
функция позволяющая проинициализировать блокчейн. Под капотом происходит создание генезис блока и добавление его в историю.
getTokenFromFaucet()
функция позволяющая получить тестовые монеты с крана. Обновляет состояние coinDatabase и баланса аккаунта, который вызвал метод.
validateBlock()
функция позволяющая выполнить проверку и добавить блок в историю.
getTokenFromFaucet()
функция позволяющая получить текущее состояния аккаунтов и балансов.

Опциональные методы (упростят жизнь при тестировании):

toString()
функция позволяющая сформировать строку с объектов блокчейна. Возвращает объект класса String.
printKeyPair()
функция для вывода объектов блокчейна. Не возвращает ничего.

Замечания и предложения:
Функция validateBlock должна содержать набор следующих проверок:
Проверка что блок содержит ссылку на последний актуальный блок в истории;
Проверка что транзакции в блоке еще не были добавлены в историю;
Проверка того что блок не содержит конфликтующих транзакций.
Проверка каждой операции в транзакции:
Проверка подписи;
Проверка того что операция платит не больше монет чем хранится на балансе аккаунта отправителя.



*/

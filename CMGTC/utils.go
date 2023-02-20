package CMGTC

import (
	"fmt"
	"golang-cmgt-coin-miner/mod10"
	"strconv"
	"time"
)

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Println()
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

// TryNonce tries different nonce's and validates generated hash
func TryNonce(block string, condition Callback, n int) (int, string) {
	hash := mod10.HashPayload(block + strconv.Itoa(n))
	if condition(hash, "0000") {
		return n, hash
	}
	return TryNonce(block, condition, n+1)
}

// isValidHash Callback function to check whether the hash is valid
func isValidHash(str string, value string) bool {
	firstFour := str[:4]
	return firstFour == value
}

// getTransactionString Creates string from transaction
func getTransactionString(transaction struct {
	from      string
	to        string
	amount    int
	timestamp int64
}) string {
	return transaction.from + transaction.to + strconv.Itoa(transaction.amount) + strconv.FormatInt(transaction.timestamp, 10)
}

// makeTransactionFromBlockData Converts BlockchainTransactionFields type to Transaction
func makeTransactionFromBlockData(transaction BlockchainTransactionFields) Transaction {
	resp := Transaction{}
	resp.from = transaction.Cfrom()
	resp.to = transaction.Cto()
	resp.amount = transaction.Camount()
	resp.timestamp = transaction.Ctimestamp()

	return resp
}

// makeTransactionFromBlock Converts BlockTransactionFields type to Transaction
func makeTransactionFromBlock(transaction BlockTransactionFields) Transaction {
	resp := Transaction{}
	resp.from = transaction.from()
	resp.to = transaction.to()
	resp.amount = transaction.amount()
	resp.timestamp = transaction.timestamp()

	return resp
}

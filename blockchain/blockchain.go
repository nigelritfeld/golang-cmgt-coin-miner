package blockchain

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"golang-cmgt-coin-miner/mod10"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

// GetLatestBlock Sends request to blockchain for new Block
// If the received block is open it returns the block
// Else the program sleeps until the countdown has finished and restarts
func GetLatestBlock(url string) BlockResponse {
	client := req.C()
	resp, err := client.R().Get(url)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var block BlockResponse
	if err := json.Unmarshal(body, &block); err != nil {
		// Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	//If block is open for mining we return the block
	if block.Open {
		return block
	} else {
		//If block not open, wait till countdown is finished and restart
		seconds := block.Countdown / 1000
		fmt.Printf("[LOOKING FOR BLOCK]Block %s is currently pending, waiting %d seconds before retreiving next block", block.Blockchain.Hash, seconds)
		fmt.Println()
		time.Sleep(time.Duration(seconds))
		return GetLatestBlock(url)
	}
}

// isValidHash Callback function to check whether the hash is valid
func isValidHash(str string, value string) bool {
	firstFour := str[:4]
	return firstFour == value
}

// FindNonce Tries different nonce's till a valid hash is generated
func FindNonce(block string) (int, string) {
	nonce, hash := TryNonce(block, isValidHash, 0)
	return nonce, hash
}

// TryNonce tries different nonce's and validates generated hash
func TryNonce(block string, condition Callback, n int) (int, string) {
	hash := mod10.HashPayload(block + strconv.Itoa(n))
	if condition(hash, "0000") {
		return n, hash
	}
	return TryNonce(block, isValidHash, n+1)
}

// RequestPayout Send request to blockchain for reward
func RequestPayout(nonce int, user string) {
	type Result struct {
		Data string `json:"data"`
	}
	client := req.C().DevMode()
	var result Result
	type Payload struct {
		Nonce string `json:"nonce"`
		User  string `json:"user"`
	}

	resp, err := client.R().
		SetBody(&Payload{Nonce: strconv.Itoa(nonce), User: user}).
		SetSuccessResult(&result).
		Post("https://programmeren9.cmgt.hr.nl:8000/api/blockchain")

	fmt.Println("response")
	fmt.Println(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
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

// GetTransactionListFromBlock Creates string with transactions
func GetTransactionListFromBlock(block BlockResponse) string {
	transactionList := ""
	for _, data := range block.Transactions {
		transactionList = transactionList + getTransactionString(makeTransactionFromBlock(data))
	}
	fmt.Println("TRANSACTIONLIST")
	fmt.Println(transactionList)
	return transactionList
}

// GetDataListFromBlock Creates string with transactions
func GetDataListFromBlock(block BlockResponse) string {
	transactionList := ""
	for _, data := range block.Blockchain.Data {
		transactionList = transactionList + getTransactionString(makeTransactionFromBlockData(data))
	}
	return transactionList
}

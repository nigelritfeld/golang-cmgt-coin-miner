package CMGTC

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
		fmt.Printf("[Found block]: Block ID #%s", block.Blockchain.ID)
		fmt.Println()
		return block
	} else {
		//If block not open, wait till countdown is finished and restart
		seconds := block.Countdown / 1000
		fmt.Printf("[BLOCKCHAIN PENDING]Block %s is currently pending, waiting %d seconds before retreiving next block", block.Blockchain.Hash, seconds)
		fmt.Println()
		time.Sleep(time.Duration(seconds))
		return GetLatestBlock(url)
	}
}

// MineBlock Starts mining process
func MineBlock(block BlockResponse) (int, string) {
	defer Timer("Mining CMGT Coin")()
	LastBlockString := block.Blockchain.Hash + GetDataListFromBlock(block) + strconv.FormatInt(block.Blockchain.Timestamp, 10) + block.Blockchain.Nonce
	lastBlockHash := mod10.HashPayload(LastBlockString)
	fmt.Printf("[Mining block]: Block hash - %s", lastBlockHash)
	fmt.Println()
	transactionsList := ""
	for _, transaction := range block.Transactions {
		trans := transaction.From + transaction.To + strconv.Itoa(transaction.Amount) + strconv.FormatInt(transaction.Timestamp, 10)
		transactionsList = transactionsList + trans
	}
	blockTimeStamp := block.Timestamp
	newBlockString := lastBlockHash + GetTransactionListFromBlock(block) + strconv.FormatInt(blockTimeStamp, 10)

	return FindNonce(newBlockString)
}

// FindNonce Tries different nonce's till a valid hash is generated
func FindNonce(block string) (int, string) {
	nonce, hash := TryNonce(block, isValidHash, 0)
	fmt.Printf("[Found nonce]: new block hash - %s", hash)
	fmt.Println()
	return nonce, hash
}

// RequestPayout Send request to CMGTC for reward
func RequestPayout(nonce int, user string) {
	type Result struct {
		Data string `json:"data"`
	}
	client := req.C()
	var result Result
	type Payload struct {
		Nonce string `json:"nonce"`
		User  string `json:"user"`
	}

	resp, err := client.R().
		SetBody(&Payload{Nonce: strconv.Itoa(nonce), User: user}).
		SetSuccessResult(&result).
		Post("https://programmeren9.cmgt.hr.nl:8000/api/blockchain")
	fmt.Printf("[Requested award]: %s", resp.String())
	if err != nil {
		log.Fatal(err)
	}
}

// GetTransactionListFromBlock Creates string with transactions
func GetTransactionListFromBlock(block BlockResponse) string {
	transactionList := ""
	for _, data := range block.Transactions {
		transactionList = transactionList + getTransactionString(makeTransactionFromBlock(data))
	}
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

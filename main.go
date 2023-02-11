package main

import (
	"fmt"
	"golang-cmgt-coin-miner/blockchain"
	"golang-cmgt-coin-miner/mod10"
	"strconv"
)

func MineBlock(block blockchain.BlockResponse) (int, string) {

	LastBlockString := block.Blockchain.Hash + blockchain.GetDataListFromBlock(block) + strconv.FormatInt(block.Blockchain.Timestamp, 10) + block.Blockchain.Nonce
	fmt.Println("LastBlock String")
	fmt.Println()
	fmt.Println(LastBlockString)
	lastBlockHash := mod10.HashPayload(LastBlockString)
	transactionsList := ""
	for _, transaction := range block.Transactions {
		trans := transaction.From + transaction.To + strconv.Itoa(transaction.Amount) + strconv.FormatInt(transaction.Timestamp, 10)
		transactionsList = transactionsList + trans
	}
	blockTimeStamp := block.Timestamp
	newBlockString := lastBlockHash + blockchain.GetTransactionListFromBlock(block) + strconv.FormatInt(blockTimeStamp, 10)

	return blockchain.FindNonce(newBlockString)

}

func main() {
	fmt.Println("Started mining..")
	//Get the latest block
	block := blockchain.GetLatestBlock("https://programmeren9.cmgt.hr.nl:8000/api/blockchain/next")
	// Start mining process
	nonce, _ := MineBlock(block)
	// Request payout
	blockchain.RequestPayout(nonce, "Nigel 1004416")
}

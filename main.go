package main

import (
	"fmt"
	"golang-cmgt-coin-miner/mod10"
	"golang-cmgt-coin-miner/network"
	"strconv"
)

func main() {
	fmt.Println("Started program..")
	//hash := mod10.HashPayload("test9oke")
	//fmt.Printf("%x\n", hash)
	block := network.GetLatestBlock("https://programmeren9.cmgt.hr.nl:8000/api/blockchain/next")
	//fmt.Println("Latest block")
	//fmt.Println("----------------")
	//fmt.Printf("Hash: %s", block.Blockchain.Hash)
	//fmt.Println()
	//fmt.Printf("Nonce: %s", block.Blockchain.Nonce)
	//fmt.Println()
	//fmt.Printf("Timestamp: %d", block.Blockchain.Timestamp)
	//fmt.Println()
	//fmt.Printf("Data: %v", block.Blockchain.Data)
	//fmt.Println()
	//fmt.Println("----------------")

	fmt.Println("Blockchain hash")
	fmt.Println(block.Blockchain.Hash)
	fmt.Println()
	LastBlockString := block.Blockchain.Hash + network.GetDataListFromBlock(block) + strconv.FormatInt(block.Blockchain.Timestamp, 10) + block.Blockchain.Nonce
	fmt.Println("LastBlock String")
	fmt.Println()
	fmt.Println(LastBlockString)
	//
	//fmt.Println("payload")
	//fmt.Println(payload)
	//BBB = hash(AAAtransactiestimestampaaa)
	lastBlockHash := mod10.HashPayload(LastBlockString)
	fmt.Println(lastBlockHash)
	fmt.Println(lastBlockHash)

	transactionsList := ""
	for _, transaction := range block.Transactions {
		trans := transaction.From + transaction.To + strconv.Itoa(transaction.Amount) + strconv.FormatInt(transaction.Timestamp, 10)
		transactionsList = transactionsList + trans
	}
	fmt.Println("New transactions")
	fmt.Println(transactionsList)
	//CCC = hash(BBBtransactiestimestampbbb)
	//from := "CMGT Mining Corporation"
	//to := "Nigel 1004416"
	//amount := "1"
	//transactionTimestamp := time.Now().Unix()
	blockTimeStamp := block.Timestamp
	////newTransaction := from + to + amount + string(transactionTimestamp)

	newBlockString := lastBlockHash + network.GetTransactionListFromBlock(block) + strconv.FormatInt(blockTimeStamp, 10)
	fmt.Println("NEW BLOCK")
	fmt.Println(newBlockString)

	//
	nonce, hash := network.FindNonce(newBlockString)
	//
	fmt.Printf("nonce: %d", nonce)
	fmt.Println()
	fmt.Printf("hash: %s", hash)
	network.RequestPayout(nonce, "Nigel 104416")
}

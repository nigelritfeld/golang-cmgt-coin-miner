package main

import (
	"fmt"
	"golang-cmgt-coin-miner/CMGTC"
)

func main() {
	fmt.Println("Mining CMGT Coin..")
	//Get the latest block
	block := CMGTC.GetLatestBlock("https://programmeren9.cmgt.hr.nl:8000/api/blockchain/next")
	// Start mining process
	nonce, _ := CMGTC.MineBlock(block)
	// Request payout
	CMGTC.RequestPayout(nonce, "Nigel 1004416")
}

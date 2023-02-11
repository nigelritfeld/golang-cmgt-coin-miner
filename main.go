package main

import (
	"fmt"
	"golang-cmgt-coin-miner/mod10"
	"golang-cmgt-coin-miner/network"
)

func main() {
	fmt.Println("Started program..")
	hash := mod10.HashPayload("test9oke")
	fmt.Printf("%x\n", hash)
	network.GetLatestBlock("https://programmeren9.cmgt.hr.nl:8000/api/blockchain/next")
}

package network

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"golang-cmgt-coin-miner/mod10"
	"io/ioutil"
	"log"
)

func GetLatestBlock(url string) {
	req.DevMode()            // Treat the package name as a Client, enable development mode
	client := req.C()        // Use C() to create a client.
	resp, err := client.R(). // Use R() to create a request.
					Get(url)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		log.Fatal(err)
	}

	// snippet only
	var result BlockResponse
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	fmt.Println(result.Blockchain.Hash)
	fmt.Println(result.Blockchain.Nonce)
	fmt.Println(result.Blockchain.Timestamp)
	fmt.Println(result.Blockchain.Data[0].From)
	fmt.Println(result.Blockchain.Data[0].To)
	fmt.Println(result.Blockchain.Data[0].Amount)
	fmt.Println(result.Blockchain.Data[0].Timestamp)

	transactions, err := json.Marshal(result.Blockchain.Data)

	payload := result.Blockchain.Hash + string(transactions) + string(result.Timestamp) + result.Blockchain.Nonce
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("payload")
	fmt.Println(payload)

	BBB := mod10.HashPayload(payload)
	fmt.Println(BBB)

	//req.EnableForceHTTP1() // Force using HTTP/1.1
	//req.MustGet("https://httpbin.org/uuid")
}

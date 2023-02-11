package network

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"golang-cmgt-coin-miner/mod10"
	"io/ioutil"
	"log"
	"strconv"
)

func GetLatestBlock(url string) BlockResponse {
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
	return result
}

type Callback func(string, string) bool

func isValidHash(str string, value string) bool {
	firstFour := str[:4]
	return firstFour == value
}

func FindNonce(block string) (int, string) {
	nonce, hash := TryNonce(block, isValidHash, 0)
	return nonce, hash
}

func TryNonce(block string, condition Callback, n int) (int, string) {
	hash := mod10.HashPayload(block + strconv.Itoa(n))
	if condition(hash, "0000") {
		return n, hash
	}
	return TryNonce(block, isValidHash, n+1)
}

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
	fmt.Println(resp)
	if err != nil {
		log.Fatal(err)
	}
}

func getTransactionString(transaction struct {
	from      string
	to        string
	amount    int
	timestamp int64
}) string {
	return transaction.from + transaction.to + strconv.Itoa(transaction.amount) + strconv.FormatInt(transaction.timestamp, 10)
}

func makeTransactionFromBlockData(transaction BlockchainTransactionFields) Transaction {
	resp := Transaction{}
	resp.from = transaction.Cfrom()
	resp.to = transaction.Cto()
	resp.amount = transaction.Camount()
	resp.timestamp = transaction.Ctimestamp()

	return resp
}
func makeTransactionFromBlock(transaction BlockTransactionFields) Transaction {
	resp := Transaction{}
	resp.from = transaction.from()
	resp.to = transaction.to()
	resp.amount = transaction.amount()
	resp.timestamp = transaction.timestamp()

	return resp
}

func GetTransactionListFromBlock(block BlockResponse) string {
	transactionList := ""
	for _, data := range block.Transactions {
		transactionList = transactionList + getTransactionString(makeTransactionFromBlock(data))
	}
	fmt.Println("TRANSACTIONLIST")
	fmt.Println(transactionList)
	return transactionList
}

func GetDataListFromBlock(block BlockResponse) string {
	transactionList := ""
	for _, data := range block.Blockchain.Data {
		transactionList = transactionList + getTransactionString(makeTransactionFromBlockData(data))
	}
	return transactionList
}

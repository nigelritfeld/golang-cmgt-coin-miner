package network

import "fmt"

//type BlockResponse []struct {
//	ID        string `json:"_id"`
//	Algorithm string `json:"algorithm"`
//	Hash      string `json:"hash"`
//	Nonce     string `json:"nonce"`
//	Timestamp int64  `json:"timestamp"`
//	V         int    `json:"__v"`
//	Data      []struct {
//		ID        string `json:"_id"`
//		From      string `json:"from"`
//		To        string `json:"to"`
//		Amount    int    `json:"amount"`
//		Timestamp int64  `json:"timestamp"`
//	} `json:"data"`
//}

type BlockResponse struct {
	Blockchain struct {
		ID        string                        `json:"_id"`
		Algorithm string                        `json:"algorithm"`
		Hash      string                        `json:"hash"`
		Nonce     string                        `json:"nonce"`
		Timestamp int64                         `json:"timestamp"`
		V         int                           `json:"__v"`
		Data      []BlockchainTransactionFields `json:"data"`
	} `json:"blockchain"`
	Transactions []BlockTransactionFields `json:"transactions"`
	Timestamp    int64                    `json:"timestamp"`
	Algorithm    string                   `json:"algorithm"`
	Open         bool                     `json:"open"`
	Countdown    int                      `json:"countdown"`
}
type Transaction struct {
	from      string
	to        string
	amount    int
	timestamp int64
}
type TransactionList struct {
	Transactions []struct {
		ID        string `json:"_id"`
		From      string `json:"from"`
		To        string `json:"to"`
		Amount    int    `json:"amount"`
		Timestamp int64  `json:"timestamp"`
		V         int    `json:"__v"`
	} `json:"transactions"`
}

//type From string

type BlockTransaction interface {
	BlockTransactionFields
	BlockTransactionMethods
}
type BlockchainTransaction interface {
	BlockchainTransactionFields
	BlockchainTransactionMethods
}
type BlockchainTransactionFields struct {
	ID        string `json:"_id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Amount    int    `json:"amount"`
	Timestamp int64  `json:"timestamp"`
}

type BlockTransactionFields struct {
	ID        string `json:"_id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Amount    int    `json:"amount"`
	Timestamp int64  `json:"timestamp"`
	V         int    `json:"__v"`
}

type BlockchainTransactionMethods interface {
	BlockTransactionFields
	Cfrom() string
	Cto() string
	Camount() int
	Ctimestamp() int64
}

type BlockTransactionMethods interface {
	BlockTransactionFields
	from() string
	to() string
	amount() int
	timestamp() int64
}

func (transaction *BlockchainTransactionFields) Cfrom() string {
	return transaction.From
}

func (transaction *BlockchainTransactionFields) Cto() string {
	return transaction.To
}
func (transaction *BlockchainTransactionFields) Camount() int {
	return transaction.Amount
}
func (transaction *BlockchainTransactionFields) Ctimestamp() int64 {
	return transaction.Timestamp
}

func (transaction *BlockTransactionFields) from() string {
	fmt.Println("return s.From")
	return transaction.From
}

func (transaction *BlockTransactionFields) to() string {
	return transaction.To
}
func (transaction *BlockTransactionFields) amount() int {
	return transaction.Amount
}
func (transaction *BlockTransactionFields) timestamp() int64 {
	return transaction.Timestamp
}

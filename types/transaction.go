package types

type Block struct {
	algorithm string
	hash      string
	nonce     string
	timestamp int
	data      []Transaction
}

// Type for block transaction
type Transaction struct {
	from      string
	to        string
	amount    string
	timestamp int
}

type BlockResponse []struct {
	ID        string `json:"_id"`
	Algorithm string `json:"algorithm"`
	Hash      string `json:"hash"`
	Nonce     string `json:"nonce"`
	Timestamp int64  `json:"timestamp"`
	V         int    `json:"__v"`
	Data      []struct {
		ID        string `json:"_id"`
		From      string `json:"from"`
		To        string `json:"to"`
		Amount    int    `json:"amount"`
		Timestamp int64  `json:"timestamp"`
	} `json:"data"`
}

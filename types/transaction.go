package main

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

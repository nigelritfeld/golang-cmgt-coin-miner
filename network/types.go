package network

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
	} `json:"blockchain"`
	Transactions []struct {
		ID        string `json:"_id"`
		From      string `json:"from"`
		To        string `json:"to"`
		Amount    int    `json:"amount"`
		Timestamp int64  `json:"timestamp"`
		V         int    `json:"__v"`
	} `json:"transactions"`
	Timestamp int64  `json:"timestamp"`
	Algorithm string `json:"algorithm"`
	Open      bool   `json:"open"`
	Countdown int    `json:"countdown"`
}

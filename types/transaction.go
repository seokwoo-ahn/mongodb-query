package types

type Tx struct {
	TxHash      string
	BlockNumber int
	Nonce       int
	GasUsed     int
	TotalIndex  int
	TxIndex     int
	TxSize      int
	From        string
	To          string
	GasPrice    int
	Amount      string
	Status      int
}

package types

type Block struct {
	BlockHash   string
	ParentHash  string
	BlockNumber int
	Time        float64
	BlockSize   int
	GasUsed     int
	TotalTxs    int
}

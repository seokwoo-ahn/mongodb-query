package types

type Event struct {
	Contract        string
	ContractAddress string
	TxHash          string
	TxIndex         int
	LogIndex        int
	Event           string
	EventFunc       string
	Data            map[string]interface{}
	BlockNumber     int
	Time            float64
	Removed         bool
}

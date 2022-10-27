package prompts

import (
	prompt "github.com/c-bata/go-prompt"
)

func SelectCollection(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "Txs", Description: "Transactions"},
		{Text: "Blocks", Description: "Blocks"},
		{Text: "Events", Description: "Events"},
		{Text: "Exit", Description: "terminate scanner"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func SelectTxQuery(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "ByHash", Description: "find transaction by transaction hash"},
		{Text: "ByBNGT", Description: "get transactions which block numbers are greater than input"},
		{Text: "Index", Description: "get Index"},
		{Text: "Exit", Description: "terminate scanner"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func SelectBlockQuery(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "ByHash", Description: "find block by block hash"},
		{Text: "ByBNGT", Description: "get blocks which block numbers are greater than input"},
		{Text: "Index", Description: "get Index"},
		{Text: "Exit", Description: "terminate scanner"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func SelectEventQuery(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "ByEventName", Description: "find events by event name"},
		{Text: "ByContractName", Description: "find events which emitted by contract name"},
		{Text: "ByContractAddress", Description: "find events which emitted by contract address"},
		{Text: "ByBlockNum", Description: "find events by block number"},
		{Text: "ByTxHash", Description: "find events by tx hash"},
		{Text: "Index", Description: "get Index"},
		{Text: "Exit", Description: "terminate scanner"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func ReceiveData(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "Exit", Description: "terminate scanner"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

package prompts

import (
	"mongodb_query/common/types"

	prompt "github.com/c-bata/go-prompt"
)

func promptExecutor(in string) {}

func NewPrompt() *types.Prompt {
	prompt := &types.Prompt{
		CollectionPrompt: NewCollectionPrompt(),
		QueryPrompt:      NewQueryPrompt(),
		TxHashPrompt:     NewTxHashPrompt(),
		BlockNumPrompt:   NewBlockNumPrompt(),
	}
	return prompt
}

func NewCollectionPrompt() *prompt.Prompt {
	return prompt.New(promptExecutor, SelectCollection, prompt.OptionPrefix("서치할 DB collection을 입력하세요 >>> "))
}

func NewQueryPrompt() *prompt.Prompt {
	return prompt.New(promptExecutor, SelectTxQuery, prompt.OptionPrefix("Tx 서치 타입을 입력하세요 >>> "))
}

func NewTxHashPrompt() *prompt.Prompt {
	return prompt.New(promptExecutor, ReceiveTxHash, prompt.OptionPrefix("Tx 해시 값을 입력하세요 >>> "))
}

func NewBlockNumPrompt() *prompt.Prompt {
	return prompt.New(promptExecutor, ReceiveBlockNum, prompt.OptionPrefix("Block number를 입력하세요 >>> "))
}

package prompts

import (
	"mongodb_query/common/types"

	prompt "github.com/c-bata/go-prompt"
)

func promptExecutor(in string) {}

func NewPrompt() *types.Prompt {
	prompt := &types.Prompt{
		CollectionPrompt: NewCollectionPrompt(),
		TxQueryPrompt:    NewTxQueryPrompt(),
		BlockQueryPrompt: NewBlockQueryPrompt(),
		EventQueryPrompt: NewEventQueryPrompt(),
		HashPrompt:       NewHashPrompt(),
		BlockNumPrompt:   NewBlockNumPrompt(),
		NamePrompt:       NewNamePrompt(),
		AddressPrompt:    NewAddressPrompt(),
	}
	return prompt
}

func NewCollectionPrompt() *prompt.Prompt {
	return prompt.New(promptExecutor, SelectCollection, prompt.OptionPrefix("서치할 DB collection을 입력하세요 >>> "))
}

func NewTxQueryPrompt() *prompt.Prompt {
	return prompt.New(promptExecutor, SelectTxQuery, prompt.OptionPrefix("Tx 서치 타입을 입력하세요 >>> "))
}

func NewBlockQueryPrompt() *prompt.Prompt {
	return prompt.New(promptExecutor, SelectBlockQuery, prompt.OptionPrefix("Block 서치 타입을 입력하세요 >>> "))
}

func NewEventQueryPrompt() *prompt.Prompt {
	return prompt.New(promptExecutor, SelectEventQuery, prompt.OptionPrefix("Event 서치 타입을 입력하세요 >>> "))
}

func NewHashPrompt() *prompt.Prompt {
	return prompt.New(promptExecutor, ReceiveData, prompt.OptionPrefix("해시 값을 입력하세요 >>> "))
}

func NewBlockNumPrompt() *prompt.Prompt {
	return prompt.New(promptExecutor, ReceiveData, prompt.OptionPrefix("Block number를 입력하세요 >>> "))
}

func NewNamePrompt() *prompt.Prompt {
	return prompt.New(promptExecutor, ReceiveData, prompt.OptionPrefix("이름을 입력하세요 >>> "))
}

func NewAddressPrompt() *prompt.Prompt {
	return prompt.New(promptExecutor, ReceiveData, prompt.OptionPrefix("주소를 입력하세요 >>> "))
}

package prompts

import prompt "github.com/c-bata/go-prompt"

func promptExecutor(in string) {}

func NewCollectionPrompt() *prompt.Prompt {
	return prompt.New(promptExecutor, SelectCollection, prompt.OptionPrefix("서치할 DB collection을 입력하세요 >>> "))
}

func NewQueryPrompt() *prompt.Prompt {
	return prompt.New(promptExecutor, SelectTxQuery, prompt.OptionPrefix("Tx 서치 타입을 입력하세요 >>> "))
}

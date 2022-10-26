package types

import "github.com/c-bata/go-prompt"

type Prompt struct {
	CollectionPrompt *prompt.Prompt
	QueryPrompt      *prompt.Prompt
	TxHashPrompt     *prompt.Prompt
	BlockNumPrompt   *prompt.Prompt
}

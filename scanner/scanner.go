package scanner

import (
	"fmt"
	"mongodb_query/db"
	"mongodb_query/libs/prompts"
	"mongodb_query/query"

	prompt "github.com/c-bata/go-prompt"
)

type Scanner struct {
	DB               *db.Database
	CollectionPrompt *prompt.Prompt
	QueryPrompt      *prompt.Prompt
	Stop             chan interface{}
}

func promptExecutor(in string) {}

func NewScanner(db *db.Database) (*Scanner, error) {
	scanner := &Scanner{
		DB:   db,
		Stop: make(chan interface{}),
	}
	scanner.CollectionPrompt = prompt.New(promptExecutor, prompts.SelectCollection)
	scanner.QueryPrompt = prompt.New(promptExecutor, prompts.SelectTxQuery)

	go scanner.ScanLoop()
	return scanner, nil
}

func (s *Scanner) ScanPrompt() {
	fmt.Println("종료하려면 Exit을 입력하세요")
	collectionType := s.CollectionPrompt.Input()
	switch collectionType {
	case "Txs":
		txCollection := s.DB.TxCollection
		txQuery := s.QueryPrompt.Input()
		switch txQuery {
		case "ByHash":
			tx, _ := query.FindTxByHash(txCollection, "0x075164408b59135a8efd2dc840147d397007552b92e14a2ca79e60d8b0d17f98")
			fmt.Println(tx)
		}
	case "Exit":
		defer close(s.Stop)
	}
}

func (s *Scanner) ScanLoop() {
	for {
		select {
		case <-s.Stop:
			return
		default:
			fmt.Println("DB 스캔을 시작합니다")
			s.ScanPrompt()
		}
	}
}

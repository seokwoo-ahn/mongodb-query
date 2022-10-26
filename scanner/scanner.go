package scanner

import (
	"fmt"
	"mongodb_query/common/prompts"
	"mongodb_query/common/types"
	"mongodb_query/db"
	"mongodb_query/db/query"
)

type Scanner struct {
	DB     *db.Database
	Prompt *types.Prompt
	Stop   chan interface{}
}

func NewScanner(db *db.Database) (*Scanner, error) {
	scanner := &Scanner{
		DB:     db,
		Stop:   make(chan interface{}),
		Prompt: prompts.NewPrompt(),
	}

	go scanner.ScanLoop()
	return scanner, nil
}

func (s *Scanner) ScanPrompt() {
	fmt.Println("종료하려면 Exit을 입력하세요")
	collectionType := s.Prompt.CollectionPrompt.Input()
	switch collectionType {
	case "Txs":
		txCollection := s.DB.TxCollection
		txQuery := s.Prompt.QueryPrompt.Input()
		switch txQuery {
		case "ByHash":
			tx, _ := query.FindTxByHash(txCollection, "0x075164408b59135a8efd2dc840147d397007552b92e14a2ca79e60d8b0d17f98")
			fmt.Println(tx)
			fmt.Println()
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

package scanner

import (
	"fmt"
	"mongodb_query/common/prompts"
	"mongodb_query/common/types"
	"mongodb_query/db"
	"mongodb_query/db/queries"
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

func (s *Scanner) ScanTx() {
	collection := s.DB.TxCollection
	query := s.Prompt.QueryPrompt.Input()
	switch query {
	case "ByHash":
		input := s.Prompt.TxHashPrompt.Input()
		if tx, err := queries.FindTxByHash(collection, input); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(tx)
		}
	case "Exit":
		defer close(s.Stop)
	}
}

func (s *Scanner) Scan() {
	fmt.Println("종료하려면 Exit을 입력하세요")
	collectionType := s.Prompt.CollectionPrompt.Input()
	switch collectionType {
	case "Txs":
		s.ScanTx()
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
			s.Scan()
		}
	}
}

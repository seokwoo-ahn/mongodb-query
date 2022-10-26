package scanner

import (
	"fmt"
	"mongodb_query/common/prompts"
	"mongodb_query/common/types"
	"mongodb_query/db"
	"mongodb_query/db/queries"
	"strconv"
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
	query := s.Prompt.TxQueryPrompt.Input()
	switch query {
	case "Index":
		if indexes, err := queries.GetIndexes(collection); err != nil {
			fmt.Println(err)
		} else {
			for _, v := range indexes {
				fmt.Println(v)
			}
		}
	case "ByHash":
		input := s.Prompt.HashPrompt.Input()
		if tx, err := queries.FindTxByHash(collection, input); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(tx)
		}
	case "ByBNGT":
		input := s.Prompt.BlockNumPrompt.Input()
		blocknum, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		if txs, err := queries.GetTxsByBlockNumberGT(collection, blocknum); err != nil {
			fmt.Println(err)
		} else {
			for _, v := range txs {
				fmt.Println(v)
			}
		}
	case "Exit":
		defer close(s.Stop)
	}
}

func (s *Scanner) ScanBlock() {
	collection := s.DB.BlockCollection
	query := s.Prompt.BlockQueryPrompt.Input()
	switch query {
	case "Index":
		if indexes, err := queries.GetIndexes(collection); err != nil {
			fmt.Println(err)
		} else {
			for _, v := range indexes {
				fmt.Println(v)
			}
		}
	case "ByHash":
		input := s.Prompt.HashPrompt.Input()
		if tx, err := queries.FindBlockByHash(collection, input); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(tx)
		}
	case "ByBNGT":
		input := s.Prompt.BlockNumPrompt.Input()
		blocknum, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		if txs, err := queries.GetBlocksByBlockNumberGT(collection, blocknum); err != nil {
			fmt.Println(err)
		} else {
			for _, v := range txs {
				fmt.Println(v)
			}
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
	case "Blocks":
		s.ScanBlock()
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

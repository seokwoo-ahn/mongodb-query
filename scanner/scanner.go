package scanner

import (
	"mongodb_query/db"
)

type Scanner struct {
	DB   *db.Database
	Quit chan struct{}
}

func NewScanner(db *db.Database) (*Scanner, error) {
	scanner := &Scanner{
		DB:   db,
		Quit: make(chan struct{}),
	}
	go scanner.loop()
	return scanner, nil
}

func (s *Scanner) loop() {
	for {
		// fmt.Println("시작합니다")
		// select {
		// case <-s.Quit:
		// 	return
		// default:
		// 	collectionType := prompt.Input(">> ", prompts.SelectCollection)
		// 	switch collectionType {
		// 	case "Txs":
		// 		txCollection := s.DB.TxCollection
		// 		txQuery := prompt.Input(">> ", prompts.SelectTxQuery)
		// 		switch txQuery {
		// 		case "ByHash":
		// 			tx, _ := query.FindTxByHash(txCollection, "0x075164408b59135a8efd2dc840147d397007552b92e14a2ca79e60d8b0d17f98")
		// 			fmt.Println(tx)
		// 		}
		// 	case "Exit":
		// 		return
		// 	}
		// }
	}
}

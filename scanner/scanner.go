package scanner

import (
	"bufio"
	"fmt"
	"mongodb_query/db"
	"os"
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
	consoleScanner := bufio.NewScanner(os.Stdin)
	cnt := 0
	for {
		fmt.Println("시작합니다")
		if !consoleScanner.Scan() {
			fmt.Println("에러발생")
			break
		}
		cnt++
		fmt.Println("hello")
		if cnt == 2 {
			break
		}
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
		// 		fmt.Println("hello world")
		// 		break test
		// 	}
		// }
	}
	// prompt.Input("아무키나 입력하시져", prompts.SelectCollection)
}

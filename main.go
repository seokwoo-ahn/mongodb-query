package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"mongodb_query/config"
	"mongodb_query/db"
	"mongodb_query/query"
)

var configFlag = flag.String("config", "./config.toml", "configuration toml file path")

func main() {
	config := config.NewConfig(*configFlag)
	database, err := db.NewDatabase(config)
	if err != nil {
		log.Fatal(err)
	}

	if tx, err := query.FindTxByHash(database.TxCollection, "0x075164408b59135a8efd2dc840147d397007552b92e14a2ca79e60d8b0d17f98"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tx)
	}

	if txs, err := query.GetTxsByBlockNumberGT(database.TxCollection, 7923820); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(txs)
	}

	err = database.Client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("몽고DB 연결을 종료했습니다!")
}

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"mongodb_query/config"
	"mongodb_query/query"
	"mongodb_query/util"
)

var configFlag = flag.String("config", "./config.toml", "configuration toml file path")

func main() {
	config := config.NewConfig(*configFlag)
	client, err := util.ConnectMongoDB(config.DataSource, config.UserName, config.PassWord)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("몽고 DB에 연결했습니다!")

	usersCollection := client.Database(config.DB).Collection(config.Collection)

	if tx, err := query.FindTxByHash(usersCollection, "0x075164408b59135a8efd2dc840147d397007552b92e14a2ca79e60d8b0d17f98"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tx)
	}

	if txs, err := query.GetTxsByBlockNumberGT(usersCollection, 7923820); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(txs)
	}

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("몽고DB 연결을 종료했습니다!")
}

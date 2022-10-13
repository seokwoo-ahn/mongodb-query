package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"mongodb_query/config"
	"mongodb_query/util"

	"go.mongodb.org/mongo-driver/bson"
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

	cursor, err := usersCollection.Find(context.TODO(), bson.D{{Key: "blocksize", Value: bson.D{{Key: "$gt", Value: 10000}}}})
	if err != nil {
		fmt.Println("error")
	}

	// var results []bson.D
	// if err = cursor.All(context.TODO(), &results); err != nil {
	// 	panic(err)
	// }
	// for _, result := range results {
	// 	fmt.Println(result)
	// }
	// fmt.Println("check")

	for cursor.Next(context.TODO()) {
		var elem bson.M
		if err := cursor.Decode(&elem); err != nil {
			fmt.Println(err)
		}
		fmt.Println(elem)
	}

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("몽고DB 연결을 종료했습니다!")
}

package db

import (
	"fmt"
	"mongodb_query/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	Config          *config.Config
	Client          *mongo.Client
	TxCollection    *mongo.Collection
	BlockCollection *mongo.Collection
	EventCollection *mongo.Collection
}

func NewDatabase(config *config.Config) (*Database, error) {
	db := &Database{
		Config: config,
	}

	client, err := ConnectMongoDB(config.DataSource, config.UserName, config.PassWord)
	if err != nil {
		return nil, err
	}
	db.Client = client
	fmt.Println("몽고 DB에 연결했습니다!")

	database := client.Database(config.DB)
	db.TxCollection = database.Collection(config.TxCollection)
	fmt.Println("트랜잭션 컬렉션에 연결했습니다!")
	db.BlockCollection = database.Collection(config.BlockCollection)
	fmt.Println("블록 컬렉션에 연결했습니다!")
	db.EventCollection = database.Collection(config.BlockCollection)
	fmt.Println("이벤트 컬렉션에 연결했습니다!")

	return db, nil
}

package db

import (
	"context"
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
	d := &Database{
		Config: config,
	}

	client, err := ConnectMongoDB(config.DataSource, config.UserName, config.PassWord)
	if err != nil {
		return nil, err
	}
	d.Client = client
	fmt.Println("몽고 DB에 연결했습니다!")

	database := client.Database(config.DB)
	d.TxCollection = database.Collection(config.TxCollection)
	fmt.Println("트랜잭션 컬렉션에 연결했습니다!")
	d.BlockCollection = database.Collection(config.BlockCollection)
	fmt.Println("블록 컬렉션에 연결했습니다!")
	d.EventCollection = database.Collection(config.BlockCollection)
	fmt.Println("이벤트 컬렉션에 연결했습니다!")

	return d, nil
}

func (d *Database) Disconnect() error {
	if d.Client != nil {
		return d.Client.Disconnect(context.TODO())
	}
	return nil
}

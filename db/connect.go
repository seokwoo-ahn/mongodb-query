package db

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/console/prompt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB(datasource, username, pw string) (*mongo.Client, error) {

	ping := func(client *mongo.Client) error {
		return client.Ping(context.TODO(), nil)
	}

	connect := func(dataSource string) (*mongo.Client, error) {
		if client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dataSource)); err != nil {
			fmt.Println("Connect")
			return nil, err
		} else if err = ping(client); err != nil {
			return nil, fmt.Errorf("ping : %v", err)
		} else {
			return client, nil
		}
	}

	makeDataSource := func(datasource, user, pw string) string {
		//mongodb:// 다음에 username과 password를 삽입.
		return string(append([]byte(datasource)[:len("mongodb://")], append([]byte(fmt.Sprintf("%s:%s@", user, pw)), []byte(datasource)[len("mongodb://"):]...)...))
	}

	if username == "" {
		return connect(datasource)
	} else if pw != "" {
		return connect(makeDataSource(datasource, username, pw))
	} else {
		var (
			err    error
			client *mongo.Client
		)

		for try := 0; try < 3; try++ {
			pw, _ = prompt.Stdin.PasswordPrompt(fmt.Sprintf("Enter the password to access the contract db (try: %d/3): ", try+1))
			if client, err = connect(makeDataSource(datasource, username, pw)); err == nil {
				return client, nil
			}
		}
		return nil, err
	}
}

func Disconnect(db *Database) error {
	if db.Client != nil {
		return db.Client.Disconnect(context.TODO())
	}
	return nil
}

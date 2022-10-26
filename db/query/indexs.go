package query

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetIndexes(collection *mongo.Collection, blockNumber int) error {
	cursor := collection.Indexes()
	fmt.Println(cursor)
	return nil
}

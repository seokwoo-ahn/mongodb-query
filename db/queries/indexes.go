package queries

import (
	"context"
	"mongodb_query/db/types"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetIndexes(collection *mongo.Collection) ([]types.Index, error) {
	cursor, _ := collection.Indexes().List(context.TODO())
	var results []types.Index
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

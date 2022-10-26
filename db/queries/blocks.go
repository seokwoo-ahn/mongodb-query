package queries

import (
	"context"
	"mongodb_query/db/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindBlockByHash(collection *mongo.Collection, blockHash string) (types.Block, error) {
	var block types.Block

	if err := collection.FindOne(context.TODO(), bson.D{{Key: "blockhash", Value: blockHash}}, nil).Decode(&block); err != nil {
		return block, err
	}
	return block, nil
}

func GetBlocksByBlockNumberGT(collection *mongo.Collection, blockNumber int) ([]types.Block, error) {
	var blocks []types.Block

	//descending
	opts := options.Find().SetSort(bson.M{"blokcnumber": 1})

	cursor, err := collection.Find(context.TODO(), bson.D{{Key: "blocknumber", Value: bson.D{{Key: "$gt", Value: blockNumber}}}}, opts)
	if err != nil {
		return nil, err
	} else if err := cursor.All(context.TODO(), &blocks); err != nil {
		return nil, err
	}
	return blocks, nil
}

package query

import (
	"context"
	"mongodb_query/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindTxByHash(collection *mongo.Collection, txHash string) (types.Tx, error) {
	var tx types.Tx

	// projection
	// opts := []*options.FindOneOptions{
	// 	options.FindOne().SetProjection(bson.M{"blocknumber": 1}),
	// }

	if err := collection.FindOne(context.TODO(), bson.D{{Key: "txhash", Value: txHash}}, nil).Decode(&tx); err != nil {
		return tx, err
	}
	return tx, nil
}

func GetTxsByBlockNumberGT(collection *mongo.Collection, blockNumber int) ([]types.Tx, error) {
	var txs []types.Tx

	//descending
	opts := options.Find().SetSort(bson.M{"blokcnumber": 1})

	cursor, err := collection.Find(context.TODO(), bson.D{{Key: "blocknumber", Value: bson.D{{Key: "$gt", Value: blockNumber}}}}, opts)
	if err != nil {
		return nil, err
	} else if err := cursor.All(context.TODO(), &txs); err != nil {
		return nil, err
	}
	return txs, nil
}

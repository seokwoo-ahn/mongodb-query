package query

import (
	"context"
	"mongodb_query/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindTxByHash(collection *mongo.Collection, txHash string) (types.Tx, error) {
	var tx types.Tx
	cursor, err := collection.Find(context.TODO(), bson.D{{Key: "txhash", Value: txHash}})
	if err != nil {
		return tx, err
	}

	for cursor.Next(context.TODO()) {
		var elem bson.M
		if err := cursor.Decode(&elem); err != nil {
			return tx, err
		}
		tx.TxHash = elem["txhash"].(string)
		tx.BlockNumber = int(elem["blocknumber"].(float64))
		tx.Nonce = int(elem["nonce"].(float64))
		tx.GasUsed = int(elem["gasused"].(float64))
		tx.TotalIndex = int(elem["totalindex"].(float64))
		tx.TxSize = int(elem["txsize"].(float64))
		tx.From = elem["from"].(string)
		tx.To = elem["to"].(string)
		tx.GasPrice = int(elem["gasprice"].(float64))
		tx.Amount = elem["amount"].(string)
		tx.Status = int(elem["status"].(float64))
	}
	return tx, nil
}

func GetTxsByBlockNumber(collection *mongo.Collection, txHash string) (types.Tx, error) {
	var tx types.Tx
	cursor, err := collection.Find(context.TODO(), bson.D{{Key: "blocksize", Value: bson.D{{Key: "$gt", Value: 10000}}}})
	if err != nil {
		return tx, err
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
			return tx, err
		}
	}
	return tx, nil
}

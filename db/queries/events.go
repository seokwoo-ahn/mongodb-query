package queries

import (
	"context"
	"mongodb_query/db/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindEventsByBlockNum(collection *mongo.Collection, blockNum int) ([]types.Event, error) {
	var events []types.Event
	opts := options.Find().SetSort(bson.M{"logindex": -1})

	cursor, err := collection.Find(context.TODO(), bson.D{{Key: "blocknumber", Value: blockNum}}, opts)
	if err != nil {
		return nil, err
	} else if err := cursor.All(context.TODO(), &events); err != nil {
		return nil, err
	}
	return events, nil
}

func FindEventsByTxHash(collection *mongo.Collection, txHash string) ([]types.Event, error) {
	var events []types.Event
	opts := options.Find().SetSort(bson.M{"logindex": -1})

	cursor, err := collection.Find(context.TODO(), bson.D{{Key: "txhash", Value: txHash}}, opts)
	if err != nil {
		return nil, err
	} else if err := cursor.All(context.TODO(), &events); err != nil {
		return nil, err
	}
	return events, nil
}

func FindEventsByEventName(collection *mongo.Collection, eventName string) ([]types.Event, error) {
	var events []types.Event
	opts := options.Find().SetSort(bson.M{"blokcnumber": 1})

	cursor, err := collection.Find(context.TODO(), bson.D{{Key: "event", Value: eventName}}, opts)
	if err != nil {
		return nil, err
	} else if err := cursor.All(context.TODO(), &events); err != nil {
		return nil, err
	}
	return events, nil
}

func FindEventsByContractName(collection *mongo.Collection, contractName string) ([]types.Event, error) {
	var events []types.Event
	opts := options.Find().SetSort(bson.M{"blokcnumber": 1})

	cursor, err := collection.Find(context.TODO(), bson.D{{Key: "contract", Value: contractName}}, opts)
	if err != nil {
		return nil, err
	} else if err := cursor.All(context.TODO(), &events); err != nil {
		return nil, err
	}
	return events, nil
}

func FindEventsByContractAddress(collection *mongo.Collection, contractAddress string) ([]types.Event, error) {
	var events []types.Event
	opts := options.Find().SetSort(bson.M{"blokcnumber": 1})

	cursor, err := collection.Find(context.TODO(), bson.D{{Key: "contractaddress", Value: contractAddress}}, opts)
	if err != nil {
		return nil, err
	} else if err := cursor.All(context.TODO(), &events); err != nil {
		return nil, err
	}
	return events, nil
}

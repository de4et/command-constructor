package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

const DBNAME = "command-constructor"

type Dropper interface {
	Drop(context.Context) error
}

type Store struct {
	User    UserStore
	Command CommandStore
}

func ConvertStructToBson(v any) (bson.M, error) {
	bsonData, err := bson.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	// Преобразование BSON в bson.M
	var bsonMap bson.M
	err = bson.Unmarshal(bsonData, &bsonMap)
	if err != nil {
		return nil, err
	}

	return bsonMap, nil
}

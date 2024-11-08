package db

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func ConvertStructToBson(v any) (bson.M, error) {
	bsonData, err := bson.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	var bsonMap bson.M
	err = bson.Unmarshal(bsonData, &bsonMap)
	if err != nil {
		return nil, err
	}

	return bsonMap, nil
}

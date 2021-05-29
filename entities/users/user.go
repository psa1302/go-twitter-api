package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db mongoDB) getAllUsers() (interface{}, error) {
	var results []UserModel
	var err error

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := db.recipes.Find(ctx, bson.D{}, options.Find())

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var elem UserModel
		err = cur.Decode(&elem)

		if err != nil {
			return nil, err
		}

		results.append(results, elem)
	}

	if err = cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(ctx)
	return results, nil
}

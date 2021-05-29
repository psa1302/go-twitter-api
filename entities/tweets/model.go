package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TweetModel struct {
	ID        primitive.ObjectID
	tweet     string
	createdAt time.Time
	updated   time.Time
}

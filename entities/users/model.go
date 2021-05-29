package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID          primitive.ObjectID
	firstName   string
	lastName    string
	email       string
	password    string
	acceptTerms bool
	isActive    bool
	createdAt   time.Time
	updatedAt   time.Time
}

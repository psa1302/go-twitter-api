package main

import (
	"github.com/graphql-go/graphql"
)

var TweetType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Tweet",
	Fields: graphql.Fields{
		"_id":       &graphql.Field{Type: ID},
		"tweet":     &graphql.Field{Type: graphql.String},
		"createdAt": &graphql.Field{Type: graphql.String},
		"updatedAt": &graphql.Field{Type: graphql.String},
	},
})

package main

import (
	"log"

	"github.com/graph-gophers/graphql-go"
	"github.com/graphql-go/graphql"
)

func initSchema() graphql.Schema {
	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"getAllTweets": &graphql.Field{
					Type:    graphql.NewList(TweetType),
					Args:    graphql.FieldConfigArgument{},
					Resolve: getAllTweets,
				},
				"getAllUsers": &graphql.Field{
					Type:    graphql.NewList(UserType),
					Args:    graphql.FieldArgument{},
					Resolve: getAllUsers,
				},
			},
		}),
	})
	if err != nil {
		return log.Fatal(err)
	}

	return graphqlSchema
}

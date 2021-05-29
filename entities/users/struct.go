package main

import "github.com/graphql-go/graphql"

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"_id":         &graphql.Field{Type: ID},
		"firstName":   &graphql.Field{Type: graphql.String},
		"lastName":    &graphql.Field{Type: graphql.String},
		"email":       &graphql.Field{Type: graphql.String},
		"password":    &graphql.Field{Type: graphql.String},
		"acceptTerms": &graphql.Field{Type: graphql.Boolean},
		"isActive":    &graphql.Field{Type: graphql.Boolean},
		"createdAt":   &graphql.Field{Type: graphql.String},
		"updatedAt":   &graphql.Field{Type: graphql.String},
	},
})

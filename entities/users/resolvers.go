package main

import "github.com/graphql-go/graphql"

func getAllUsers(_ graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	results, err = db.getAllUsers()
	if err != nil {
		return nil, err
	}

	return results, nil
}

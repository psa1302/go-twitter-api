package main

import "github.com/graphql-go/graphql"

func getAllTweets(_ graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	results, err = db.getAllTweets()
	if err != nil {
		return nil, err
	}

	return results, nil
}

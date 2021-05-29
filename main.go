package main

import (
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go/relay"
)

var db mongoDB

func main() {
	db = connectDB()
	defer db.close()

	schema := initSchema()
	h := &relay.Handler{Schema: &schema}

	http.Handle("/graphql", disableCors(headerAuthorization(h)))
	err := http.ListenAndServe(config.serveUri, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

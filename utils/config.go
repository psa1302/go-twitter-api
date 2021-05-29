package main

type configModel struct {
	mongoUri    string
	mongoDb     string
	tokenSecret string
	tokenExp    string
	serveUri    string
}

var config = configModel{
	mongoUri:    "mongodb://localhost:27017/twitter",
	mongoDb:     "react-recipes",
	tokenSecret: "secret",
	tokenExp:    "1h",
	serveUri:    ":4444",
}

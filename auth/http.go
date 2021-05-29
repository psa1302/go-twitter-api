package main

import (
	"context"
	"net/http"
)

func headerAuthorization(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context
		ctx = context.WithValue(r.Context(), "email", "")
		auth := r.Header.Get("authorization")

		if auth != "" {
			token, _ := verifyToken(auth)
			if token != nil {
				ctx = context.WithValue(r.Context(), "email", token["email"])
			}
		}
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func disableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Method", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}

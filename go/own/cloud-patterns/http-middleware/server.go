/**
 * curl -H "Token:secret" localhost:8080/restricted
 */
package main

import (
	"context"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
)

type ctxKey string

const keyUserName ctxKey = "user_id"

func main() {
	mux := http.NewServeMux()
	mux.Handle("/restricted", authMiddleware(handleRestricted()))
	mux.Handle("/anonymous", handleAnonymous())
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func handleAnonymous() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Token")
		if token != "secret" {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, "invalid token")
			return
		}

		rctx := context.WithValue(r.Context(), keyUserName, "Dj. Groove")
		next.ServeHTTP(w, r.WithContext(rctx))
	})
}

func handleRestricted() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		userName, ok := r.Context().Value(keyUserName).(string)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "internal error")
			return
		}
		fmt.Fprintf(
			w,
			"Access granted for user %q to resource %q",
			userName,
			html.EscapeString(r.URL.Path))
	})
}

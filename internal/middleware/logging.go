package middleware

import (
	"fmt"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Method %s | Path: %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

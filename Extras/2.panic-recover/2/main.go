package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

// Middleware
func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Into Middleware...")
		defer func() {
			if r := recover(); r != nil {
				log.Printf("recovered panic: %v\n", r)
				debug.PrintStack()
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("panic")
	})

	log.Println("Listening on server :8080")
	if err := http.ListenAndServe(":8080", recoverMiddleware(mux)); err != nil {
		log.Fatalf("Error while listening on server: %v", err)
	}
}

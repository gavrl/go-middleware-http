package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", myMiddleware(myHandler))
	_ = http.ListenAndServe(":8080", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello from Handler")
	w.WriteHeader(http.StatusOK)
}

func myMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello from middleware")
		next(w, r)
	}
}

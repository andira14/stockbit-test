package main

import (
	"net/http"
	"questions/question2/src/httpHandler"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env")
	}

	server := http.NewServeMux()
	httpHandler.Request("/movies", "GET", server)

	http.ListenAndServe(":8080", server)
}

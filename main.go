package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	log.Println("Server listening on: :8800")
	server := http.ListenAndServe(":8800", router)
	log.Fatal(server)
}

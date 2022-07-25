package main

import (
	"log"
	"main/config"
	"net/http"
)

func main() {
	router := NewRouter()
	config := config.ReadConfig()
	log.Println("Server listening on: " + config.Port)
	server := http.ListenAndServe(config.Port, router)
	log.Fatal(server)
}

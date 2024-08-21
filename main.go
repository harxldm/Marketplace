package main

import (
	handler "backend_en_go/Handler"
	"backend_en_go/Storage"
	"log"
	"net/http"
)

func main() {
	Storage.NewPostgresDB()

	http.HandleFunc("/register", handler.Register)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

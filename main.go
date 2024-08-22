package main

import (
	handler "backend_en_go/Handler"
	"backend_en_go/Storage"
	"log"
	"net/http"
)

func main() {
	Storage.NewPostgresDB()

	// db := Storage.Pool()

	// userStorage := Storage.NewPsqlUser(db)

	// mux := http.NewServeMux()

	// handler.RouteUser(mux, userStorage)

	// log.Println("Server is running on port 8080")
	// if err := http.ListenAndServe(":8080", mux); err != nil {
	// 	log.Fatalf("Server failed to start: %v", err)
	// }

	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/Login", handler.Login)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

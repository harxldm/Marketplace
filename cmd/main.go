package main

import (
	authorization "backend_en_go/Authorization"
	handler "backend_en_go/Handler"
	"backend_en_go/Storage"
	"backend_en_go/middleware"
	"log"
	"net/http"
)

func handleCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Permite cualquier origen
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	Storage.NewPostgresDB()

	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("No se pudo cargar los certificados: %v", err)
	}
	// db := Storage.Pool()

	// userStorage := Storage.NewPsqlUser(db)

	// mux := http.NewServeMux()

	// handler.RouteUser(mux, userStorage)

	// log.Println("Server is running on port 8080")
	// if err := http.ListenAndServe(":8080", mux); err != nil {
	// 	log.Fatalf("Server failed to start: %v", err)
	// }

	mux := http.NewServeMux()
	mux.HandleFunc("/register", handler.Register)
	mux.HandleFunc("/login", handler.Login)
	mux.Handle("/createProduct", middleware.AuthMiddleware(http.HandlerFunc(handler.NewProduct)))
	mux.HandleFunc("/getAll", handler.GetAll)
	mux.HandleFunc("/getById", handler.GetProductsBySellerID)

	// mux.HandleFunc("/createProduct", handler.NewProduct)

	log.Fatal(http.ListenAndServe(":8080", handleCORS(mux)))
}

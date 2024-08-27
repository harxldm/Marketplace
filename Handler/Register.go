package handler

import (
	model "backend_en_go/Model"
	"backend_en_go/Storage"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func setupCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Register(w http.ResponseWriter, r *http.Request) {

	setupCORS(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Error al decodificar el cuerpo de la solicitud:", err)
		http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	log.Println("Datos recibidos:", user)

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Println("Error al hashear la contraseña:", err)
		http.Error(w, "Error al hashear la contraseña", http.StatusInternalServerError)
		return
	}

	user.Password = hashedPassword
	user.Created_date = time.Now()

	// Obtener la conexión a la base de datos
	db := Storage.Pool()

	// Instanciar el repositorio de usuario
	userRepo := Storage.NewPsqlUser(db)

	// Crear el usuario
	err = userRepo.CreateUser(&user)
	if err != nil {
		log.Println("Error al registrar el usuario:", err)
		http.Error(w, "Error al registrar el usuario: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Usuario registrado exitosamente",
		"user":    user.UserID,
	})
}

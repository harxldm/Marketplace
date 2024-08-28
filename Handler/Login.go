package handler

import (
	authorization "backend_en_go/Authorization"
	model "backend_en_go/Model"
	"backend_en_go/Storage"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Ejemplo en Go para evitar el cacheo

func UserExist(w http.ResponseWriter, r *http.Request) {
	// Obtener el parámetro de la consulta
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	// Obtener la conexión a la base de datos y crear el repositorio de usuarios
	db := Storage.Pool()
	userRepo := Storage.NewPsqlUser(db)

	// Verificar si el usuario existe
	_, err := userRepo.GetUserByEmail(email)
	if err != nil {
		if err.Error() == "user not found" {
			// El correo electrónico no está registrado
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]bool{"exists": false})
			return
		}
		// Error al consultar la base de datos
		http.Error(w, "Error checking email existence", http.StatusInternalServerError)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials model.Login

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return

	}
	db := Storage.Pool()
	userRepo := Storage.NewPsqlUser(db)

	user, err := userRepo.GetUserByEmail(credentials.Email)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	credentials.Rol = user.Rol

	credentials.UserID = int(user.UserID)

	token, err := authorization.GenerateToken(&credentials)
	if err != nil {
		http.Error(w, "No se pudo generar el token", http.StatusInternalServerError)
		return

	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful",
		"user":    user.UserID,
		"Token":   token,
	})

}

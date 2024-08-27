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

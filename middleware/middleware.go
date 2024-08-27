package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Token no proporcionado", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("mi_clave_secreta"), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "No se pudieron extraer las reclamaciones del token", http.StatusUnauthorized)
			return
		}

		userID, ok := claims["userID"].(float64)
		if !ok {
			http.Error(w, "No se pudo obtener el ID del usuario del token", http.StatusUnauthorized)
			return
		}

		// Coloca el userID en el contexto de la solicitud
		ctx := context.WithValue(r.Context(), "userID", int(userID))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

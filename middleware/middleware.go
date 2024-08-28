package middleware

import (
	authorization "backend_en_go/Authorization"
	"context"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Excluir la ruta /getById de la autenticación
		if r.URL.Path == "/getById" && r.Method == "GET" {
			next.ServeHTTP(w, r)
			return
		}

		if r.URL.Path == "/createProducts" && r.Method == "POST" {
			next.ServeHTTP(w, r)
			return
		}

		// Obtener el token del encabezado Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "No autorizado", http.StatusUnauthorized)
			return
		}

		// Eliminar el prefijo "Bearer " si se está utilizando
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validar y decodificar el token
		claims, err := authorization.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "No autorizado", http.StatusUnauthorized)
			return
		}

		// Colocar el userID en el contexto de la solicitud
		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

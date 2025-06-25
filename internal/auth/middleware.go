package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// ContextKey тип для ключей контекста
type ContextKey string

const (
	UserContextKey ContextKey = "user"
)

// JWTMiddleware middleware для проверки JWT токенов
func JWTMiddleware(jwtConfig *JWTConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Получаем токен из заголовка Authorization
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				sendErrorResponse(w, "Authorization header required", http.StatusUnauthorized)
				return
			}

			// Проверяем формат "Bearer TOKEN"
			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
				sendErrorResponse(w, "Invalid authorization header format", http.StatusUnauthorized)
				return
			}

			// Валидируем токен
			claims, err := jwtConfig.ValidateToken(bearerToken[1])
			if err != nil {
				sendErrorResponse(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// Создаем пользователя из claims
			user := User{
				ID:       claims.UserID,
				Username: claims.Username,
				Email:    claims.Email,
			}

			// Добавляем пользователя в контекст
			ctx := context.WithValue(r.Context(), UserContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetUserFromContext извлекает пользователя из контекста
func GetUserFromContext(ctx context.Context) (User, bool) {
	user, ok := ctx.Value(UserContextKey).(User)
	return user, ok
}

// sendErrorResponse отправляет JSON ответ с ошибкой
func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]string{
		"error": message,
	}

	json.NewEncoder(w).Encode(response)
}

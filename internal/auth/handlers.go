package auth

import (
	"encoding/json"
	"net/http"
)

// Handler структура для обработчиков авторизации
type Handler struct {
	service Service
}

// NewHandler создает новый handler
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Register godoc
// @Summary Регистрация нового пользователя
// @Description Создает нового пользователя в системе
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Данные для регистрации"
// @Success 201 {object} AuthResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/register [post]
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Простая валидация
	if req.Username == "" || req.Email == "" || req.Password == "" {
		sendJSONError(w, "Username, email and password are required", http.StatusBadRequest)
		return
	}

	if len(req.Password) < 6 {
		sendJSONError(w, "Password must be at least 6 characters", http.StatusBadRequest)
		return
	}

	response, err := h.service.Register(req)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// Login godoc
// @Summary Авторизация пользователя
// @Description Авторизует пользователя и возвращает JWT токен
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Данные для входа"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Простая валидация
	if req.Email == "" || req.Password == "" {
		sendJSONError(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	response, err := h.service.Login(req)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Profile godoc
// @Summary Получить профиль пользователя
// @Description Возвращает информацию о текущем пользователе
// @Tags auth
// @Produce json
// @Security Bearer
// @Success 200 {object} User
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /auth/profile [get]
func (h *Handler) Profile(w http.ResponseWriter, r *http.Request) {
	// Получаем пользователя из контекста
	user, ok := GetUserFromContext(r.Context())
	if !ok {
		sendJSONError(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	// Получаем полную информацию о пользователе
	profile, err := h.service.GetProfile(user.ID)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

// sendJSONError отправляет JSON ответ с ошибкой
func sendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]string{
		"error": message,
	}

	json.NewEncoder(w).Encode(response)
}

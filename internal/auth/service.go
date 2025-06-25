package auth

import (
	"errors"
)

// Service интерфейс для бизнес-логики авторизации
type Service interface {
	Register(req RegisterRequest) (*AuthResponse, error)
	Login(req LoginRequest) (*AuthResponse, error)
	GetProfile(userID int) (*User, error)
}

// service реализация Service
type service struct {
	repo      Repository
	jwtConfig *JWTConfig
}

// NewService создает новый сервис
func NewService(repo Repository, jwtConfig *JWTConfig) Service {
	return &service{
		repo:      repo,
		jwtConfig: jwtConfig,
	}
}

// Register регистрирует нового пользователя
func (s *service) Register(req RegisterRequest) (*AuthResponse, error) {
	// Создаем пользователя
	user, err := s.repo.CreateUser(req)
	if err != nil {
		return nil, err
	}

	// Создаем JWT токен
	token, err := s.jwtConfig.CreateToken(*user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

// Login авторизует пользователя
func (s *service) Login(req LoginRequest) (*AuthResponse, error) {
	// Получаем пользователя по email
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Проверяем пароль
	if !s.repo.VerifyPassword(user.Password, req.Password) {
		return nil, errors.New("invalid credentials")
	}

	// Создаем JWT токен
	token, err := s.jwtConfig.CreateToken(*user)
	if err != nil {
		return nil, err
	}

	// Убираем пароль из ответа
	user.Password = ""

	return &AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

// GetProfile получает профиль пользователя
func (s *service) GetProfile(userID int) (*User, error) {
	return s.repo.GetUserByID(userID)
}

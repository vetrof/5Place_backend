package auth

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Repository интерфейс для работы с пользователями
type Repository interface {
	CreateUser(req RegisterRequest) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	VerifyPassword(hashedPassword, password string) bool
}

// repository реализация Repository
type repository struct {
	db *sql.DB
}

// NewRepository создает новый репозиторий
func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

// CreateUser создает нового пользователя
func (r *repository) CreateUser(req RegisterRequest) (*User, error) {
	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Проверяем, не существует ли уже пользователь с таким email
	var exists bool
	checkQuery := "SELECT EXISTS(SELECT 1 FROM app_user WHERE email = $1)"
	err = r.db.QueryRow(checkQuery, req.Email).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("user with this email already exists")
	}

	// Создаем пользователя
	query := `
		INSERT INTO app_user (username, email, password) 
		VALUES ($1, $2, $3) 
		RETURNING id, username, email`

	user := &User{}
	err = r.db.QueryRow(query, req.Username, req.Email, string(hashedPassword)).
		Scan(&user.ID, &user.Username, &user.Email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByEmail получает пользователя по email
func (r *repository) GetUserByEmail(email string) (*User, error) {
	query := "SELECT id, username, email, password FROM app_user WHERE email = $1"

	user := &User{}
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

// GetUserByID получает пользователя по ID
func (r *repository) GetUserByID(id int) (*User, error) {
	query := "SELECT id, username, email FROM app_user WHERE id = $1"

	user := &User{}
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

// VerifyPassword проверяет пароль
func (r *repository) VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

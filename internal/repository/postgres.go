package repository

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// PostgresDB представляет реализацию репозитория с использованием PostgreSQL
type PostgresDB struct {
	DB *sql.DB
}

// NewPostgresDB создает и инициализирует новое подключение к PostgreSQL
func NewPostgresDB() (*PostgresDB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
		// Продолжаем работу, возможно переменные окружения установлены иначе
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	fmt.Println("Подключение к базе данных:", host)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("database is not reachable: %w", err)
	}

	// Проверка доступности PostGIS
	var version string
	err = db.QueryRow("SELECT PostGIS_Version()").Scan(&version)
	if err != nil {
		return nil, fmt.Errorf("PostGIS не установлен или не доступен: %w", err)
	}
	fmt.Println("Версия PostGIS:", version)

	// Создаем таблицы
	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("cannot create tables: %w", err)
	}

	return &PostgresDB{DB: db}, nil
}

// Close закрывает соединение с базой данных
func (p *PostgresDB) Close() error {
	return p.DB.Close()
}

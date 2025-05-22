package repository

import (
	"5Place/internal/models"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

// Вспомогательная функция для создания таблиц
func createTables(db *sql.DB) error {
	city := `
   CREATE TABLE IF NOT EXISTS city (
      id SERIAL PRIMARY KEY,
      name TEXT UNIQUE NOT NULL
   );`

	user := `
   CREATE TABLE IF NOT EXISTS "user" (
      id SERIAL PRIMARY KEY,
      external_id TEXT UNIQUE, -- для вашего текстового ID
      imei TEXT,
      telegram_id TEXT,
      email TEXT,
      phone TEXT
   );`

	userPlace := `
   CREATE TABLE IF NOT EXISTS user_place (
      id SERIAL PRIMARY KEY,
      user_id INTEGER REFERENCES "user"(id) ON UPDATE CASCADE ON DELETE CASCADE,
      info TEXT,
      geom GEOGRAPHY(POINT,4326)
   );`

	place := `
   CREATE TABLE IF NOT EXISTS place (
      id SERIAL PRIMARY KEY,
      city_id INTEGER REFERENCES city(id) ON UPDATE CASCADE ON DELETE SET NULL,
      name TEXT,
      geom GEOGRAPHY(POINT,4326),
      descr TEXT
   );`

	for i, q := range []string{city, user, userPlace, place} {
		tableName := []string{"city", "user", "user_place", "place"}[i]
		if _, err := db.Exec(q); err != nil {
			return fmt.Errorf("ошибка создания таблицы %s: %w", tableName, err)
		} else {
			fmt.Printf("Таблица %s создана или уже существует\n", tableName)
		}
	}

	return nil
}

// GetNearPlaces находит места рядом с указанными координатами
func (db *PostgresDB) GetNearPlaces(lat, long float64) ([]models.Place, error) {
	query := fmt.Sprintf(`
		SELECT id, city_name, name, ST_AsText(geom) as geom, descr, 
		ST_Distance(geom::geography, ST_SetSRID(ST_MakePoint($1, $2), 4326)::geography) AS distance
		FROM %s.place
		ORDER BY distance ASC
		LIMIT 20
	`, os.Getenv("DB_SCHEMA")) // Или передай схему явно, если удобнее

	rows, err := db.DB.Query(query, long, lat)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var places []models.Place
	for rows.Next() {
		var p models.Place
		if err := rows.Scan(&p.ID, &p.CityName, &p.Name, &p.Geom, &p.Desc, &p.Distance); err != nil {
			return nil, fmt.Errorf("row scan error: %w", err)
		}
		places = append(places, p)
	}

	return places, nil
}

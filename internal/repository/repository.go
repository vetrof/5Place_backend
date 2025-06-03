package repository

import (
	"5Place/internal/models"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

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

// GetNearPlaces находит места рядом с указанными координатами
func (db *PostgresDB) GetNearPlaces(lat, long float64) ([]models.Place, error) {
	query := fmt.Sprintf(`
		SELECT p.id, c.name AS city_name, p.name, ST_AsText(p.geom) as geom, p.descr, 
		ST_Distance(p.geom::geography, ST_SetSRID(ST_MakePoint($1, $2), 4326)::geography) AS distance
		FROM %[1]s.place p
		JOIN %[1]s.city c ON p.city_id = c.id
		ORDER BY distance ASC
		LIMIT 20
	`, os.Getenv("DB_SCHEMA"))

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

		photos, err := db.GetPhotosByPlaceID(p.ID)
		if err != nil {
			return nil, err
		}
		p.Photos = photos

		places = append(places, p)
	}

	return places, nil
}

func (db *PostgresDB) GetAllCities() ([]models.City, error) {
	query := fmt.Sprintf(`
		SELECT c.id, c.name, ST_AsText(c.geom) as geom, COUNT(p.id) as points
		FROM %s.city c
		LEFT JOIN %s.place p ON p.city_id = c.id
		GROUP BY c.id, c.name, c.geom
		ORDER BY c.id
		LIMIT 20
	`, os.Getenv("DB_SCHEMA"), os.Getenv("DB_SCHEMA"))

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var cities []models.City
	for rows.Next() {
		var c models.City
		if err := rows.Scan(&c.ID, &c.Name, &c.Geom, &c.Points); err != nil {
			return nil, fmt.Errorf("row scan error: %w", err)
		}
		cities = append(cities, c)
	}

	return cities, nil
}

func (db *PostgresDB) GetPhotosByPlaceID(placeID int) ([]string, error) {
	query := fmt.Sprintf(`
		SELECT file_link FROM %s.photo WHERE place_id = $1
	`, os.Getenv("DB_SCHEMA"))

	rows, err := db.DB.Query(query, placeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []string
	for rows.Next() {
		var filePath string
		if err := rows.Scan(&filePath); err != nil {
			return nil, err
		}
		photos = append(photos, filePath)
	}

	return photos, nil
}

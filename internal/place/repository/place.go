package repository

import (
	"5Place/internal/place/models"
	"fmt"
	"os"
)

// GetNearPlaces находит места рядом с указанными координатами
func (db *PostgresDB) GetNearPlaces(lat, long float64, limit int, radius float64) ([]models.Place, error) {
	query := fmt.Sprintf(`
		SELECT p.id, c.name AS city_name, p.name, ST_AsText(p.geom) as geom, p.descr, 
		ST_Distance(p.geom::geography, ST_SetSRID(ST_MakePoint($1, $2), 4326)::geography) AS distance
		FROM app_place p
		JOIN app_city c ON p.city_id = c.id
		ORDER BY distance ASC
		LIMIT 20
	`)

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

// PlaceDetail
func (db *PostgresDB) GetPlaceDetail(placeID int) ([]models.Place, error) {
	// Простой запрос без сортировки по расстоянию
	query := `
        SELECT p.id, c.name AS city_name, p.name, ST_AsText(p.geom) as geom, p.descr
        FROM app_place p
        JOIN app_city c ON p.city_id = c.id
        WHERE p.id = $1
        ORDER BY p.name ASC
        LIMIT 20`

	rows, err := db.DB.Query(query, placeID)
	if err != nil {
		return nil, fmt.Errorf("failed to query places: %w", err)
	}
	defer rows.Close()

	var places []models.Place
	for rows.Next() {
		var place models.Place
		var cityName string
		var geomText string

		err := rows.Scan(
			&place.ID,
			&cityName,
			&place.Name,
			&geomText,
			&place.Desc, // предполагаю что поле называется Description
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan place: %w", err)
		}

		// Если нужно сохранить название города в структуре
		place.CityName = cityName
		place.Geom = geomText // или парсить геометрию если нужно

		places = append(places, place)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating places: %w", err)
	}

	return places, nil
}

// GetAllCityPlaces выводит все места города
func (db *PostgresDB) GetAllCityPlaces(cityID int) ([]models.Place, error) {
	// Простой запрос без сортировки по расстоянию
	query := `
        SELECT p.id, c.name AS city_name, p.name, ST_AsText(p.geom) as geom, p.descr
        FROM app_place p
        JOIN app_city c ON p.city_id = c.id
        WHERE p.city_id = $1
        ORDER BY p.name ASC
        LIMIT 20`

	rows, err := db.DB.Query(query, cityID)
	if err != nil {
		return nil, fmt.Errorf("failed to query places: %w", err)
	}
	defer rows.Close()

	var places []models.Place
	for rows.Next() {
		var place models.Place
		var cityName string
		var geomText string

		err := rows.Scan(
			&place.ID,
			&cityName,
			&place.Name,
			&geomText,
			&place.Desc, // предполагаю что поле называется Description
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan place: %w", err)
		}

		// Если нужно сохранить название города в структуре
		place.CityName = cityName
		place.Geom = geomText // или парсить геометрию если нужно

		places = append(places, place)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating places: %w", err)
	}

	return places, nil
}

// функция для получения фото к месту
func (db *PostgresDB) GetPhotosByPlaceID(placeID int) ([]string, error) {
	query := fmt.Sprintf(`
		SELECT image FROM %s.app_photo WHERE place_id = $1
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

func (db *PostgresDB) GetRandomPlaces(countryId *int64, cityId *int64) ([]models.Place, error) {
	query := `
        SELECT p.id, c.name AS city_name, p.name, ST_AsText(p.geom) as geom, p.descr
        FROM app_place p
        JOIN app_city c ON p.city_id = c.id
    `
	var args []any

	if countryId != nil {
		query += "WHERE c.country_id = $1\n"
		args = append(args, *countryId)
	} else if cityId != nil {
		query += "WHERE c.id = $1\n"
		args = append(args, *cityId)
	}

	query += "ORDER BY random() LIMIT 100"

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query places: %w", err)
	}
	defer rows.Close()

	var places []models.Place
	for rows.Next() {
		var place models.Place
		var cityName string
		var geomText string

		if err := rows.Scan(
			&place.ID,
			&cityName,
			&place.Name,
			&geomText,
			&place.Desc,
		); err != nil {
			return nil, fmt.Errorf("failed to scan place: %w", err)
		}

		place.CityName = cityName
		place.Geom = geomText
		places = append(places, place)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating places: %w", err)
	}

	return places, nil
}

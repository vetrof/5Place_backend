package repository

import (
	"5Place/internal/place/models"
	"database/sql"
	"fmt"
	"os"
)

// функция для получения фото к месту
func (db *PostgresDB) GetPhotosByPlaceID(placeID int, limit int) ([]string, error) {
	query := fmt.Sprintf(`
		SELECT image FROM %s.app_photo WHERE place_id = $1 LIMIT $2
	`, os.Getenv("DB_SCHEMA"))

	rows, err := db.DB.Query(query, placeID, limit)
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

// GetNearPlaces находит места рядом с указанными координатами
func (db *PostgresDB) GetNearPlaces(lat, long float64, limit int, radius float64) ([]models.Place, error) {
	query := fmt.Sprintf(`
		SELECT p.id, c.name AS city_name, p.name, ST_AsText(p.geom) as geom, p.descr,
		ST_Distance(p.geom::geography, ST_SetSRID(ST_MakePoint($1, $2), 4326)::geography) AS distance,
		ST_Y(p.geom::geometry) AS latitude,
    	ST_X(p.geom::geometry) AS longitude
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
	var latVal, lngVal float64

	for rows.Next() {
		var p models.Place
		if err := rows.Scan(&p.ID, &p.CityName, &p.Name, &p.Geom, &p.Desc, &p.Distance, &latVal,
			&lngVal); err != nil {
			return nil, fmt.Errorf("row scan error: %w", err)
		}

		photos, err := db.GetPhotosByPlaceID(p.ID, 1) // лимит на отображение фоток - 1
		if err != nil {
			return nil, err
		}
		p.Photos = photos
		p.Coordinates = models.Coordinates{
			Lat: latVal,
			Lng: lngVal,
		}

		places = append(places, p)
	}

	return places, nil
}

// GetPlaceDetail возвращает подробную информацию об одном месте по его ID
func (db *PostgresDB) GetPlaceDetail(placeID int, lat, long float64) (models.Place, error) {
	query := `
        SELECT t.name, p.id, c.name AS city_name, p.name, ST_AsText(p.geom) as geom, p.descr, p.price, country.currency,
               ST_Distance(p.geom::geography, ST_SetSRID(ST_MakePoint($2, $3), 4326)::geography) AS distance,
               ST_Y(p.geom::geometry) AS latitude,
    			ST_X(p.geom::geometry) AS longitude
        FROM app_place p
        JOIN app_city c ON p.city_id = c.id
        JOIN app_country country ON c.country_id = country.id
        JOIN app_place_type t ON p.type_id = t.id
        WHERE p.id = $1
    `

	var place models.Place
	var latVal, lngVal float64
	photos, _ := db.GetPhotosByPlaceID(placeID, 10) // лимит фоток для карточки 10

	err := db.DB.QueryRow(query, placeID, long, lat).Scan(
		&place.Type,
		&place.ID,
		&place.CityName,
		&place.Name,
		&place.Geom,
		&place.Desc,
		&place.Price,
		&place.Currency,
		&place.Distance,
		&latVal,
		&lngVal,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Place{}, fmt.Errorf("place with ID %d not found", placeID)
		}
		return models.Place{}, fmt.Errorf("failed to query place detail: %w", err)
	}

	place.Coordinates = models.Coordinates{
		Lat: latVal,
		Lng: lngVal,
	}

	place.Photos = photos

	return place, nil
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

		photos, err := db.GetPhotosByPlaceID(place.ID, 1) // задаем лимит фото для каждой карточки
		if err != nil {
			return nil, err
		}
		place.Photos = photos
		place.CityName = cityName
		place.Geom = geomText
		places = append(places, place)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating places: %w", err)
	}

	return places, nil
}

func (db *PostgresDB) RepoFavoritesPlaces(user_id int) ([]models.Place, error) {
	query := `
        SELECT place.id, c.name, place.name, ST_AsText(place.geom) as geom, place.descr
        FROM app_place place
        JOIN app_user user ON place.id = user.id
    `

	rows, err := db.DB.Query(query)
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

		photos, err := db.GetPhotosByPlaceID(place.ID, 1) // задаем лимит фото для каждой карточки
		if err != nil {
			return nil, err
		}
		place.Photos = photos
		place.CityName = cityName
		place.Geom = geomText
		places = append(places, place)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating places: %w", err)
	}

	return places, nil
}

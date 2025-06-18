package repository

import (
	"5Place/internal/models"
	"fmt"
	"os"
)

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

func (db *PostgresDB) GetPhotosByPlaceID(placeID int) ([]string, error) {
	query := fmt.Sprintf(`
		SELECT image FROM %s.photo WHERE place_id = $1
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

// GetAllCityPlaces выводит все места города
func (db *PostgresDB) GetAllCityPlaces(id int) ([]models.Place, error) {

	return nil, nil
}

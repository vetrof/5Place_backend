package repository

import (
	"5Place/internal/models"
	"fmt"
	"os"
)

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

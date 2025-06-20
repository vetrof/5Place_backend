package repository

import (
	"5Place/internal/models"
	"fmt"
)

func (db *PostgresDB) GetAllCities(country_id int) ([]models.City, error) {
	fmt.Println("country_id -->> ", country_id)

	query := fmt.Sprintf(`
	SELECT 
		c.id,
		c.name,
		ST_AsText(c.geom) as geom,
		COUNT(p.id) as points,
		co.name AS country_name
	FROM app_city c
	LEFT JOIN app_place p ON p.city_id = c.id
	JOIN app_country co ON c.country_id = co.id
	WHERE c.country_id = $1
	GROUP BY c.id, c.name, c.geom, co.name
	ORDER BY c.id
	LIMIT 20
`)

	rows, err := db.DB.Query(query, country_id)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var cities []models.City
	for rows.Next() {
		var c models.City
		if err := rows.Scan(&c.ID, &c.Name, &c.Geom, &c.Points, &c.Country); err != nil {
			return nil, fmt.Errorf("row scan error: %w", err)
		}
		cities = append(cities, c)
	}

	return cities, nil
}

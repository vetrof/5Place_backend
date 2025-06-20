package repository

import (
	"5Place/internal/models"
	"fmt"
)

func (db *PostgresDB) GetCountries() ([]models.Country, error) {

	query := fmt.Sprintf(`
	SELECT 
		c.id,
		c.name
	FROM app_country c
`)

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var countries []models.Country
	for rows.Next() {
		var c models.Country
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, fmt.Errorf("row scan error: %w", err)
		}
		countries = append(countries, c)
	}

	return countries, nil
}

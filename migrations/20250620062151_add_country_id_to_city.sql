-- +goose Up
ALTER TABLE app_city
ADD COLUMN country_id INTEGER;

ALTER TABLE app_city
ADD CONSTRAINT fk_city_country
FOREIGN KEY (country_id) REFERENCES app_country(id);

-- +goose Down
ALTER TABLE app_city
DROP CONSTRAINT fk_city_country;

ALTER TABLE app_city
DROP COLUMN country_id;
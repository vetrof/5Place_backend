5Place

![](banner.png)

***

dev mode: 

for start postgis-db:  
```docker compose up -d```

.env:  
```
DB_HOST=localhost
DB_PORT=55000
DB_USER=postgres
DB_NAME=place5
DB_PASSWORD=postgrespw
DB_SCHEMA=public
PORT=8080
```

start project:  
```
go mod tidy
go run cmd/api/main.go
```

connect to database in ide: 
```jdbc:postgresql://localhost:55000/place5```

sql for test insert places
```
INSERT INTO city (name)
VALUES ('Astana')
ON CONFLICT (name) DO NOTHING;
```
```
INSERT INTO place (city_name, name, geom, descr)
VALUES (
  'Astana',
  'Independence Square',
  ST_GeogFromText('SRID=4326;POINT(71.4304 51.1169)'),
  'Главная площадь Астаны — площадь Независимости'
);
```

```
INSERT INTO place (city_name, name, geom, descr)
VALUES (
  'Astana',
  'Independence Square',
  ST_GeogFromText('SRID=4326;POINT(71.419738 51.154179)'),
  'центральный парк Астаны'
);
```

в тестовом запросе передаем свои координаты  
```GET {{domain}}/near_place?long=71.408771&lat=51.162030```


адмика

django-admin .env
```
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_NAME=place5
DB_PASSWORD=postgrespw
```
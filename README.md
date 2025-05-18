5Place
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
```-- 1. Добавим центр Астаны
INSERT INTO city (name)
VALUES ('Astana')
ON CONFLICT (name) DO NOTHING;

-- 2. Добавим главную площадь Астаны
INSERT INTO place (city_name, name, geom, descr)
VALUES (
  'Astana',
  'Independence Square',
  ST_GeogFromText('SRID=4326;POINT(71.4304 51.1169)'),
  'Главная площадь Астаны — площадь Независимости'
);
```

```-- 1. Добавим центральный парк Астаны
INSERT INTO city (name)
VALUES ('Astana')
ON CONFLICT (name) DO NOTHING;

-- 2. Добавим главную площадь Астаны
INSERT INTO place (city_name, name, geom, descr)
VALUES (
  'Astana',
  'Independence Square',
  ST_GeogFromText('SRID=4326;POINT(71.419738 51.154179)'),
  'центральный парк Астаны'
);
```
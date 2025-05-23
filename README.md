# 5Place

![Banner](banner.png)

---

## 🚧 Dev Mode

### 📦 Запуск PostGIS

```bash
docker compose up -d
```

---

### 🔐 Переменные окружения (`.env`)

```env
DB_HOST=localhost
DB_PORT=55000
DB_USER=postgres
DB_NAME=place5
DB_PASSWORD=postgrespw
DB_SCHEMA=public
PORT=8080
```

---

### 🚀 Запуск проекта

```bash
go mod tidy
go run cmd/api/main.go
```

---

## 🧠 Подключение к базе данных в IDE

```
jdbc:postgresql://localhost:55000/place5
```

---

## 🧪 SQL для тестов

### ➕ Добавить город

```sql
INSERT INTO city (name)
VALUES ('Astana')
ON CONFLICT (name) DO NOTHING;
```

### ➕ Добавить место (без координат)

```sql
INSERT INTO place (city_name, name, geom, descr)
VALUES (
  'Astana',
  'Главная площадь Астаны — площадь Независимости',
  NULL,
  NULL
);
```

### ➕ Добавить место с координатами

```sql
INSERT INTO place (city_name, name, geom, descr)
VALUES (
  'Astana',
  'Independence Square',
  ST_GeogFromText('SRID=4326;POINT(71.419738 51.154179)'),
  'центральный парк Астаны'
);
```

---

## 📍 Тестовые запросы

Передаем свои координаты в запросе:


### Get near place
GET http://127.0.0.1:8080/near_place?long=71.108771&lat=51.962030

### All cities
GET http://127.0.0.1:8080/city


---

## 🔐 Админка Directus

[http://127.0.0.1:8055/](http://localhost:8055/)
```
EMAIL: 'admin@example.com'
PASSWORD: 'password'
```
# 5Place 
## Go | Python | Flutter | Postgres | Docker | Django | Minio | Chi 

![Banner](docs/banner.png)

Интерфейс админки
![Admin](docs/admin_interface.png)

---

## 🌍 О проекте 5Place

**5Place** — это open-source проект для тех, кто любит открывать новое рядом с собой. Приложение помогает туристам и жителям города узнавать об интересных и малоизвестных местах вокруг.

📍 **Вы заходите — и сразу видите 5 интересных локаций рядом с вами.**  
Маршруты строятся автоматически, а во время прогулки можно:

- 📖 почитать историю места,
- 📷 посмотреть фотографии,
- 🎧 послушать аудио или 🎥 посмотреть видеорассказ — на разных языках.

Все материалы создаются **исключительно людьми** — никакого AI-контента. Мы верим, что настоящий интерес скрыт в деталях, которые известны лишь местным жителям и опытным гидам.

---

## 🤝 Контент от людей для людей

- Все **фото**, **тексты**, **аудио** и **видео** сделаны вручную.
- Места подбираются **по рекомендации гидов и местных знатоков**.
- Мы стремимся показать **уникальные точки на карте**, которые часто теряются в шуме туристических гигантов.

---

## 🧭 Будущее приложения

- 📌 Курируемые маршруты по интересам: пабы, памятники, видовые точки, скверы и многое другое.
- 🗺️ Совместные проекты с профессиональными гидами.
- 🧩 Возможность настраивать количество отображаемых локаций, фильтры и категории интересов.

---

**Проект в активной разработке** и распространяется как **open-source** — присоединяйтесь, если хотите помочь сделать города ближе, интереснее и человечнее.

---
---

### APP
ссылка на мобильное приложение (flutter)  https://github.com/vetrof/5place_flutter

### DB Diagram
![db](docs/db_diagram.png)

## Структура проекта

```text
├── admin
│   ├── _settings
│   └── places
│       └── migrations
├── cmd
├── docs
├── internal
│   ├── place
│   │   ├── handlers
│   │   ├── models
│   │   ├── repository
│   │   │   └── mocks
│   │   ├── router
│   │   ├── services
│   │   └── utils
│   │       ├── gps
│   │       ├── logger
│   │       └── validator
│   └── user
│       └── router
└── migrations

```


# 🚧🚧🚧🚧🚧🚧🚧🚧🚧 Dev Mode 🚧🚧🚧🚧🚧🚧🚧🚧

### 🔐 Переменные окружения (`.env`)

env main project
```env main project
#REPO=fake

DB_HOST=localhost
DB_PORT=55000
DB_USER=postgres
DB_NAME=place5
DB_PASSWORD=postgrespw
DB_SCHEMA=public
PORT=8080

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-change-in-production-12345
JWT_EXPIRATION_HOURS=24
```
env django admin
```env django admin
HOST=localhost:9000

DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_NAME=place5
DB_PASSWORD=postgrespw
DB_SCHEMA=public

AWS_ACCESS_KEY_ID=minioadmin
AWS_SECRET_ACCESS_KEY=minioadmin
AWS_STORAGE_BUCKET_NAME=5place-files
AWS_S3_ENDPOINT_URL=http://minio:9000
AWS_S3_FILE_OVERWRITE=False
AWS_S3_ADDRESSING_STYLE=path
AWS_DEFAULT_ACL=None
```

если в .env REPO=fake то базу запускать ненужно

---

### 🚀 Запуск проекта

```bash
// запуск контейнеров
docker compose up --build -d
goose -dir migrations postgres "postgres://postgres:postgrespw@localhost:55000/place5?sslmode=disable&search_path=public" up

// запуск админки
docker compose exec web python manage.py migrate
docker compose exec web python manage.py createsuperuser

// запуск api
go run cmd/api/main.go
```
## Джанго админка
http://127.0.0.1:8000/admin

## 🧠 Подключение к базе данных в IDE

```
jdbc:postgresql://localhost:55000/place5
```  

## миграции goose
```
go install github.com/pressly/goose/v3/cmd/goose@latest
goose -dir migrations create create_users_table sql
goose -dir migrations postgres "postgres://postgres:postgrespw@localhost:55000/place5?sslmode=disable&search_path=public" up
```

## генерация swagger
```json
swag init -g cmd/main.go       

```

---

## 🧪 Наполнение базы для тестов

```$
PGPASSWORD=postgrespw psql -U postgres -h localhost -p 55000 -d place5 -f docs/init_data.sql

```

---
---

# 📍 Public API

## swagger

http://localhost:8080/swagger/

```json
### api info
GET {{domain}}/

### health
GET {{domain}}/health

### auth/register
POST {{domain}}/auth/register
Content-Type: application/json

{
"username": "testuser",
"email": "test@example.com",
"password": "password123"
}

### auth/login
POST {{domain}}/auth/login
Content-Type: application/json

{
"email": "test@example.com",
"password": "password123"
}

### auth/profile
GET {{domain}}/auth/profile
Authorization: Bearer "token"

### countries
GET {{domain}}/api/v1/place/countries

### cities
GET {{domain}}/api/v1/place/cities/country/1

### near 5 places
GET {{domain}}/api/v1/place/places/near?long=71.408771&lat=51.162030&limit=5&radius=5000

### place detail
GET {{domain}}/api/v1/place/places/1

### random places
GET {{domain}}/api/v1/place/places/random

### random places country
GET {{domain}}/api/v1/place/places/random?country=1

### random places city
GET {{domain}}/api/v1/place/places/random?city=2

### places list in city
GET {{domain}}/api/v1/place/places/city/1

```


```
### all cities
GET {{domain}}/api/v1//place/cities
```
```json
{
  "data": [
    {
      "id": 1,
      "name": "Astana",
      "geom": "POINT(71.429745 51.128479)",
      "points": 2
    }
  ],
  "meta": {
    "count": 0,
    "limit": 0,
    "searchRadius": 0,
    "center": {
      "lat": 0,
      "lon": 0
    }
  }
}

```

```
### near 5 places
GET {{domain}}/api/v1//place/places/near?long=71.408771&lat=51.162030
```
```json
{
  "data": [
    {
      "id": 2,
      "cityName": "Astana",
      "name": "central park",
      "geom": "POINT(71.419953 51.154506)",
      "desc": "центральный парк Астаны",
      "distance": 1145.69542435,
      "photos": [
        "https://astana.citypass.kz/wp-content/uploads/7db97aa358c9dcf7b27cd405bceba5e3.jpeg"
      ]
    },
    {
      "id": 1,
      "cityName": "Astana",
      "name": "Independence Square",
      "geom": "POINT(71.429745 51.128479)",
      "desc": "центральная площадь",
      "distance": 4010.78532212,
      "photos": [
        "https://media-cdn.tripadvisor.com/media/photo-s/0b/89/fb/fc/caption.jpg"
      ]
    }
  ],
  "meta": {
    "count": 2,
    "limit": 5,
    "searchRadius": 5000,
    "center": {
      "lat": 51.16203,
      "lon": 71.408771
    }
  }
}
```

```
### place detail
GET {{domain}}/api/v1//place/places/1
```
```json
{
  "data": [
    {
      "id": 1,
      "cityName": "Astana",
      "name": "Independence Square",
      "geom": "POINT(71.429745 51.128479)",
      "desc": "центральная площадь",
      "distance": 0,
      "photos": null
    }
  ],
  "meta": {
    "count": 0,
    "limit": 0,
    "searchRadius": 0,
    "center": {
      "lat": 0,
      "lon": 0
    }
  }
}
```

```
### places list in city
GET {{domain}}/api/v1//place/places/city/1
```
```json
{
  "data": [
    {
      "id": 2,
      "cityName": "Astana",
      "name": "central park",
      "geom": "POINT(71.419953 51.154506)",
      "desc": "центральный парк Астаны",
      "distance": 0,
      "photos": null
    },
    {
      "id": 1,
      "cityName": "Astana",
      "name": "Independence Square",
      "geom": "POINT(71.429745 51.128479)",
      "desc": "центральная площадь",
      "distance": 0,
      "photos": null
    }
  ],
  "meta": {
    "count": 0,
    "limit": 0,
    "searchRadius": 0,
    "center": {
      "lat": 0,
      "lon": 0
    }
  }
}
```




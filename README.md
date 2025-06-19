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

# 🚧🚧🚧🚧🚧🚧🚧🚧🚧 Dev Mode 🚧🚧🚧🚧🚧🚧🚧🚧

### 🔐 Переменные окружения (`.env`)

```env
REPO=fake

DB_HOST=localhost
DB_PORT=55000
DB_USER=postgres
DB_NAME=place5
DB_PASSWORD=postgrespw
DB_SCHEMA=public
PORT=8080
```

если в .env REPO=fake то базу запускать ненужно

---

### 🚀 Запуск проекта

```bash
docker compose up --build -d
goose -dir migrations postgres "postgres://postgres:postgrespw@localhost:55000/place5?sslmode=disable&search_path=public" up
docker compose exec web python manage.py migrate
docker compose exec web python manage.py createsuperuser
go run cmd/api/main.go
```


## 🧠 Подключение к базе данных в IDE

```
jdbc:postgresql://localhost:55000/place5
```  

## миграции
```
go install github.com/pressly/goose/v3/cmd/goose@latest
goose -dir migrations create create_users_table sql
goose -dir migrations postgres "postgres://postgres:postgrespw@localhost:55000/place5?sslmode=disable&search_path=public" up
```
---

## Джанго админка
http://127.0.0.1:8000/admin

---

## 🧪 Наполнение базы для тестов

### ➕ Добавить город + 2 места + фото на каждое место

```sql
INSERT INTO city (name, geom)
VALUES ('Astana', ST_GeogFromText('SRID=4326;POINT(71.429745 51.128479)'))
ON CONFLICT (name) DO NOTHING;

INSERT INTO place (city_id, name, geom, descr)
VALUES (
           1,
           'Independence Square',
           ST_GeogFromText('SRID=4326;POINT(71.429745 51.128479)'),
           'центральная площадь'
       );

INSERT INTO place (city_id, name, geom, descr)
VALUES (
           1,
           'central park',
           ST_GeogFromText('SRID=4326;POINT(71.419953 51.154506)'),
           'центральный парк Астаны'
       );


INSERT INTO photo (place_id, image, description)
VALUES (
           1,
           'https://media-cdn.tripadvisor.com/media/photo-s/0b/89/fb/fc/caption.jpg',
           'центральнаня площадь'
       );

INSERT INTO photo (place_id, image, description)
VALUES (
           2,
           'https://astana.citypass.kz/wp-content/uploads/7db97aa358c9dcf7b27cd405bceba5e3.jpeg',
           'центральный парк Астаны'
       );

```

---
---

# 📍 Public API

## 🔹 `GET /cities` — All Cities

**Request:**
```
GET {{domain}}/cities
```

**Response:**
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

---

## 🔹 `GET /places/near` — Near 5 Places

**Request:**
```
GET {{domain}}/places/near?long=71.408771&lat=51.162030
```

**Response:**
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

---

## 🔹 `GET /places/{id}` — Place Detail

**Request:**
```
GET {{domain}}/places/1
```

**Response:**
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

---

## 🔹 `GET /places/city/{city_id}` — Places in City

**Request:**
```
GET {{domain}}/places/city/1
```

**Response:**
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




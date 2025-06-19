# 5Place 
## Go | Python | Dart | Postgres | Docker | Django | Minio | Chi 

![Banner](docs/banner.png)

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


### TODO
 -[x] "Чистая" архитектура для проекта  
 -[x] Приложение на flutter для проекта  
 -[x] Таблицы для юзера м избранного  
 -[x] Подключение к БД через интерфейс  
 -[ ] Подключить миграции  
 -[ ] Избранное  
 -[ ] Посещенное  
 -[ ] Поля для ссылок на видео и аудио  
 -[ ] S3 хранилище для изображений (на входе оптимизировать)  
 -[ ] Подключить кеш  
 -[ ] Таблицу для маршрутов  
 -[ ] Таблицу для моих точек (сохранять свое положение)  
 -[ ] Авторизация через телегу или whatsapp  
 -[ ] Добавить таблицу - рейтинг  

### APP
ссылка на мобильное приложение (flutter)  https://github.com/vetrof/5place_flutter

### DB Diagram
![db](docs/db_diagram.png)

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
---
---
---
---
---
# 🚧🚧🚧🚧🚧🚧🚧🚧🚧 Dev Mode 🚧🚧🚧🚧🚧🚧🚧🚧
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
генерация моделей (в проекте уже сгенерированы)  
```docker compose exec web python manage.py inspectdb > places/models.py```

Прописать модели джанги в базу
```bash
docker compose exec web python manage.py migrate
```

Создать суперюзера
```bash
docker compose exec web python manage.py createsuperuser
```

затем
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

## 📍 Тестовые запросы

```json
### all cities
GET {{domain}}/cities

### near 5 places
GET {{domain}}/places/near?long=71.408771&lat=51.162030

### place detail
GET {{domain}}/places/1

### places list in city
GET {{domain}}/places/city/1
```

# 5Place

![Banner](banner.png)

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

## 🧪 Наполнение базы для тестов

### ➕ Добавить город + 2 места + фото на каждое место

```sql
INSERT INTO city (name)
VALUES ('Astana')
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


INSERT INTO photo (place_id, file_link, description)
VALUES (
           1,
           'https://media-cdn.tripadvisor.com/media/photo-s/0b/89/fb/fc/caption.jpg',
           'центральнаня площадь'
       );

INSERT INTO photo (place_id, file_link, description)
VALUES (
           2,
           'https://astana.citypass.kz/wp-content/uploads/7db97aa358c9dcf7b27cd405bceba5e3.jpeg',
           'центральный парк Астаны'
       );

```

---

## 📍 Тестовые запросы

GET http://127.0.0.1:8080/near_place?long=71.108771&lat=51.962030
```json
[
  {
    "ID": 2,
    "CityName": "1",
    "Name": "house",
    "Geom": "POINT(71.40779381461866 51.16246940914874)",
    "Desc": "house",
    "Distance": 84.03784534
  },
  {
    "ID": 1,
    "CityName": "1",
    "Name": "park",
    "Geom": "POINT(71.41844700046002 51.15587254065716)",
    "Desc": "park",
    "Distance": 963.05044711
  },
  {
    "ID": 3,
    "CityName": "2",
    "Name": "mountain",
    "Geom": "POINT(77.05954207441113 43.203151137847215)",
    "Desc": "mountain",
    "Distance": 982318.60610853
  }
]

```

GET http://127.0.0.1:8080/city
```json
[
  {
    "ID": 1,
    "Name": "astana",
    "Geom": "POINT(71.41042574459857 51.15162433549682)",
    "Points": 2
  },
  {
    "ID": 2,
    "Name": "almaty",
    "Geom": "POINT(76.93752703214142 43.25406013672736)",
    "Points": 1
  }
]

```

---

## 🔐 Админка Directus

[http://127.0.0.1:8055/](http://localhost:8055/)
```
EMAIL: 'admin@example.com'
PASSWORD: 'password'
```
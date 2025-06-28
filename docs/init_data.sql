-- init_data.sql

INSERT INTO app_country (name, currency)
VALUES ('Uzbekistan', 'SOM')
ON CONFLICT (name) DO NOTHING;

INSERT INTO app_city (name, geom, country_id)
VALUES ('Astana3', ST_GeogFromText('SRID=4326;POINT(71.429745 51.128479)'), 1)
ON CONFLICT (name) DO NOTHING;

INSERT INTO app_place_type (name)
VALUES ('Bridge')
ON CONFLICT (name) DO NOTHING;

INSERT INTO app_place (city_id, name, geom, descr, type_id)
VALUES (
           1,
           'Independence Square',
           ST_GeogFromText('SRID=4326;POINT(71.429745 51.128479)'),
           'центральная площадь',
           1
       );

INSERT INTO app_place (city_id, name, geom, descr, type_id)
VALUES (
           1,
           'central park',
           ST_GeogFromText('SRID=4326;POINT(71.419953 51.154506)'),
           'центральный парк Астаны',
           1
       );

INSERT INTO app_photo (place_id, image, description)
VALUES (
           1,
           'https://media-cdn.tripadvisor.com/media/photo-s/0b/89/fb/fc/caption.jpg',
           'центральнаня площадь'
       );

INSERT INTO app_photo (place_id, image, description)
VALUES (
           2,
           'https://astana.citypass.kz/wp-content/uploads/7db97aa358c9dcf7b27cd405bceba5e3.jpeg',
           'центральный парк Астаны'
       );

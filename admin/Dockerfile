# Dockerfile для Django с PostGIS
FROM python:3.11-slim

ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1

# Устанавливаем зависимости для PostGIS и сборки Python пакетов
RUN apt-get update && apt-get install -y \
    binutils \
    libproj-dev \
    gdal-bin \
    libgdal-dev \
    libgeos-dev \
    libpq-dev \
    gcc \
    && rm -rf /var/lib/apt/lists/*

# Устанавливаем рабочую директорию
WORKDIR /code

# Копируем requirements.txt и устанавливаем зависимости
COPY ./requirements.txt /code/
RUN pip install --upgrade pip \
    && pip install -r requirements.txt

# Копируем весь проект
COPY . /code/

# Открываем порт
EXPOSE 8000

# Стартовая команда (замени на gunicorn, если нужен прод-режим)
CMD ["python", "manage.py", "runserver", "0.0.0.0:8000"]
# go-todo-api

Это тестовый проект API, написанный на Go, с использованием Docker и Docker Compose.

## Установка и запуск

1. Склонируйте репозиторий:

   ```sh
   git clone https://github.com/Llifintsefv/go-todo-api.git
   ```

2. Файл `.env` в корневой директории со следующим содержимым:

   ```env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=mysecretpassword
    DB_NAME=postgres-db
    DB_SSL_MODE=disable
    APP_PORT=:8080

   ```

   Вы можете изменить значения по своему усмотрению.

3. Соберите и запустите приложение с помощью Docker Compose:

   ```sh
    docker build --no-cache -t go-todo-api .
    docker compose --env-file ./cmd/api/.env up -d --build
   ```

4. API будет доступно по адресу: `http://localhost:8080`.

## Дополнительно

- Миграции базы данных автоматически применяются при запуске.


#  URL Shortener API 

Это проект сервиса сокращения ссылок, реализованный на языке Go с использованием стандартной библиотеки net/http, JWT, Docker и PostgreSQL.

## Стек технологий

- Golang
- PostgreSQL
- Docker Compose
- JWT

## Запуск приложения

### 1. Клонируем репозиторий

```bash
git clone https://github.com/laxyshkaaa/url-shortener
```

### 2. Создаём `.env` файл

Скопируй шаблон:

```bash
cp .env.example .env
```

### 3. Запускаем через Docker

```bash
docker-compose up --build
```

Сервис будет доступен на `http://localhost:8080`.


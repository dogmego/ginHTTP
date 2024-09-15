# Используем обновленный образ Go 1.22 для сборки приложения
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код в контейнер
COPY . .

# Сборка приложения
RUN go build -o main ./cmd/main.go

# Используем стабильный образ Alpine для запуска приложения
FROM alpine:3.17

WORKDIR /root/

# Копируем бинарный файл из предыдущего контейнера
COPY --from=builder /app/main .
COPY --from=builder /app/schema ./schema

# Обновляем индексы пакетов и устанавливаем необходимые утилиты
RUN apk update && apk add --no-cache bash curl netcat-openbsd

# Устанавливаем утилиту migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate /usr/local/bin/ \
    && chmod +x /usr/local/bin/migrate

# Копируем скрипт entrypoint
COPY entrypoint.sh .

# Копируем директорию configs и файл .env
COPY configs ./configs
COPY .env .env

# Делаем скрипт исполняемым
RUN chmod +x entrypoint.sh

# Открываем порт для приложения
EXPOSE 8000

# Указываем точку входа
ENTRYPOINT ["./entrypoint.sh"]


#!/bin/sh

# Ожидание доступности базы данных
until nc -z -v -w30 $DB_HOST $DB_PORT
do
  echo "Ожидание запуска базы данных..."
  sleep 1
done

echo "База данных доступна, применяем миграции..."

# Применяем миграции
migrate -path ./schema -database "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" up

echo "Миграции применены, запускаем приложение..."

# Запускаем приложение
./main

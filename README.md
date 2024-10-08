**ToDo List API на Go и PostgreSQL**
Этот проект представляет собой RESTful API для управления списками задач (ToDo List), разработанный на языке Go с использованием фреймворка Gin и базы данных PostgreSQL.

**Запуск локально в docker контейнере**
1. Скачать docker
2. Склонировать к себе репозиторий
3. В корневой директории написать команду "docker-compose up --build"

**Функциональность**
1. Аутентификация и авторизация пользователей:
2. Регистрация новых пользователей.
3. Авторизация с использованием JWT токенов.

**Управление списками задач:**
1. Создание новых списков.
2. Получение списка всех созданных списков.
3. Получение конкретного списка по его ID.
4. Обновление информации о списке.
5. Удаление списка.

**Управление задачами внутри списков:**
1. Добавление задач в список.
2. Получение всех задач из списка.
3. Получение конкретной задачи по ее ID.
4. Обновление информации о задаче.
5. Удаление задачи из списка.

**Технологический стек**
- Язык программирования: Go.
- Веб-фреймворк: Gin.
- База данных: PostgreSQL.
- Работа с БД: sqlx.
- Аутентификация: JWT (JSON Web Tokens).
- Тестирование API: Postman или любой другой инструмент для API-тестирования.

**Структура проекта**
- cmd/main.go: Точка входа в приложение.
- pkg/handler: Обработчики HTTP-запросов.
- pkg/service: Бизнес-логика приложения.
- pkg/repository: Взаимодействие с базой данных.
- maxHTTP: Общие модели и структуры данных.

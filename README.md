# Тестовое задание на позицию стажера в Avito

Этот проект создан в рамках тестового задания на позицию DevOps стажера в Avito. Он представляет собой минималистичное приложение-интерфейс для работы с Redis с проксированием трафика через Nginx.

## Функционал

Проект реализует следующий функционал:

1. Создание пары ключ-значение в Redis по заданному ключу.
2. Получение значения по ключу из Redis (возвращает 404, если ключ отсутствует).
3. Удаление пары ключ-значение из Redis.
4. Проксирование трафика через Nginx, порт 8089.

## Технологии

Проект реализован с использованием следующих технологий:

- Golang для реализации приложения.
- Redis для хранения данных.
- Nginx для проксирования трафика.
- Docker и Docker Compose для контейнеризации и управления компонентами проекта.

## Требования

Для работы с этим проектом вам понадобятся:

- Docker
- Docker Compose

## Развертывание

1. Склонируйте репозиторий:

```sh
git clone git@github.com:the-cutest-ogre-ever/go-redis-nginx-app.git
```

2. Перейдите в каталог проекта:
```sh
cd go-redis-nginx-app
```
3. Соберите и запустите контейнеры для приложения, Redis и Nginx с помощью Docker Compose:
```sh
docker-compose up --build
```

## Использование

После успешного развертывания, вы можете использовать следующие эндпоинты:

- Создать или перезаписать пару ключ-значение: POST http://app:8089/set_key
- Получить значение по ключу: GET http://app:8089/get_key?key=<key>
- Удалить пару ключ-значение: POST http://app:8089/del_key

### Примеры использования эндпоинтов с *curl*
1. Создание или перезапись пары ключ-значение
```sh
curl -X POST -H "Content-Type: application/json" -d '{"my_key": "my_value"}' http://app:8089/set_key
```
2. Получение значения по ключу
```sh
curl http://app:8089/get_key?key=my_key
```
3. Удаление пары ключ-значение
```sh
curl -X POST -H "Content-Type: application/json" -d '{"key": "my_key"}' http://app:8089/del_key
```

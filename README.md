# Lamoda: тестовое задание
---
## Описание задачи:
### #1. API для работы с товарами на складе
Необходимо спроектировать и реализовать API методы для работы с товарами на одном складе. 
Учесть, что вызов API может быть одновременно из разных систем и они могут работать с одинаковыми товарами.
Методы API можно расширять доп. параметрами на своё усмотрение

Спроектировать и реализовать БД для хранения следующих сущностей

- Склад

1) название
2) признак доступности

- Товар

1) название
2) размер
3) уникальный код
4) количество

Реализовать методы API:

- Резервирование товара на складе для доставки. На вход принимает:

1) массив уникальных кодов товара

- Освобождение резерва товаров. На вход принимает

1) массив уникальных кодов товара

- получение кол-ва оставшихся товаров на складе. На вход принимает:

1) идентификатор склада

- Будет плюсом
  - Реализация логики работы с товарами, которые одновременно могут находиться на нескольких складах

---

# Запуск приложения:

### Приложение поднимается через docker-compose, что позволит сразу поднять все необходимые компоненты сервиса: само приложение, postgres и мигратор

```bash
docker compose up
```

### Для отключения: 

```bash
docker compose down
```

---

# Тестирование API

## Запросы

### Примеры запросов можно найти в файле `.http` 

## Examples

1) **Reserve products request**

``` .http
#ReserveProducts
POST <http://localhost:8081/v1/products/reserve>
Content-Type: application/json

{
  "unique_codes": ["ABC123", "DEF456"]
}
```

- Response

```
{
	"Message": "Products reserved successfully"
}
```

2) **Release reservations request**

``` .http
# ReleaseReservations
POST <http://localhost:8081/v1/products/release>
Content-Type: application/json
```

- Response

```
{
	"Message": "Reservations released successfully"
}
```

3) **Release reservations request**

``` .http
{
  "unique_codes": ["ABC123", "DEF456"]
}

# GetRemainingProducts
GET <http://localhost:8081/v1/products/remaining/1>
```

- Response

```
{
    "count":175
}
```

### Также, создан swagger файл в *(лежит одноименной папке)*

---

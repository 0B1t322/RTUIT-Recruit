# Система продажи и производства товаров
## Описание
Система состоит из 3 микросервисов:
1. Сервис покупок
2. Сервис магазинов
3. Заводы

Сервис покупок позволяет совершать покупки в магазинах и отслеживать их в виде чека

Сервис магазины позволяет отслеживтаь имеющие товары, добавлять новые 

Сервис заводы производит эти товары и доставляет их в магазины

Каждый из сервисов доступен по одному порту указаному в docker-compose.yml

Сервис магазины обращается к сервису покупки, чтобы совершить покупку
Сервис заводы обращается к магазинам, чтобы доставлять в них товары

Сам по себе сервис не имеет какого либо интерфейса для взаимодействия с ним
Он конфигурируется на примере:
```json
{
    // список продуктов которые будут созданы
    "products": [
        {
            "name": "iPhone X 256 GB",
            "description": "Apple iphone with 256 GB",
            "cost": 70000,
            "category": "Phones"
        },
        {
            "name": "iPhone X 512 GB",
            "description": "Apple iphone with 512 GB",
            "cost": 90000,
            "category": "Phones"
        },
        {
            "name": "MacBook 17 ARM 1TB",
            "description": "MacBook 16 ARM with 32GB RAM",
            "cost": 200000,
            "category": "Laptops"
        }
    ],
    // время указываеться в формате целого числа с припиской "m" "s" "ms" "h"

    // через какое время создавать продукцию 
    "production_capisity": "15s",
    // сколько продукции делать за это время
    "product_per_cap": 2,
    // Сеть до магазина
    "shop_network": "http://service.shops:8082",
    // как часто запрашивать список магазигов и искать в них производимые товары
    "update_time": "3s",
    // как часто пытаться доставить продукт
    "delivery_time": "2s"
}
```

По умолчанию nginx прослушивает порт 8084

## Запуск
для запуска необходимо
1. [Docker]("https://www.docker.com")

Чтобы запустить с исходного кода необходимо клонировать репозиторий и прописать следующие команды
```
    docker-compose pull
    docker-compose up
```

Для остановки проекта нажимте Ctrl^C и пропишите
```
    docker-compose down
```

Для редактирование проекта необходимо
1. [Golang 1.15](https://golang.org/dl/)


Также по http://localhost:8001 доступен swaggerUI

# API for level_2
# service.purchases
---
## GET /purchases/{uid}/{id} 
return a json purchase according to id in path
if not find return 404 code and empty body
if success return code 200
### Example
/purchases/1/1
```json
    {
        "id": 1,
        "uid": 1,
        "shop_id": 1,
        "shop": {
            "ID": 1,
            "name": "shop_1",
            "adress": "shop_1_adress",
            "phone_number": "89991234567"
        },
        "buy_date": "2021-03-07T14:38:09.476Z",
        "product_id": 1,
        "product": {
            "ID": 1,
            "name": "product_1",
            "description": "product_1_desc",
            "cost": 1000,
            "category": "products"
        },
        "payment": "cash",
        "count": 9,
        "cost": 9000
    }
```
    
## GET /purchases/{uid}
return all purchases for uid
if not find any return code 404
if success return code 200
### Example
/purchases/1
```json
    [
        {
            "id": 1,
            "uid": 1,
            "shop_id": 1,
            "shop": {
                "ID": 1,
                "name": "shop_1",
                "adress": "shop_1_adress",
                "phone_number": "89991234567"
            },
            "buy_date": "2021-03-07T14:38:09.476Z",
            "product_id": 1,
            "product": {
                "ID": 1,
                "name": "product_1",
                "description": "product_1_desc",
                "cost": 1000,
                "category": "products"
            },
            "payment": "cash",
            "count": 9,
            "cost": 9000
        },
        {
            "id": 2,
            "uid": 1,
            "shop_id": 1,
            "shop": {
                "ID": 1,
                "name": "shop_1",
                "adress": "shop_1_adress",
                "phone_number": "89991234567"
            },
            "buy_date": "2021-03-07T14:42:53.339Z",
            "product_id": 1,
            "product": {
                "ID": 1,
                "name": "product_1",
                "description": "product_1_desc",
                "cost": 1000,
                "category": "products"
            },
            "payment": "card",
            "count": 1,
            "cost": 1000
        }
    ]
```
## POST /purchases/{uid}
Add to db a purchases with uid in path
if add return code 201 and uint id of added purchase
require a Authorization header with key if token not found return 401
if count bigger than count of product or zero return 401
don't reduce a count of product make this from shops api
### Example
/purchases/1
body req:
```json
    {
        "product_id": 1,
        "shop_id": 1,
        "count": 9,
        "payment": "cash"
    }
```
answer:
```
1
```
## DELETE /purchases/{uid}/{id}
Delete a purchase with id in path
if not found purchase with this id return code 404
If success retorn code 200
### Example
/purchases/1/1

---
# service.shops


## GET /shops/{id}
return a information abou shop
if not find shop with this id return 404
### Example
/shops/1
answer:
```json
{
    "ID": 1,
    "name": "shop_1",
    "adress": "shop_1_adress",
    "phone_number": "89991234567",
    "shop_products": [
        {
            "shop_id": 1,
            "product_id": 1,
            "product": {
                "ID": 1,
                "name": "product_1",
                "description": "product_1_desc",
                "cost": 1000,
                "category": "products"
            },
            "count": 10,
            "UpdatedAt": "0001-01-01T00:00:00Z"
        }
    ]
}
```
## PUT /shops/{id}/{pid}
buy product
if shop or product not found return 404
if success returb code 200
if not found uid or count or payment in body return 400
if body empty return 400
### Example
/shops/1/1
req body:
```json
{
    "uid": 1,
    "count": 2,
    "payment": "cash"
}
```
## GET /shops/purchases/{uid}
return all purchases for user
if not find return 404
### Example
/shops/purchases/1
answer:
```json
[
    {
        "id": 1,
        "uid": 1,
        "shop_id": 1,
        "shop": {
            "ID": 1,
            "name": "shop_1",
            "adress": "shop_1_adress",
            "phone_number": "89991234567"
        },
        "buy_date": "2021-03-07T14:38:09.476Z",
        "product_id": 1,
        "product": {
            "ID": 1,
            "name": "product_1",
            "description": "product_1_desc",
            "cost": 1000,
            "category": "products"
        },
        "payment": "cash",
        "count": 9,
        "cost": 9000
    },
    {
        "id": 2,
        "uid": 1,
        "shop_id": 1,
        "shop": {
            "ID": 1,
            "name": "shop_1",
            "adress": "shop_1_adress",
            "phone_number": "89991234567"
        },
        "buy_date": "2021-03-07T14:42:53.339Z",
        "product_id": 1,
        "product": {
            "ID": 1,
            "name": "product_1",
            "description": "product_1_desc",
            "cost": 1000,
            "category": "products"
        },
        "payment": "card",
        "count": 1,
        "cost": 1000
    },
    {
        "id": 3,
        "uid": 1,
        "shop_id": 1,
        "shop": {
            "ID": 1,
            "name": "shop_1",
            "adress": "shop_1_adress",
            "phone_number": "89991234567"
        },
        "buy_date": "2021-03-07T15:16:57.961Z",
        "product_id": 1,
        "product": {
            "ID": 1,
            "name": "product_1",
            "description": "product_1_desc",
            "cost": 1000,
            "category": "products"
        },
        "payment": "cash",
        "count": 2,
        "cost": 2000
    }
]
```

## PUT /shops/{id}/{pid}/{count}
add count to product in shop
if success return 200 
if don't find a product or shop return 404
require a header Authirization with token
### Example
/shops/1/1/2

## POST /shops/{id}/{pid}
add product to shop
if don't find product or shop return 404
require a header Authirization with token
if find return 200
### Example
/shops/1/2

## GET /shops
return all shops

### Example
/shops
```json
[
    {
        "ID": 1,
        "name": "shop_1",
        "adress": "shop_1_adress",
        "phone_number": "89991234567",
        "shop_products": [
            {
                "shop_id": 1,
                "product_id": 1,
                "product": {
                    "ID": 1,
                    "name": "product_1",
                    "description": "product_1_desc",
                    "cost": 1000,
                    "category": "products"
                },
                "count": 10,
                "UpdatedAt": "2021-03-07T15:35:47.971Z"
            },
            {
                "shop_id": 1,
                "product_id": 2,
                "product": {
                    "ID": 2,
                    "name": "product_2",
                    "description": "product_2_desc",
                    "cost": 2000,
                    "category": "products"
                },
                "count": 0,
                "UpdatedAt": "2021-03-07T15:38:44.691Z"
            }
        ]
    }
]
```

## POST /shops/
Create shop
if shop with this name exist return 400
if create return status 200
if not found name in body return 400
if body empty return 400
### Example
/shops/
body:
```json
{
    "name": "name_of_some_shop_1",
    "adress": "adress_of_some_shop",
    "phone_number": "89991234567"
}
```
answer:
```
15
```

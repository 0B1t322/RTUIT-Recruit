# API for level_2

## GET /purchases/:uid/:id 
return a json purchase according to id in path
if not find return 404 code and empty body
if success return code 200
require a Authorization header with key
### Example
/purchases/1/1
```json
    {
        "id": 1,
        "uid": 1,
        "shop_id": 9,
        "shop": {
            "ID": 9,
            "name": "cool_shop",
            "adress": "adress_1",
            "phone_number": "89991234567"
        },
        "buy_date": "2021-02-27T18:32:47.959Z",
        "product_id": 29,
        "product": {
            "ID": 29,
            "name": "phone_1",
            "description": "cool phone",
            "cost": 13000,
            "category": "Phone"
        },
        "payment": "",
        "count": 1
    }
```
    
## GET /purchases/:uid
return all purchases for uid
if not find any return code 404
if success return code 200
require a Authorization header with key
### Example
/purchases/1
```json
    [
        {
            "id": 1,
            "uid": 1,
            "shop_id": 9,
            "shop": {
                "ID": 0,
                "name": "",
                "adress": "",
                "phone_number": ""
            },
            "buy_date": "2021-02-27T18:32:47.959Z",
            "product_id": 29,
            "product": {
                "ID": 0,
                "name": "",
                "description": "",
                "cost": 0,
                "category": ""
            },
            "payment": "",
            "count": 1
        },
        {
            "id": 2,
            "uid": 1,
            "shop_id": 9,
            "shop": {
                "ID": 0,
                "name": "",
                "adress": "",
                "phone_number": ""
            },
            "buy_date": "2021-02-27T18:34:45.919Z",
            "product_id": 29,
            "product": {
                "ID": 0,
                "name": "",
                "description": "",
                "cost": 0,
                "category": ""
            },
            "payment": "",
            "count": 5
        }
    ]
```
## POST /purchases/:uid
Add to db a purchases with uid in path
if add return code 201 and uint id of added purchase
require a Authorization header with key
### Example
/purchases/1
body req:
```json
    {
        "product_id": 29,
        "shop_id": 9,
        "cost": 99923,
        "count": 5,
        "payment": "",
    }
```

## DELETE /purchases/:uid/:id
Delete a purchase with id in path
if not found purchase with this id return code 404
If success retorn code 200
require a Authorization header with key
### Example
/purchases/1/1
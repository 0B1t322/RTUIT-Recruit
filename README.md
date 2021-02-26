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
        "ID": 1,
        "CreatedAt": "2021-02-18T19:51:42Z",
        "UpdatedAt": "2021-02-18T19:51:42Z",
        "DeletedAt": null,
        "UID": "1",
        "buy_date": "2021-02-18T19:51:41.999Z",
        "product_name": "prodcut_1",
        "cost": 123.2234567
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
    {
        [
            {
                "ID": 1,
                "CreatedAt": "2021-02-18T19:51:42Z",
                "UpdatedAt": "2021-02-18T19:51:42Z",
                "DeletedAt": null,
                "UID": "1",
                "buy_date": "2021-02-18T19:51:41.999Z",
                "product_name": "prodcut_1",
                "cost": 123.2234567
            },
            {
                "ID": 2,
                "CreatedAt": "2021-02-18T19:54:42Z",
                "UpdatedAt": "2021-02-18T19:54:42Z",
                "DeletedAt": null,
                "UID": "1",
                "buy_date": "2021-02-18T19:54:41.999Z",
                "product_name": "prodcut_2",
                "cost": 124
            },
        ],
    }
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
        "product_name": "product_1",
        "cost": 213,	
    }
```

## DELETE /purchases/:uid/:id
Delete a purchase with id in path
if not found purchase with this id return code 404
If success retorn code 200
require a Authorization header with key
### Example
/purchases/1/1
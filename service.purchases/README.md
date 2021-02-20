# API for level_1

## GET /purchases/:uid/:id 
return a json purchase according to id in path
if not find return 404 code and empty body

### Example
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
    
## GET /purchases/:uid
return all purchases for uid
if not find any return code 404

### Example

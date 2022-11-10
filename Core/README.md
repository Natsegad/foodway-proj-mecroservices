# Api
Get
```url
/store/getStores
```
Ответ
```json
[
    {
        "store_name": "Гасан жиес",
        "store_id": 42402557
    },
    {
        "store_name": "Гасан жие",
        "store_id": 2349912723
    }
]
```
Create store
Post
```url
/store/create
```

Тело
```json
{
    "store_name":"Дерьмо"
}
```

Ответ 
```json
{
    "store_name": "Дерьмо",
    "store_id": 3327271853
}
```

Products

POST 
Create product
```url
/products/create
```

Тело
```json
{
    "ID":0,
    "store_id":0,
    "store_name":"Дерьмо",
    "name":"Govno",
    "price":22,
    "image_path":"asdasd.jpg",
    "grams":2323,
    "calories":213,
    "sqrls":213123,
    "fats":123,
    "carbonts":213123
}
```

Ответ
```json
"Дерьмо"
```

Get products in store by store id 
Get
```url
/products/getProducts?id=3327271853
```

Ответ
```json
[
    {
        "ID": 3,
        "store_id": 3327271853,
        "store_name": "Дерьмо",
        "name": "Govno",
        "price": 22,
        "image_path": "asdasd.jpg",
        "grams": 2323,
        "calories": 213,
        "sqrls": 213123,
        "fats": 123,
        "carbonts": 213123
    },
    {
        "ID": 4,
        "store_id": 3327271853,
        "store_name": "Дерьмо",
        "name": "Gvno",
        "price": 22,
        "image_path": "asdasd.jpg",
        "grams": 2323,
        "calories": 213,
        "sqrls": 213123,
        "fats": 123,
        "carbonts": 213123
    }
]
```
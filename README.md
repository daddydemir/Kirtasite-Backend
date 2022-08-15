# Kırtasiye Projesi


> Bu projede kullanılan veritabanı büyük ölçüde değiştirildiği için proje [bu adreste](https://github.com/daddydemir/Kirtasite) geliştirilmeye devam etmektedir.

Bu proje;

- Mobil
- Backend
- Frontend

'den oluşmaktadır.

## API Docs

**USER**

1. GetAllUsers
    - Method : GET
    - `/api/users`
2. GetUserById
    - Method : GET
    - `/api/users/{id}`
3. AddUser
    - Method : POST
    - `/api/users`
4. UpdateUser
    - Method : PUT
    - `/api/users/{id}`
5. DeleteUser
    - Method : DELETE
    - `/api/users/{id}`
6. UpdateImage
    - Method : PUT
    - `/api/users/{id}`

### Request Body

``` json
{
    "username":"",
    "password":"",
    "image_path":"",
    "mail":"",
    "tel_no":"",
    "role_id":0
}
```

**STATIONERY**

1. GetAllStationery
    - Method : GET
    - `/api/stationery`
2. GetStationeryById
    - Method : GET
    - `/api/stationery/{id}`
3. AddStationery
    - Method : POST
    - `/api/stationery`
4. UpdateStationery
    - Method : PUT
    - `/api/stationery/{id}`
5. DeleteStationery
    - Method : DELETE
    - `/api/stationery/{id}`
6. UpdateStationery
    - Method : PUT
    - ``/api/stationery/{id}`

### Request Body

```json
{
    "role_id":0,
    "company_name":"",
    "password":"",
    "image_path":"",
    "mail":"",
    "tel_no":"",
    "address_id":0,
    "score":0
}
```

**ROLES**

1. GetAllRoles
    - Method : GET
    - `/api/roles`

### Response Body

```json
{
    "id":""
    "name":""
}
```

**FILES**

1. GetFileByUserId
    - Method : GET
    - `/api/files/{id}`
2. FileDelete
    - Method : DELETE
    - `/api/files/{id}`
3. FileAdd
    - Method : POST
    - `/api/files`
4. FileById
    - Method : GET
    - `/api/file/{id}`

### Request Body

```json
{
    "user_id":0,
    "private":true,
    "file_path":""
}
```

**PRICES**

1. GetAllPrices
    - Method : GET
    - `/api/prices`
2. PriceById
    - Method : GET
    - `/api/prices/{id}`
3. PriceAdd
    - Method : POST
    - `/api/prices`
4. PriceUpdate
    - Method : PUT
    - `/api/prices/{id}`
5. PriceDelete
    - Method : DELETE
    - `/api/prices/{id}`

### Response Body

```json
{
    "stationery_id":0,
    "info":"",
    "price":0
}
```

**ORDERS**

1. OrderByUserId
    - Method : GET
    - `/api/orders/user/{id}`
2. OrderByStationeryId
    - Method : GET
    - `/api/orders/stationery/{id}`
3. OrderAdd
    - Method : POST
    - `/api/orders/`

### Request Body

```json
{
    "file_id":0,
    "user_id":0,
    "stationery_id":0,
    "delivery_date":"",
    "verification_code":"",
    "total_price":0
}
```

**COMMENTS**

1. GetCommentByUserId
    - Method : GET
    - `/api/comments/user/{id}`
2. CommentByStationeryId
    - Method : GET
    - `/api/comments/stationery/{id}`
3. CommentAdd
    - Method : POST
    - `/api/comments/`

### Request Body

```json
{
    "user_id":0,
    "stationery_id":0,
    "content":"",
    "date":"",
    "score":0
}
```

**ADDRESS**

1. GetAddressById
    - Method : GET
    - `/api/locate/{id}`
2. GetAddressByCity
    - Method : GET
    - `/api/city/{id}`
3. GetAddressByDistrict
    - Method : POST
    - `/api/district/{id}`

### Request Body 

``` json
{
    "id":0,
    "city_id":0,
    "district_id":0,
    "name":"",
    "cordinate_x":"",
    "cordinate_y":""
}
```

**CITY**
1. GetCityById
    - Method : GET
    - `/api/cities/{id}`

### Request Body
``` json
{
    "id":0,
    "name":""
}
```

**DISTRICT**
1. GetDistrictById
    - Method : GET
    - `/api/districts/{id}`
2. GetDistrictByCity
    - Method : GET
    - `/api/district-city/{id}`

### Response Body
```json
{
    "id":0,
    "city_id":0,
    "name":""
}
```
> NOT: Request Body POST ve PUT metodunda kullanılacak

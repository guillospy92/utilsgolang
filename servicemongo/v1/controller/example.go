package controller

var jsonCreated = `{
  "subscription_id": "zzz-zzz-zzz",
  "service_id": "0101",
  "type": "prime",
  "country_code": "co",
  "city_code": "br",
  "timezone": "America/Bogota",
  "activation_date": "2022-01-17T10:57:05-05:00",
  "expiration_date": "2022-02-17T10:57:05-05:00",
  "is_renovated": true,
  "cancel_date": null,
  "expiration_date": null,
  "status": "active",
  "user": {
    "id": 22,
    "name": "name",
    "email": "name@gmail.com",
    "cellphone": "300000000",
    "address": "address"
  },
  "plan": {
    "id": "xxx-xxx-xxx",
    "name": "year",
    "price": 5000.34,
    "promotional_price": 0,
    "period": 3,
    "benefits": [
      "free deliver",
      "discount total"
    ]
  }
}`

var jsonCharge = `{
  "subscription_id": "xx-xx-xxx",
  "is_renovated": false,
  "payment": {
    "card_id": 23433453453,
    "card_type": "t.c",
    "last_four": 2345,
    "installments": 4
  }
}`

var jsonCancel = `{
  "subscription_id": "xx-xx-xxx",
  "is_renovated": true,
  "cancel_date": "2023-01-17T10:57:05-05:00",										
  "status": "cancel 2"
}`

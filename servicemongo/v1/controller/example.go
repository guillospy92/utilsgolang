package controller

var jsonCreated = `{
  "subscription_id": "zzz-zzz-zzz-ggg-lll",
  "service_id": "01010101010",
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
  "subscription_id": "621e4bc855d790e9d6958e37-co-dd0595cf-80dd-41a9-a8ab-93b4b16669af",
  "is_renovated": false,
  "payment": {
    "credit_card": {
      "id": 387926,
      "franchise": "visa",
      "last_four": 1234,
      "installments": 1
    },
    "history": [
      {
        "transaction_id": "eyetyetyter37373676",
        "transaction_status": "pending_payment",
        "transaction_date":"2022-01-17T10:57:05-05:00",
        "reference_code": "string"
      }
    ]
  }
}`

var jsonCancel = `{
  "subscription_id": "xx-xx-xxx",
  "is_renovated": true,
  "cancel_date": "2023-01-17T10:57:05-05:00",										
  "status": "cancel 2"
}`

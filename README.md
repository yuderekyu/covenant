# Covenant
[![Build Status](https://travis-ci.org/yuderekyu/covenant.svg?branch=master)](https://travis-ci.org/yuderekyu/covenant)
[![Coverage Status](https://coveralls.io/repos/github/yuderekyu/covenant/badge.svg?branch=master)](https://coveralls.io/github/yuderekyu/covenant?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/yuderekyu/covenant)](https://goreportcard.com/report/github.com/yuderekyu/covenant)

A Go service handling Expresso subscriptions

## API

### Subscription

#### `POST /api/subscription` creates and adds a new subscription to the database

Example:

*Request:*
```
{
  "userId" : "52e02c0a-f480-11e6-89f2-0242ac13000b",
  "frequency" : "MONTHLY",
  "roasterId" : "5afc1d5c-f562-11e6-a55c-0242ac130009",
  "itemId" : "9d97b574-f487-11e6-bc64-92361f002671",
  "quantity" : 1
}  
```

*Response:*
```
{
	"data": {
	   "userId" : "52e02c0a-f480-11e6-89f2-0242ac13000b",
       "frequency" : "MONTHLY",
       "roasterId" : "5afc1d5c-f562-11e6-a55c-0242ac130009",
       "itemId" : "9d97b574-f487-11e6-bc64-92361f002671",
       "quantity" : 1
	}
}
```

#### `GET /api/subscription/:subscriptionId` returns the subscription with the given subscriptionId

Example:

*Request:*
```
GET localhost:8082/api/subscription/397eb95d-1993-11e7-b75e-0a0027000004
```

*Response*
```
{
  "data": {
    "id": "397eb95d-1993-11e7-b75e-0a0027000004",
    "userId": "52e02c0a-f480-11e6-89f2-0242ac13000b",
    "status": "ACTIVE",
    "createdAt": "2017-04-05T00:02:57Z",
    "frequency": "MONTHLY",
    "roasterId": "5afc1d5c-f562-11e6-a55c-0242ac130009",
    "itemId": "9d97b574-f487-11e6-bc64-92361f002671",
    "quantity": 1
  },
  "success": true
}
```

#### `GET /api/subscription/` returns a list of subscriptions 

Example:

*Request:*
```
GET localhost:8082/api/subscription?offset=0&limit=3
```

*Response:*
```
{
  "data": [
    {
      "id": "0652eaa8-1a39-11e7-b56e-0a0027000004",
      "userId": "52e02c0a-f480-11e6-89f2-0242ac13000b",
      "status": "ACTIVE",
      "createdAt": "2017-04-05T19:49:47Z",
      "frequency": "MONTHLY",
      "roasterId": "5afc1d5c-f562-11e6-a55c-0242ac130009",
      "itemId": "9d97b574-f487-11e6-bc64-92361f002671",
      "quantity": 1
    },
    {
      "id": "1013cc94-1a39-11e7-b107-0a0027000004",
      "userId": "52e02c0a-f480-11e6-89f2-0242ac13000b",
      "status": "ACTIVE",
      "createdAt": "2017-04-05T19:50:04Z",
      "frequency": "MONTHLY",
      "roasterId": "5afc1d5c-f562-11e6-a55c-0242ac130009",
      "itemId": "9d97b574-f487-11e6-bc64-92361f002671",
      "quantity": 1
    },
    {
      "id": "1690b5f1-fbc7-11e6-b6da-0242ac130008",
      "userId": "52e02c0a-f480-11e6-89f2-0242ac13000b",
      "status": "ACTIVE",
      "createdAt": "2017-02-26T01:58:37Z",
      "frequency": "DAILY",
      "roasterId": "5afc1d5c-f562-11e6-a55c-0242ac130009",
      "itemId": "9d97b574-f487-11e6-bc64-92361f002671",
      "quantity": 0
    },
  ],
  "success": true
}
```

#### `GET /api/roaster/subscription/:roasterId` returns a list of subscriptions with the given roasterId

Example:

*Request:*
```
GET localhost:8082/api/roaster/subscription/5afc1d5c-f562-11e6-a55c-0242ac130009?offset=0&limit=2
```

*Response:*
```
{
  "data": [
    {
      "id": "0652eaa8-1a39-11e7-b56e-0a0027000004",
      "userId": "52e02c0a-f480-11e6-89f2-0242ac13000b",
      "status": "ACTIVE",
      "createdAt": "2017-04-05T19:49:47Z",
      "frequency": "MONTHLY",
      "roasterId": "5afc1d5c-f562-11e6-a55c-0242ac130009",
      "itemId": "9d97b574-f487-11e6-bc64-92361f002671",
      "quantity": 1
    },
    {
      "id": "1013cc94-1a39-11e7-b107-0a0027000004",
      "userId": "52e02c0a-f480-11e6-89f2-0242ac13000b",
      "status": "ACTIVE",
      "createdAt": "2017-04-05T19:50:04Z",
      "frequency": "MONTHLY",
      "roasterId": "5afc1d5c-f562-11e6-a55c-0242ac130009",
      "itemId": "9d97b574-f487-11e6-bc64-92361f002671",
      "quantity": 1
    }
  ],
  "success": true
}
```

#### `GET /api/user/subscription/:userId` returns a list of subscriptions with the given userId

Example:

*Request:*
```
GET localhost:8082/api/user/subscription/52e02c0a-f480-11e6-89f2-0242ac13000b?offset=1&limit=2
```

*Response:*
```
{
  "data": [
    {
      "id": "1013cc94-1a39-11e7-b107-0a0027000004",
      "userId": "52e02c0a-f480-11e6-89f2-0242ac13000b",
      "status": "ACTIVE",
      "createdAt": "2017-04-05T19:50:04Z",
      "frequency": "MONTHLY",
      "roasterId": "5afc1d5c-f562-11e6-a55c-0242ac130009",
      "itemId": "9d97b574-f487-11e6-bc64-92361f002671",
      "quantity": 1
    },
    {
      "id": "1690b5f1-fbc7-11e6-b6da-0242ac130008",
      "userId": "52e02c0a-f480-11e6-89f2-0242ac13000b",
      "status": "ACTIVE",
      "createdAt": "2017-02-26T01:58:37Z",
      "frequency": "DAILY",
      "roasterId": "5afc1d5c-f562-11e6-a55c-0242ac130009",
      "itemId": "9d97b574-f487-11e6-bc64-92361f002671",
      "quantity": 0
    }
  ],
  "success": true
}
```

#### `PUT /api/subscription/:subscriptionId` updates the subscription with the given subscriptionId

Example:

*Request:*
```
{
"id":"952e029f-f9fe-11e6-8253-0a002700001b",
"status": "TEST",
"frequency" : "TEST",
"roasterID": "5afc1d5c-f562-11e6-a55c-0242ac130009",
"itemID": "9d97b574-f487-11e6-bc64-92361f002671",
"quantity": 1
}
```

*Response:*
```
{
  "data": {
    "id": "952e029f-f9fe-11e6-8253-0a002700001b",
    "userId": "",
    "status": "TEST",
    "createdAt": "0001-01-01T00:00:00Z",
    "frequency": "TEST",
    "roasterId": "5afc1d5c-f562-11e6-a55c-0242ac130009",
    "itemId": "9d97b574-f487-11e6-bc64-92361f002671",
    "quantity": 1
  },
  "success": true
}
```
#### `DELETE /api/subscription/:subscriptionId` deletes the subscription with the given subscriptionId

Example:

*Request:*
```
DELETE localhost:8082/api/subscription/952e029f-f9fe-11e6-8253-0a002700001b
```

*Response:*
```
{
  "data": null,
  "success": true
}
```

#### `POST /api/order` creates an order with the given userID and itemID

Example:

*Request:*
```
{
"userId": "5afc1d5c-f562-11e6-a55c-0242ac130009",
"itemId": "9d97b574-f487-11e6-bc64-92361f002671",
"quantity": 1
}
```

*Response:*
```
{
  "id":"",
  "userId":"",
  "subscriptionId":"",
  "requestDate":"",
  "shipDate":"",
  "quantity":"",
  "status":"",
  "labelUrl":""
}
```
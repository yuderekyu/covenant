# Covenant
[![Build Status](https://travis-ci.org/yuderekyu/covenant.svg?branch=master)](https://travis-ci.org/yuderekyu/covenant)
[![Coverage Status](https://coveralls.io/repos/github/yuderekyu/covenant/badge.svg?branch=master)](https://coveralls.io/github/yuderekyu/covenant?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/yuderekyu/covenant)](https://goreportcard.com/report/github.com/yuderekyu/covenant)

A Go service handling Expresso subscriptions

## API

### Subscription

#### `POST /api/subscription` creates and adds a new subscription to the database

#### `GET /api/subscription/:subscriptionId` returns the subscription with the given subscriptionId

#### `GET /api/subscription/` returns a list of subscriptions 

#### `GET /api/roaster/subscription/:roasterId` returns a list of subscriptions with the given roasterId

#### `GET /api/user/subscription/:userId` returns a list of subscriptions with the given userId

#### `PUT /api/subscription/:subscriptionId` updates the subscription with the given subscriptionId

#### `DELETE /api/subscription/:subscriptionId` deletes the subscription with the given subscriptionId

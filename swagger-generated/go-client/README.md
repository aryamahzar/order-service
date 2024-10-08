# Go API client for swagger

API for managing orders in the e-commerce platform

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8080/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**CancelOrder**](docs/DefaultApi.md#cancelorder) | **Post** /orders/{orderId}/cancel | Cancel an order
*DefaultApi* | [**CreateOrder**](docs/DefaultApi.md#createorder) | **Post** /orders | Create a new order
*DefaultApi* | [**GetOrder**](docs/DefaultApi.md#getorder) | **Get** /orders/{orderId} | Get an order by ID
*DefaultApi* | [**ListOrders**](docs/DefaultApi.md#listorders) | **Get** /orders | List orders
*DefaultApi* | [**UpdateOrder**](docs/DefaultApi.md#updateorder) | **Put** /orders/{orderId} | Update an order

## Documentation For Models

 - [Address](docs/Address.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [Order](docs/Order.md)
 - [OrderCreate](docs/OrderCreate.md)
 - [OrderItem](docs/OrderItem.md)
 - [OrderStatus](docs/OrderStatus.md)
 - [OrderUpdate](docs/OrderUpdate.md)

## Documentation For Authorization

## apiKeyAuth
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author



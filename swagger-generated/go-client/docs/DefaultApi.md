# {{classname}}

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrder**](DefaultApi.md#CancelOrder) | **Post** /orders/{orderId}/cancel | Cancel an order
[**CreateOrder**](DefaultApi.md#CreateOrder) | **Post** /orders | Create a new order
[**GetOrder**](DefaultApi.md#GetOrder) | **Get** /orders/{orderId} | Get an order by ID
[**ListOrders**](DefaultApi.md#ListOrders) | **Get** /orders | List orders
[**UpdateOrder**](DefaultApi.md#UpdateOrder) | **Put** /orders/{orderId} | Update an order

# **CancelOrder**
> Order CancelOrder(ctx, orderId)
Cancel an order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | [**string**](.md)|  | 

### Return type

[**Order**](Order.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrder**
> Order CreateOrder(ctx, body)
Create a new order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrderCreate**](OrderCreate.md)|  | 

### Return type

[**Order**](Order.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrder**
> Order GetOrder(ctx, orderId)
Get an order by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | [**string**](.md)|  | 

### Return type

[**Order**](Order.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrders**
> InlineResponse200 ListOrders(ctx, optional)
List orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiListOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListOrdersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | [**optional.Interface of string**](.md)|  | 
 **status** | [**optional.Interface of OrderStatus**](.md)|  | 
 **page** | **optional.Int32**|  | [default to 1]
 **size** | **optional.Int32**|  | [default to 20]

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrder**
> Order UpdateOrder(ctx, body, orderId)
Update an order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrderUpdate**](OrderUpdate.md)|  | 
  **orderId** | [**string**](.md)|  | 

### Return type

[**Order**](Order.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


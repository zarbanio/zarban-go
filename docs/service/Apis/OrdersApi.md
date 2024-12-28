# OrdersApi

All URIs are relative to *https://api.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getUnfilledOrders**](OrdersApi.md#getUnfilledOrders) | **GET** /v2/orders | Fetch Unfilled Orders |
| [**syncOrder**](OrdersApi.md#syncOrder) | **POST** /v2/orders/sync | Updates Order Entity |


<a name="getUnfilledOrders"></a>
# **getUnfilledOrders**
> OrderResponse getUnfilledOrders(type, hash, status, offerer, filler, decayStartTime, decayEndTime, deadline, cursor, limit)

Fetch Unfilled Orders

    Get a list of unfilled orders filtered by different parameters.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **type** | **String**| Type of the order | [optional] [default to null] [enum: limit, dutch] |
| **hash** | **String**| order hash | [optional] [default to null] |
| **status** | **String**| Status of the order | [optional] [default to null] [enum: open, expired, error, cancelled, filled, insufficient-funds] |
| **offerer** | **String**| Ethereum address of the offerer | [optional] [default to null] |
| **filler** | **String**| Ethereum address of the filler | [optional] [default to null] |
| **decayStartTime** | [**TimeRange**](../Models/.md)| Decay start time | [optional] [default to null] |
| **decayEndTime** | [**TimeRange**](../Models/.md)| Decay end time | [optional] [default to null] |
| **deadline** | [**TimeRange**](../Models/.md)| Order deadline | [optional] [default to null] |
| **cursor** | **Integer**| Cursor for pagination | [optional] [default to null] |
| **limit** | **Integer**| Limit the number of orders returned (default is 10) | [optional] [default to 10] |

### Return type

[**OrderResponse**](../Models/OrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="syncOrder"></a>
# **syncOrder**
> Error syncOrder(UpdateOrderRequest)

Updates Order Entity

    updates an order entity in database

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdateOrderRequest** | [**UpdateOrderRequest**](../Models/UpdateOrderRequest.md)|  | |

### Return type

[**Error**](../Models/Error.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


# PricesApi

All URIs are relative to *https://api.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**listPrices**](PricesApi.md#listPrices) | **GET** /v2/prices | List prices based on query parameters |


<a name="listPrices"></a>
# **listPrices**
> PriceListResponse listPrices(symbol)

List prices based on query parameters

    Retrieve a list of price details based on specified query parameters

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **symbol** | **String**| Symbol of the price | [optional] [default to null] |

### Return type

[**PriceListResponse**](../Models/PriceListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


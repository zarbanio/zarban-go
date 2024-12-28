# AddressesApi

All URIs are relative to *https://api.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getAllAddresses**](AddressesApi.md#getAllAddresses) | **GET** /v2/addresses | Get all addresses |


<a name="getAllAddresses"></a>
# **getAllAddresses**
> AddressResponse getAllAddresses(format)

Get all addresses

    All addresses knows to Zarban.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **format** | **String**| The type of addresses to return | [optional] [default to null] |

### Return type

[**AddressResponse**](../Models/AddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


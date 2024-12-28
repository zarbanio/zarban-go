# DepositsApi

All URIs are relative to *https://api.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getUserDeposits**](DepositsApi.md#getUserDeposits) | **GET** /v2/lendingpool/deposits | Get user deposits of Lendingpool |


<a name="getUserDeposits"></a>
# **getUserDeposits**
> UserDepositsResponse getUserDeposits(user, reserve, cursor, limit)

Get user deposits of Lendingpool

    Get user deposits of Lendingpool

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| Ethereum address of the user | [optional] [default to null] |
| **reserve** | **String**| Ethereum address of the reserve | [optional] [default to null] |
| **cursor** | **Integer**| Cursor for pagination | [optional] [default to null] |
| **limit** | **Integer**| Limit the number of deposits returned (default is 50) | [optional] [default to 50] |

### Return type

[**UserDepositsResponse**](../Models/UserDepositsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


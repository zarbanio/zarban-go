# BorrowsApi

All URIs are relative to *https://api.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getUserBorrows**](BorrowsApi.md#getUserBorrows) | **GET** /v2/lendingpool/borrows | Get user borrows of lendingpool |


<a name="getUserBorrows"></a>
# **getUserBorrows**
> UserBorrowsResponse getUserBorrows(user, reserve, cursor, limit)

Get user borrows of lendingpool

    Get user borrows of lendingpool

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| Ethereum address of the user | [optional] [default to null] |
| **reserve** | **String**| Ethereum address of the reserve | [optional] [default to null] |
| **cursor** | **Integer**| Cursor for pagination | [optional] [default to null] |
| **limit** | **Integer**| Limit the number of deposits returned (default is 50) | [optional] [default to 50] |

### Return type

[**UserBorrowsResponse**](../Models/UserBorrowsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


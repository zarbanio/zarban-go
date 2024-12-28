# AccountsApi

All URIs are relative to *https://api.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getAccountByAddress**](AccountsApi.md#getAccountByAddress) | **GET** /v2/accounts/{address} | Get account by address |


<a name="getAccountByAddress"></a>
# **getAccountByAddress**
> Account getAccountByAddress(address)

Get account by address

    Get account by address

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **address** | **String**| Ethereum address of the account | [default to null] |

### Return type

[**Account**](../Models/Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


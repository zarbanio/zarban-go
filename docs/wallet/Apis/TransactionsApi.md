# TransactionsApi

All URIs are relative to *https://wapi.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getUserTransactions**](TransactionsApi.md#getUserTransactions) | **GET** /transactions | Get user transactions |


<a name="getUserTransactions"></a>
# **getUserTransactions**
> TransactionResponse getUserTransactions(cursor, limit)

Get user transactions

    Get a list of the recent transactions of the user.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **cursor** | **Integer**| Cursor for pagination | [optional] [default to 0] |
| **limit** | **Integer**| Limit the number of transactions returned (default is 100) | [optional] [default to 100] |

### Return type

[**TransactionResponse**](../Models/TransactionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


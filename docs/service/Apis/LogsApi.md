# LogsApi

All URIs are relative to *https://api.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getLogsByTransactionHash**](LogsApi.md#getLogsByTransactionHash) | **GET** /v2/logs/{txHash} | Get raw and decoded logs by transaction hash |


<a name="getLogsByTransactionHash"></a>
# **getLogsByTransactionHash**
> EventDetailsResponse getLogsByTransactionHash(txHash)

Get raw and decoded logs by transaction hash

    Get raw and decoded logs by transaction hash

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **txHash** | **String**| Transaction hash | [default to null] |

### Return type

[**EventDetailsResponse**](../Models/EventDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


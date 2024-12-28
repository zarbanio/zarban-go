# WithdrawApi

All URIs are relative to *https://wapi.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getUserWithdrawRequests**](WithdrawApi.md#getUserWithdrawRequests) | **GET** /withdraws | Get user withdraw requests |
| [**getWithdrawalStatus**](WithdrawApi.md#getWithdrawalStatus) | **GET** /withdraws/{id} | Get withdrawal status |
| [**previewWithdrawal**](WithdrawApi.md#previewWithdrawal) | **POST** /withdraws/preview | Withdraw request |
| [**requestWithdrawal**](WithdrawApi.md#requestWithdrawal) | **POST** /withdraws/request | Withdraw request |


<a name="getUserWithdrawRequests"></a>
# **getUserWithdrawRequests**
> WithdrawRequestResponse getUserWithdrawRequests()

Get user withdraw requests

    Get user withdraw requests

### Parameters
This endpoint does not need any parameter.

### Return type

[**WithdrawRequestResponse**](../Models/WithdrawRequestResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getWithdrawalStatus"></a>
# **getWithdrawalStatus**
> WithdrawRequest getWithdrawalStatus(id)

Get withdrawal status

    Get the status of a withdrawal request.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | **Long**| Withdrawal request ID | [default to null] |

### Return type

[**WithdrawRequest**](../Models/WithdrawRequest.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="previewWithdrawal"></a>
# **previewWithdrawal**
> WithdrawRequestPreview previewWithdrawal(WithdrawRequestBody)

Withdraw request

    Request a withdrawal from the wallet.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **WithdrawRequestBody** | [**WithdrawRequestBody**](../Models/WithdrawRequestBody.md)|  | |

### Return type

[**WithdrawRequestPreview**](../Models/WithdrawRequestPreview.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="requestWithdrawal"></a>
# **requestWithdrawal**
> WithdrawResponseBody requestWithdrawal(WithdrawRequestBody)

Withdraw request

    Request a withdrawal from the wallet.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **WithdrawRequestBody** | [**WithdrawRequestBody**](../Models/WithdrawRequestBody.md)|  | |

### Return type

[**WithdrawResponseBody**](../Models/WithdrawResponseBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


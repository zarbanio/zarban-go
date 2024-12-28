# PermitApi

All URIs are relative to *https://api.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getSingleTokenPermit**](PermitApi.md#getSingleTokenPermit) | **GET** /v2/permit/single | Get permit for single token |


<a name="getSingleTokenPermit"></a>
# **getSingleTokenPermit**
> PermitSingle getSingleTokenPermit(token, user)

Get permit for single token

    Get permit for single token

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **token** | **String**| Ethereum address of the token | [default to null] |
| **user** | **String**| Ethereum address of the user | [default to null] |

### Return type

[**PermitSingle**](../Models/PermitSingle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


# DepositApi

All URIs are relative to *https://wapi.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**depositMoney**](DepositApi.md#depositMoney) | **GET** /deposit | Deposit money |


<a name="depositMoney"></a>
# **depositMoney**
> DepositResponse depositMoney(network, symbol)

Deposit money

    Deposit money to the wallet.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **network** | **String**| Network to deposit | [default to null] |
| **symbol** | **String**| Coin symbol | [default to null] |

### Return type

[**DepositResponse**](../Models/DepositResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


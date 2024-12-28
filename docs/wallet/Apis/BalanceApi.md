# BalanceApi

All URIs are relative to *https://wapi.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getBalanceBySymbol**](BalanceApi.md#getBalanceBySymbol) | **GET** /balance/{symbol} | Get balance |
| [**getWalletBalance**](BalanceApi.md#getWalletBalance) | **GET** /balance | Get wallet balance |


<a name="getBalanceBySymbol"></a>
# **getBalanceBySymbol**
> Balance getBalanceBySymbol(symbol)

Get balance

    Get the balance by a symbol.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **symbol** | [**Symbol**](../Models/.md)| Coin symbol | [default to null] [enum: USD, DAI, USDT, USDC, WBTC, ETH, WETH, ZAR, TMN] |

### Return type

[**Balance**](../Models/Balance.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getWalletBalance"></a>
# **getWalletBalance**
> WalletBalance getWalletBalance()

Get wallet balance

    Get the balance of the wallet.

### Parameters
This endpoint does not need any parameter.

### Return type

[**WalletBalance**](../Models/WalletBalance.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


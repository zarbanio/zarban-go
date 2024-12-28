# CoinsApi

All URIs are relative to *https://wapi.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getCoinDetails**](CoinsApi.md#getCoinDetails) | **GET** /coins/{symbol} | Get coin |
| [**getSupportedCoins**](CoinsApi.md#getSupportedCoins) | **GET** /coins | Get coins |


<a name="getCoinDetails"></a>
# **getCoinDetails**
> Coin getCoinDetails(symbol)

Get coin

    Get the details of a coin.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **symbol** | [**Symbol**](../Models/.md)| Coin symbol | [default to null] [enum: USD, DAI, USDT, USDC, WBTC, ETH, WETH, ZAR, TMN] |

### Return type

[**Coin**](../Models/Coin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getSupportedCoins"></a>
# **getSupportedCoins**
> CoinResponse getSupportedCoins()

Get coins

    Get the list of supported coins.

### Parameters
This endpoint does not need any parameter.

### Return type

[**CoinResponse**](../Models/CoinResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


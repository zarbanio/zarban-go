# LendingPoolApi

All URIs are relative to *https://api.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createLendingPoolBorrow**](LendingPoolApi.md#createLendingPoolBorrow) | **POST** /v2/lendingpool/tx/borrow | Borrow from lending pool |
| [**createLendingPoolDeposit**](LendingPoolApi.md#createLendingPoolDeposit) | **POST** /v2/lendingpool/tx/deposit | Deposit to lending pool |
| [**createLendingPoolRepay**](LendingPoolApi.md#createLendingPoolRepay) | **POST** /v2/lendingpool/tx/repay | Repay to lending pool |
| [**createLendingPoolWithdraw**](LendingPoolApi.md#createLendingPoolWithdraw) | **POST** /v2/lendingpool/tx/withdraw | Withdraw from lending pool |
| [**fetchReserveDataByAsset**](LendingPoolApi.md#fetchReserveDataByAsset) | **GET** /v2/lendingpool/reserves | Fetch Reserve Data By Asset |
| [**getUserBorrows**](LendingPoolApi.md#getUserBorrows) | **GET** /v2/lendingpool/borrows | Get user borrows of lendingpool |
| [**getUserDeposits**](LendingPoolApi.md#getUserDeposits) | **GET** /v2/lendingpool/deposits | Get user deposits of Lendingpool |
| [**setLendingPoolAssetCollateral**](LendingPoolApi.md#setLendingPoolAssetCollateral) | **POST** /v2/lendingpool/tx/useassetascollateral | Enable/Disable asset as collateral |


<a name="createLendingPoolBorrow"></a>
# **createLendingPoolBorrow**
> LendingpoolBorrowTxResponse createLendingPoolBorrow(LendingpoolBorrowTxRequest)

Borrow from lending pool

    Borrow from lending pool

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **LendingpoolBorrowTxRequest** | [**LendingpoolBorrowTxRequest**](../Models/LendingpoolBorrowTxRequest.md)| Borrow transaction request, if amount is not provided, it will be calculated based on the token balance | |

### Return type

[**LendingpoolBorrowTxResponse**](../Models/LendingpoolBorrowTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createLendingPoolDeposit"></a>
# **createLendingPoolDeposit**
> LendingpoolDepositTxResponse createLendingPoolDeposit(LendingpoolDepositTxRequest)

Deposit to lending pool

    Deposit to lending pool

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **LendingpoolDepositTxRequest** | [**LendingpoolDepositTxRequest**](../Models/LendingpoolDepositTxRequest.md)| Deposit transaction request, if amount is not provided, it will be calculated based on the token balance | |

### Return type

[**LendingpoolDepositTxResponse**](../Models/LendingpoolDepositTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createLendingPoolRepay"></a>
# **createLendingPoolRepay**
> LendingpoolRepayTxResponse createLendingPoolRepay(LendingpoolRepayTxRequest)

Repay to lending pool

    Repay to lending pool

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **LendingpoolRepayTxRequest** | [**LendingpoolRepayTxRequest**](../Models/LendingpoolRepayTxRequest.md)| Repay transaction request, if amount is not provided, it will repay the maximum possible amount | |

### Return type

[**LendingpoolRepayTxResponse**](../Models/LendingpoolRepayTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createLendingPoolWithdraw"></a>
# **createLendingPoolWithdraw**
> LendingpoolWithdrawTxResponse createLendingPoolWithdraw(LendingpoolWithdrawTxRequest)

Withdraw from lending pool

    Withdraw from lending pool

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **LendingpoolWithdrawTxRequest** | [**LendingpoolWithdrawTxRequest**](../Models/LendingpoolWithdrawTxRequest.md)| Withdraw transaction request, if amount is not provided, it will be calculated based on the user account status | |

### Return type

[**LendingpoolWithdrawTxResponse**](../Models/LendingpoolWithdrawTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="fetchReserveDataByAsset"></a>
# **fetchReserveDataByAsset**
> FormattedReserveData fetchReserveDataByAsset(asset)

Fetch Reserve Data By Asset

    Retrieve data of reserves

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **asset** | **String**| The asset address in hexadecimal format. | [optional] [default to null] |

### Return type

[**FormattedReserveData**](../Models/FormattedReserveData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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

<a name="setLendingPoolAssetCollateral"></a>
# **setLendingPoolAssetCollateral**
> LendingpoolUseAssetAsCollateralTxResponse setLendingPoolAssetCollateral(LendingpoolUseAssetAsCollateralTxRequest)

Enable/Disable asset as collateral

    Allows to enable/disable a specific asset as collateral

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **LendingpoolUseAssetAsCollateralTxRequest** | [**LendingpoolUseAssetAsCollateralTxRequest**](../Models/LendingpoolUseAssetAsCollateralTxRequest.md)| UseAssetAsCollateral transaction request | |

### Return type

[**LendingpoolUseAssetAsCollateralTxResponse**](../Models/LendingpoolUseAssetAsCollateralTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


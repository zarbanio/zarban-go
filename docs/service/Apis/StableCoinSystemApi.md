# StableCoinSystemApi

All URIs are relative to *https://api.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**approveAndJoinZarTransaction**](StableCoinSystemApi.md#approveAndJoinZarTransaction) | **POST** /v2/stablecoinsystem/auctions/tx/zarjoin | approve and join ZAR token into Vat contract |
| [**createStableCoinVault**](StableCoinSystemApi.md#createStableCoinVault) | **POST** /v2/stablecoinsystem/tx/createvault | Create vault |
| [**depositStableCoinCollateral**](StableCoinSystemApi.md#depositStableCoinCollateral) | **POST** /v2/stablecoinsystem/tx/depositcollateral | Deposit collateral |
| [**exitGemTransaction**](StableCoinSystemApi.md#exitGemTransaction) | **POST** /v2/stablecoinsystem/auctions/tx/gemexit | exit Gem token (which can be used as collateral) from Vat contract |
| [**exitZarTransaction**](StableCoinSystemApi.md#exitZarTransaction) | **POST** /v2/stablecoinsystem/auctions/tx/zarexit | exit ZAR token from Vat contract |
| [**getAllIlks**](StableCoinSystemApi.md#getAllIlks) | **GET** /v2/ilks | Get all Ilks |
| [**getCollectorData**](StableCoinSystemApi.md#getCollectorData) | **GET** /v2/stats | Get collector data |
| [**getIlkByName**](StableCoinSystemApi.md#getIlkByName) | **GET** /v2/ilks/{name} | Get Ilk by name |
| [**getVaultById**](StableCoinSystemApi.md#getVaultById) | **GET** /v2/vaults/{id} | Get a vault by ID |
| [**getVaultEventsById**](StableCoinSystemApi.md#getVaultEventsById) | **GET** /v2/vaults/{id}/events | Get vault events by ID |
| [**getVaultsByOwner**](StableCoinSystemApi.md#getVaultsByOwner) | **GET** /v2/vaults | Get vaults by owner query |
| [**liquidateVaultTransaction**](StableCoinSystemApi.md#liquidateVaultTransaction) | **POST** /v2/stablecoinsystem/tx/bark | liquidate a vault |
| [**mintZarTransaction**](StableCoinSystemApi.md#mintZarTransaction) | **POST** /v2/stablecoinsystem/tx/mintzar | Mint ZAR |
| [**repayZarTransaction**](StableCoinSystemApi.md#repayZarTransaction) | **POST** /v2/stablecoinsystem/tx/repayzar | Repay ZAR |
| [**resetAuctionTransaction**](StableCoinSystemApi.md#resetAuctionTransaction) | **POST** /v2/stablecoinsystem/auctions/tx/redo | reset a auction |
| [**takeAuctionTransaction**](StableCoinSystemApi.md#takeAuctionTransaction) | **POST** /v2/stablecoinsystem/auctions/tx/take | take a auction |
| [**withdrawCollateralTransaction**](StableCoinSystemApi.md#withdrawCollateralTransaction) | **POST** /v2/stablecoinsystem/tx/withdrawcollateral | Withdraw collateral |


<a name="approveAndJoinZarTransaction"></a>
# **approveAndJoinZarTransaction**
> ChainActivity approveAndJoinZarTransaction(StablecoinSystemZarjoinTxRequest)

approve and join ZAR token into Vat contract

    approve and join ZAR token into Vat contract

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **StablecoinSystemZarjoinTxRequest** | [**StablecoinSystemZarjoinTxRequest**](../Models/StablecoinSystemZarjoinTxRequest.md)| Approve and join Zar token into Vat contract transaction request | |

### Return type

[**ChainActivity**](../Models/ChainActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createStableCoinVault"></a>
# **createStableCoinVault**
> ChainActivity createStableCoinVault(StablecoinSystemCreateVaultTxRequest)

Create vault

    Create vault

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **StablecoinSystemCreateVaultTxRequest** | [**StablecoinSystemCreateVaultTxRequest**](../Models/StablecoinSystemCreateVaultTxRequest.md)| Create vault transaction request | |

### Return type

[**ChainActivity**](../Models/ChainActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="depositStableCoinCollateral"></a>
# **depositStableCoinCollateral**
> ChainActivity depositStableCoinCollateral(StablecoinSystemDepositCollateralTxRequest)

Deposit collateral

    Deposit collateral

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **StablecoinSystemDepositCollateralTxRequest** | [**StablecoinSystemDepositCollateralTxRequest**](../Models/StablecoinSystemDepositCollateralTxRequest.md)| Deposit collateral transaction request | |

### Return type

[**ChainActivity**](../Models/ChainActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="exitGemTransaction"></a>
# **exitGemTransaction**
> ChainActivity exitGemTransaction(StablecoinSystemGemexitTxRequest)

exit Gem token (which can be used as collateral) from Vat contract

    exit Gem token (which can be used as collateral) from Vat contract

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **StablecoinSystemGemexitTxRequest** | [**StablecoinSystemGemexitTxRequest**](../Models/StablecoinSystemGemexitTxRequest.md)| exit Zar token from Vat contract transaction request | |

### Return type

[**ChainActivity**](../Models/ChainActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="exitZarTransaction"></a>
# **exitZarTransaction**
> ChainActivity exitZarTransaction(StablecoinSystemZarexitTxRequest)

exit ZAR token from Vat contract

    exit ZAR token from Vat contract

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **StablecoinSystemZarexitTxRequest** | [**StablecoinSystemZarexitTxRequest**](../Models/StablecoinSystemZarexitTxRequest.md)| exit Zar token from Vat contract transaction request | |

### Return type

[**ChainActivity**](../Models/ChainActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="getAllIlks"></a>
# **getAllIlks**
> IlksResponse getAllIlks()

Get all Ilks

    Retrieve a list of all available ilks.

### Parameters
This endpoint does not need any parameter.

### Return type

[**IlksResponse**](../Models/IlksResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getCollectorData"></a>
# **getCollectorData**
> Stats getCollectorData()

Get collector data

    Get collector data

### Parameters
This endpoint does not need any parameter.

### Return type

[**Stats**](../Models/Stats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getIlkByName"></a>
# **getIlkByName**
> Ilk getIlkByName(name)

Get Ilk by name

    Retrieve an Ilk by providing its name.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **name** | **String**| Name of the ILK | [default to null] |

### Return type

[**Ilk**](../Models/Ilk.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getVaultById"></a>
# **getVaultById**
> Vault getVaultById(id)

Get a vault by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | **Integer**| Vault ID | [default to null] |

### Return type

[**Vault**](../Models/Vault.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getVaultEventsById"></a>
# **getVaultEventsById**
> VaultEventsResponse getVaultEventsById(id, type)

Get vault events by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | **Integer**| Vault ID | [default to null] |
| **type** | **String**| Event type | [optional] [default to null] [enum: repay, deposit, withdraw, mint] |

### Return type

[**VaultEventsResponse**](../Models/VaultEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getVaultsByOwner"></a>
# **getVaultsByOwner**
> VaultsResponse getVaultsByOwner(owner)

Get vaults by owner query

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **owner** | **String**| Ethereum address of the owner | [optional] [default to null] |

### Return type

[**VaultsResponse**](../Models/VaultsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="liquidateVaultTransaction"></a>
# **liquidateVaultTransaction**
> ChainActivity liquidateVaultTransaction(StablecoinSystemBarkTxRequest)

liquidate a vault

    liquidate a vault

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **StablecoinSystemBarkTxRequest** | [**StablecoinSystemBarkTxRequest**](../Models/StablecoinSystemBarkTxRequest.md)| Liquidate a Vault transaction request | |

### Return type

[**ChainActivity**](../Models/ChainActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="mintZarTransaction"></a>
# **mintZarTransaction**
> ChainActivity mintZarTransaction(StablecoinSystemMintZarTxRequest)

Mint ZAR

    Mint ZAR

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **StablecoinSystemMintZarTxRequest** | [**StablecoinSystemMintZarTxRequest**](../Models/StablecoinSystemMintZarTxRequest.md)| Mint ZAR transaction request | |

### Return type

[**ChainActivity**](../Models/ChainActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="repayZarTransaction"></a>
# **repayZarTransaction**
> ChainActivity repayZarTransaction(StablecoinSystemRepayZarTxRequest)

Repay ZAR

    Repay ZAR

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **StablecoinSystemRepayZarTxRequest** | [**StablecoinSystemRepayZarTxRequest**](../Models/StablecoinSystemRepayZarTxRequest.md)| Repay ZAR transaction request | |

### Return type

[**ChainActivity**](../Models/ChainActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="resetAuctionTransaction"></a>
# **resetAuctionTransaction**
> ChainActivity resetAuctionTransaction(StablecoinSystemRedoTxRequest)

reset a auction

    reset a auction

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **StablecoinSystemRedoTxRequest** | [**StablecoinSystemRedoTxRequest**](../Models/StablecoinSystemRedoTxRequest.md)| Reset a Auction transaction request | |

### Return type

[**ChainActivity**](../Models/ChainActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="takeAuctionTransaction"></a>
# **takeAuctionTransaction**
> ChainActivity takeAuctionTransaction(StablecoinSystemTakeTxRequest)

take a auction

    take a auction

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **StablecoinSystemTakeTxRequest** | [**StablecoinSystemTakeTxRequest**](../Models/StablecoinSystemTakeTxRequest.md)| Take a Auction transaction request | |

### Return type

[**ChainActivity**](../Models/ChainActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="withdrawCollateralTransaction"></a>
# **withdrawCollateralTransaction**
> ChainActivity withdrawCollateralTransaction(StablecoinSystemWithdrawCollateralTxRequest)

Withdraw collateral

    Withdraw collateral

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **StablecoinSystemWithdrawCollateralTxRequest** | [**StablecoinSystemWithdrawCollateralTxRequest**](../Models/StablecoinSystemWithdrawCollateralTxRequest.md)| Withdraw collateral transaction request | |

### Return type

[**ChainActivity**](../Models/ChainActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


# VaultsApi

All URIs are relative to *https://api.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getVaultById**](VaultsApi.md#getVaultById) | **GET** /v2/vaults/{id} | Get a vault by ID |
| [**getVaultEventsById**](VaultsApi.md#getVaultEventsById) | **GET** /v2/vaults/{id}/events | Get vault events by ID |
| [**getVaultsByOwner**](VaultsApi.md#getVaultsByOwner) | **GET** /v2/vaults | Get vaults by owner query |


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


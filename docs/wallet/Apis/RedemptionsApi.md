# RedemptionsApi

All URIs are relative to *https://wapi.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getAllRedemptions**](RedemptionsApi.md#getAllRedemptions) | **GET** /admin/redemptions | Get all redemptions |
| [**getRedemptionDetails**](RedemptionsApi.md#getRedemptionDetails) | **GET** /admin/redemptions/{id} | Get redemption |
| [**redeemZar**](RedemptionsApi.md#redeemZar) | **POST** /redemptions | Redeem zar |
| [**updateRedemptionStatus**](RedemptionsApi.md#updateRedemptionStatus) | **POST** /admin/redemptions/{id} | Update redemption |


<a name="getAllRedemptions"></a>
# **getAllRedemptions**
> RedemptionResponse getAllRedemptions(state)

Get all redemptions

    Get a list of all redemptions.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **state** | **String**| Cursor for pagination | [optional] [default to null] [enum: pending, approved, completed, rejected] |

### Return type

[**RedemptionResponse**](../Models/RedemptionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getRedemptionDetails"></a>
# **getRedemptionDetails**
> Redemption getRedemptionDetails(id)

Get redemption

    Get the details of a redemption.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | **Long**| Redemption ID | [default to null] |

### Return type

[**Redemption**](../Models/Redemption.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="redeemZar"></a>
# **redeemZar**
> Redemption redeemZar(RedemptionRequest)

Redeem zar

    Redeem zar

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **RedemptionRequest** | [**RedemptionRequest**](../Models/RedemptionRequest.md)|  | |

### Return type

[**Redemption**](../Models/Redemption.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateRedemptionStatus"></a>
# **updateRedemptionStatus**
> Redemption updateRedemptionStatus(id, AdminRedemptionUpdateRequest)

Update redemption

    Update the status of a redemption.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | **String**| Redemption ID | [default to null] |
| **AdminRedemptionUpdateRequest** | [**AdminRedemptionUpdateRequest**](../Models/AdminRedemptionUpdateRequest.md)|  | |

### Return type

[**Redemption**](../Models/Redemption.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


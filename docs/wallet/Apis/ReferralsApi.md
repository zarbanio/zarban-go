# ReferralsApi

All URIs are relative to *https://wapi.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getReferralById**](ReferralsApi.md#getReferralById) | **GET** /referrals/{referralId} | Get referral by ID |
| [**getReferrals**](ReferralsApi.md#getReferrals) | **GET** /referrals | Get referrals |
| [**redeemReferral**](ReferralsApi.md#redeemReferral) | **POST** /referrals/{referralId}/redeem | Redeem a referral |
| [**validateReferral**](ReferralsApi.md#validateReferral) | **POST** /referrals/{referralId}/validate | Validate a referral |


<a name="getReferralById"></a>
# **getReferralById**
> Referral getReferralById(referralId)

Get referral by ID

    Get a referral by its ID.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **referralId** | **Long**| Referral ID | [default to null] |

### Return type

[**Referral**](../Models/Referral.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getReferrals"></a>
# **getReferrals**
> ReferralResponse getReferrals(name)

Get referrals

    Get referrals by user ID or name.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **name** | **String**| Referral name | [optional] [default to null] |

### Return type

[**ReferralResponse**](../Models/ReferralResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="redeemReferral"></a>
# **redeemReferral**
> redeemReferral(referralId)

Redeem a referral

    Redeem a referral for a user.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **referralId** | **Long**| Referral ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="validateReferral"></a>
# **validateReferral**
> validateReferral(referralId)

Validate a referral

    Validate a referral to check if it can be redeemed.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **referralId** | **Long**| Referral ID | [default to null] |

### Return type

null (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


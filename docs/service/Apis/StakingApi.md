# StakingApi

All URIs are relative to *https://api.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**collectStakingReward**](StakingApi.md#collectStakingReward) | **POST** /v2/staking/tx/collectreward | Collect staking reward |
| [**getStakingPlans**](StakingApi.md#getStakingPlans) | **GET** /v2/staking/plans | Get staking plans |
| [**getUserStakingStats**](StakingApi.md#getUserStakingStats) | **GET** /v2/staking/stats | Get user staking stats |
| [**stakeToStakingContract**](StakingApi.md#stakeToStakingContract) | **POST** /v2/staking/tx/stake | Stake to staking contract |
| [**withdrawStakedAsset**](StakingApi.md#withdrawStakedAsset) | **POST** /v2/staking/tx/withdraw | Withdraw staked asset |


<a name="collectStakingReward"></a>
# **collectStakingReward**
> StakingCollectRewardTxResponse collectStakingReward(StakingCollectRewardTxRequest)

Collect staking reward

    Collect staking reward

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **StakingCollectRewardTxRequest** | [**StakingCollectRewardTxRequest**](../Models/StakingCollectRewardTxRequest.md)| Collect reward transaction request | |

### Return type

[**StakingCollectRewardTxResponse**](../Models/StakingCollectRewardTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="getStakingPlans"></a>
# **getStakingPlans**
> StakePlansResponse getStakingPlans()

Get staking plans

    Get staking plans

### Parameters
This endpoint does not need any parameter.

### Return type

[**StakePlansResponse**](../Models/StakePlansResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserStakingStats"></a>
# **getUserStakingStats**
> UserStakesResponse getUserStakingStats(user, address, active, cursor, limit)

Get user staking stats

    Get user staking stats

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | **String**| Ethereum address of the user | [optional] [default to null] |
| **address** | **String**| Ethereum address of the staking contract | [optional] [default to null] |
| **active** | **Boolean**| Filter by active stakes | [optional] [default to null] |
| **cursor** | **Integer**| Cursor for pagination | [optional] [default to null] |
| **limit** | **Integer**| Limit the number of stakes returned (default is 50) | [optional] [default to 50] |

### Return type

[**UserStakesResponse**](../Models/UserStakesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="stakeToStakingContract"></a>
# **stakeToStakingContract**
> StakingStakeTxResponse stakeToStakingContract(StakingStakeTxRequest)

Stake to staking contract

    Stake to staking contract

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **StakingStakeTxRequest** | [**StakingStakeTxRequest**](../Models/StakingStakeTxRequest.md)| Stake transaction request, if amount is not provided, then the whole wallet balance will be used | |

### Return type

[**StakingStakeTxResponse**](../Models/StakingStakeTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="withdrawStakedAsset"></a>
# **withdrawStakedAsset**
> StakingWithdrawTxResponse withdrawStakedAsset(StakingWithdrawTxRequest)

Withdraw staked asset

    Withdraw staked asset

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **StakingWithdrawTxRequest** | [**StakingWithdrawTxRequest**](../Models/StakingWithdrawTxRequest.md)| Withdraw transaction request, if amount is not provided, then the whole staked amount will be withdrawn | |

### Return type

[**StakingWithdrawTxResponse**](../Models/StakingWithdrawTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


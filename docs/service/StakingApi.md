# zarban.service.openapi_client.StakingApi

All URIs are relative to *https://api.zarban.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**collect_staking_reward**](StakingApi.md#collect_staking_reward) | **POST** /v2/staking/tx/collectreward | Collect staking reward
[**get_staking_plans**](StakingApi.md#get_staking_plans) | **GET** /v2/staking/plans | Get staking plans
[**get_user_staking_stats**](StakingApi.md#get_user_staking_stats) | **GET** /v2/staking/stats | Get user staking stats
[**stake_to_staking_contract**](StakingApi.md#stake_to_staking_contract) | **POST** /v2/staking/tx/stake | Stake to staking contract
[**withdraw_staked_asset**](StakingApi.md#withdraw_staked_asset) | **POST** /v2/staking/tx/withdraw | Withdraw staked asset


# **collect_staking_reward**
> StakingCollectRewardTxResponse collect_staking_reward(staking_collect_reward_tx_request)

Collect staking reward

Collect staking reward

### Example

```python
from __future__ import print_function
import time
import zarban.service.openapi_client
from zarban.service.openapi_client.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.zarban.io
# See configuration.py for a list of all supported configuration parameters.
configuration = zarban.service.openapi_client.Configuration(
    host = "https://api.zarban.io"
)


# Enter a context with an instance of the API client
with zarban.service.openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = zarban.service.openapi_client.StakingApi(api_client)
    staking_collect_reward_tx_request = zarban.service.openapi_client.StakingCollectRewardTxRequest() # StakingCollectRewardTxRequest | Collect reward transaction request

    try:
        # Collect staking reward
        api_response = api_instance.collect_staking_reward(staking_collect_reward_tx_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StakingApi->collect_staking_reward: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **staking_collect_reward_tx_request** | [**StakingCollectRewardTxRequest**](StakingCollectRewardTxRequest.md)| Collect reward transaction request | 

### Return type

[**StakingCollectRewardTxResponse**](StakingCollectRewardTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Bad request |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_staking_plans**
> StakePlansResponse get_staking_plans()

Get staking plans

Get staking plans

### Example

```python
from __future__ import print_function
import time
import zarban.service.openapi_client
from zarban.service.openapi_client.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.zarban.io
# See configuration.py for a list of all supported configuration parameters.
configuration = zarban.service.openapi_client.Configuration(
    host = "https://api.zarban.io"
)


# Enter a context with an instance of the API client
with zarban.service.openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = zarban.service.openapi_client.StakingApi(api_client)
    
    try:
        # Get staking plans
        api_response = api_instance.get_staking_plans()
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StakingApi->get_staking_plans: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**StakePlansResponse**](StakePlansResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of stakes |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_user_staking_stats**
> UserStakesResponse get_user_staking_stats(user=user, address=address, active=active, cursor=cursor, limit=limit)

Get user staking stats

Get user staking stats

### Example

```python
from __future__ import print_function
import time
import zarban.service.openapi_client
from zarban.service.openapi_client.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.zarban.io
# See configuration.py for a list of all supported configuration parameters.
configuration = zarban.service.openapi_client.Configuration(
    host = "https://api.zarban.io"
)


# Enter a context with an instance of the API client
with zarban.service.openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = zarban.service.openapi_client.StakingApi(api_client)
    user = 'user_example' # str | Ethereum address of the user (optional)
address = 'address_example' # str | Ethereum address of the staking contract (optional)
active = True # bool | Filter by active stakes (optional)
cursor = 56 # int | Cursor for pagination (optional)
limit = 50 # int | Limit the number of stakes returned (default is 50) (optional) (default to 50)

    try:
        # Get user staking stats
        api_response = api_instance.get_user_staking_stats(user=user, address=address, active=active, cursor=cursor, limit=limit)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StakingApi->get_user_staking_stats: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **str**| Ethereum address of the user | [optional] 
 **address** | **str**| Ethereum address of the staking contract | [optional] 
 **active** | **bool**| Filter by active stakes | [optional] 
 **cursor** | **int**| Cursor for pagination | [optional] 
 **limit** | **int**| Limit the number of stakes returned (default is 50) | [optional] [default to 50]

### Return type

[**UserStakesResponse**](UserStakesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of stakes |  -  |
**400** | Bad request |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **stake_to_staking_contract**
> StakingStakeTxResponse stake_to_staking_contract(staking_stake_tx_request)

Stake to staking contract

Stake to staking contract

### Example

```python
from __future__ import print_function
import time
import zarban.service.openapi_client
from zarban.service.openapi_client.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.zarban.io
# See configuration.py for a list of all supported configuration parameters.
configuration = zarban.service.openapi_client.Configuration(
    host = "https://api.zarban.io"
)


# Enter a context with an instance of the API client
with zarban.service.openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = zarban.service.openapi_client.StakingApi(api_client)
    staking_stake_tx_request = zarban.service.openapi_client.StakingStakeTxRequest() # StakingStakeTxRequest | Stake transaction request, if amount is not provided, then the whole wallet balance will be used

    try:
        # Stake to staking contract
        api_response = api_instance.stake_to_staking_contract(staking_stake_tx_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StakingApi->stake_to_staking_contract: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **staking_stake_tx_request** | [**StakingStakeTxRequest**](StakingStakeTxRequest.md)| Stake transaction request, if amount is not provided, then the whole wallet balance will be used | 

### Return type

[**StakingStakeTxResponse**](StakingStakeTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Bad request |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **withdraw_staked_asset**
> StakingWithdrawTxResponse withdraw_staked_asset(staking_withdraw_tx_request)

Withdraw staked asset

Withdraw staked asset

### Example

```python
from __future__ import print_function
import time
import zarban.service.openapi_client
from zarban.service.openapi_client.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.zarban.io
# See configuration.py for a list of all supported configuration parameters.
configuration = zarban.service.openapi_client.Configuration(
    host = "https://api.zarban.io"
)


# Enter a context with an instance of the API client
with zarban.service.openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = zarban.service.openapi_client.StakingApi(api_client)
    staking_withdraw_tx_request = zarban.service.openapi_client.StakingWithdrawTxRequest() # StakingWithdrawTxRequest | Withdraw transaction request, if amount is not provided, then the whole staked amount will be withdrawn

    try:
        # Withdraw staked asset
        api_response = api_instance.withdraw_staked_asset(staking_withdraw_tx_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StakingApi->withdraw_staked_asset: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **staking_withdraw_tx_request** | [**StakingWithdrawTxRequest**](StakingWithdrawTxRequest.md)| Withdraw transaction request, if amount is not provided, then the whole staked amount will be withdrawn | 

### Return type

[**StakingWithdrawTxResponse**](StakingWithdrawTxResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Bad request |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


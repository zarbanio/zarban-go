# zarban.service.openapi_client.LendingPoolApi

All URIs are relative to *https://api.zarban.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_lending_pool_borrow**](LendingPoolApi.md#create_lending_pool_borrow) | **POST** /v2/lendingpool/tx/borrow | Borrow from lending pool
[**create_lending_pool_deposit**](LendingPoolApi.md#create_lending_pool_deposit) | **POST** /v2/lendingpool/tx/deposit | Deposit to lending pool
[**create_lending_pool_repay**](LendingPoolApi.md#create_lending_pool_repay) | **POST** /v2/lendingpool/tx/repay | Repay to lending pool
[**create_lending_pool_withdraw**](LendingPoolApi.md#create_lending_pool_withdraw) | **POST** /v2/lendingpool/tx/withdraw | Withdraw from lending pool
[**fetch_reserve_data_by_asset**](LendingPoolApi.md#fetch_reserve_data_by_asset) | **GET** /v2/lendingpool/reserves | Fetch Reserve Data By Asset
[**get_user_borrows**](LendingPoolApi.md#get_user_borrows) | **GET** /v2/lendingpool/borrows | Get user borrows of lendingpool
[**get_user_deposits**](LendingPoolApi.md#get_user_deposits) | **GET** /v2/lendingpool/deposits | Get user deposits of Lendingpool
[**set_lending_pool_asset_collateral**](LendingPoolApi.md#set_lending_pool_asset_collateral) | **POST** /v2/lendingpool/tx/useassetascollateral | Enable/Disable asset as collateral


# **create_lending_pool_borrow**
> LendingpoolBorrowTxResponse create_lending_pool_borrow(lendingpool_borrow_tx_request)

Borrow from lending pool

Borrow from lending pool

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
    api_instance = zarban.service.openapi_client.LendingPoolApi(api_client)
    lendingpool_borrow_tx_request = zarban.service.openapi_client.LendingpoolBorrowTxRequest() # LendingpoolBorrowTxRequest | Borrow transaction request, if amount is not provided, it will be calculated based on the token balance

    try:
        # Borrow from lending pool
        api_response = api_instance.create_lending_pool_borrow(lendingpool_borrow_tx_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling LendingPoolApi->create_lending_pool_borrow: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lendingpool_borrow_tx_request** | [**LendingpoolBorrowTxRequest**](LendingpoolBorrowTxRequest.md)| Borrow transaction request, if amount is not provided, it will be calculated based on the token balance | 

### Return type

[**LendingpoolBorrowTxResponse**](LendingpoolBorrowTxResponse.md)

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

# **create_lending_pool_deposit**
> LendingpoolDepositTxResponse create_lending_pool_deposit(lendingpool_deposit_tx_request)

Deposit to lending pool

Deposit to lending pool

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
    api_instance = zarban.service.openapi_client.LendingPoolApi(api_client)
    lendingpool_deposit_tx_request = zarban.service.openapi_client.LendingpoolDepositTxRequest() # LendingpoolDepositTxRequest | Deposit transaction request, if amount is not provided, it will be calculated based on the token balance

    try:
        # Deposit to lending pool
        api_response = api_instance.create_lending_pool_deposit(lendingpool_deposit_tx_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling LendingPoolApi->create_lending_pool_deposit: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lendingpool_deposit_tx_request** | [**LendingpoolDepositTxRequest**](LendingpoolDepositTxRequest.md)| Deposit transaction request, if amount is not provided, it will be calculated based on the token balance | 

### Return type

[**LendingpoolDepositTxResponse**](LendingpoolDepositTxResponse.md)

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

# **create_lending_pool_repay**
> LendingpoolRepayTxResponse create_lending_pool_repay(lendingpool_repay_tx_request)

Repay to lending pool

Repay to lending pool

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
    api_instance = zarban.service.openapi_client.LendingPoolApi(api_client)
    lendingpool_repay_tx_request = zarban.service.openapi_client.LendingpoolRepayTxRequest() # LendingpoolRepayTxRequest | Repay transaction request, if amount is not provided, it will repay the maximum possible amount

    try:
        # Repay to lending pool
        api_response = api_instance.create_lending_pool_repay(lendingpool_repay_tx_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling LendingPoolApi->create_lending_pool_repay: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lendingpool_repay_tx_request** | [**LendingpoolRepayTxRequest**](LendingpoolRepayTxRequest.md)| Repay transaction request, if amount is not provided, it will repay the maximum possible amount | 

### Return type

[**LendingpoolRepayTxResponse**](LendingpoolRepayTxResponse.md)

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

# **create_lending_pool_withdraw**
> LendingpoolWithdrawTxResponse create_lending_pool_withdraw(lendingpool_withdraw_tx_request)

Withdraw from lending pool

Withdraw from lending pool

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
    api_instance = zarban.service.openapi_client.LendingPoolApi(api_client)
    lendingpool_withdraw_tx_request = zarban.service.openapi_client.LendingpoolWithdrawTxRequest() # LendingpoolWithdrawTxRequest | Withdraw transaction request, if amount is not provided, it will be calculated based on the user account status

    try:
        # Withdraw from lending pool
        api_response = api_instance.create_lending_pool_withdraw(lendingpool_withdraw_tx_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling LendingPoolApi->create_lending_pool_withdraw: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lendingpool_withdraw_tx_request** | [**LendingpoolWithdrawTxRequest**](LendingpoolWithdrawTxRequest.md)| Withdraw transaction request, if amount is not provided, it will be calculated based on the user account status | 

### Return type

[**LendingpoolWithdrawTxResponse**](LendingpoolWithdrawTxResponse.md)

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

# **fetch_reserve_data_by_asset**
> FormattedReserveData fetch_reserve_data_by_asset(asset=asset)

Fetch Reserve Data By Asset

Retrieve data of reserves

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
    api_instance = zarban.service.openapi_client.LendingPoolApi(api_client)
    asset = 'asset_example' # str | The asset address in hexadecimal format. (optional)

    try:
        # Fetch Reserve Data By Asset
        api_response = api_instance.fetch_reserve_data_by_asset(asset=asset)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling LendingPoolApi->fetch_reserve_data_by_asset: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **str**| The asset address in hexadecimal format. | [optional] 

### Return type

[**FormattedReserveData**](FormattedReserveData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response. |  -  |
**400** | Bad request (e.g. invalid asset address). |  -  |
**500** | Internal Server Error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_user_borrows**
> UserBorrowsResponse get_user_borrows(user=user, reserve=reserve, cursor=cursor, limit=limit)

Get user borrows of lendingpool

Get user borrows of lendingpool

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
    api_instance = zarban.service.openapi_client.LendingPoolApi(api_client)
    user = 'user_example' # str | Ethereum address of the user (optional)
reserve = 'reserve_example' # str | Ethereum address of the reserve (optional)
cursor = 56 # int | Cursor for pagination (optional)
limit = 50 # int | Limit the number of deposits returned (default is 50) (optional) (default to 50)

    try:
        # Get user borrows of lendingpool
        api_response = api_instance.get_user_borrows(user=user, reserve=reserve, cursor=cursor, limit=limit)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling LendingPoolApi->get_user_borrows: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **str**| Ethereum address of the user | [optional] 
 **reserve** | **str**| Ethereum address of the reserve | [optional] 
 **cursor** | **int**| Cursor for pagination | [optional] 
 **limit** | **int**| Limit the number of deposits returned (default is 50) | [optional] [default to 50]

### Return type

[**UserBorrowsResponse**](UserBorrowsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of user borrows |  -  |
**400** | Bad request |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_user_deposits**
> UserDepositsResponse get_user_deposits(user=user, reserve=reserve, cursor=cursor, limit=limit)

Get user deposits of Lendingpool

Get user deposits of Lendingpool

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
    api_instance = zarban.service.openapi_client.LendingPoolApi(api_client)
    user = 'user_example' # str | Ethereum address of the user (optional)
reserve = 'reserve_example' # str | Ethereum address of the reserve (optional)
cursor = 56 # int | Cursor for pagination (optional)
limit = 50 # int | Limit the number of deposits returned (default is 50) (optional) (default to 50)

    try:
        # Get user deposits of Lendingpool
        api_response = api_instance.get_user_deposits(user=user, reserve=reserve, cursor=cursor, limit=limit)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling LendingPoolApi->get_user_deposits: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **str**| Ethereum address of the user | [optional] 
 **reserve** | **str**| Ethereum address of the reserve | [optional] 
 **cursor** | **int**| Cursor for pagination | [optional] 
 **limit** | **int**| Limit the number of deposits returned (default is 50) | [optional] [default to 50]

### Return type

[**UserDepositsResponse**](UserDepositsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of user deposits |  -  |
**400** | Bad request |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **set_lending_pool_asset_collateral**
> LendingpoolUseAssetAsCollateralTxResponse set_lending_pool_asset_collateral(lendingpool_use_asset_as_collateral_tx_request)

Enable/Disable asset as collateral

Allows to enable/disable a specific asset as collateral

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
    api_instance = zarban.service.openapi_client.LendingPoolApi(api_client)
    lendingpool_use_asset_as_collateral_tx_request = zarban.service.openapi_client.LendingpoolUseAssetAsCollateralTxRequest() # LendingpoolUseAssetAsCollateralTxRequest | UseAssetAsCollateral transaction request

    try:
        # Enable/Disable asset as collateral
        api_response = api_instance.set_lending_pool_asset_collateral(lendingpool_use_asset_as_collateral_tx_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling LendingPoolApi->set_lending_pool_asset_collateral: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lendingpool_use_asset_as_collateral_tx_request** | [**LendingpoolUseAssetAsCollateralTxRequest**](LendingpoolUseAssetAsCollateralTxRequest.md)| UseAssetAsCollateral transaction request | 

### Return type

[**LendingpoolUseAssetAsCollateralTxResponse**](LendingpoolUseAssetAsCollateralTxResponse.md)

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


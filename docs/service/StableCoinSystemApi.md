# zarban.service.openapi_client.StableCoinSystemApi

All URIs are relative to *https://api.zarban.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**approve_and_join_zar_transaction**](StableCoinSystemApi.md#approve_and_join_zar_transaction) | **POST** /v2/stablecoinsystem/auctions/tx/zarjoin | approve and join ZAR token into Vat contract
[**create_stable_coin_vault**](StableCoinSystemApi.md#create_stable_coin_vault) | **POST** /v2/stablecoinsystem/tx/createvault | Create vault
[**deposit_stable_coin_collateral**](StableCoinSystemApi.md#deposit_stable_coin_collateral) | **POST** /v2/stablecoinsystem/tx/depositcollateral | Deposit collateral
[**exit_gem_transaction**](StableCoinSystemApi.md#exit_gem_transaction) | **POST** /v2/stablecoinsystem/auctions/tx/gemexit | exit Gem token (which can be used as collateral) from Vat contract
[**exit_zar_transaction**](StableCoinSystemApi.md#exit_zar_transaction) | **POST** /v2/stablecoinsystem/auctions/tx/zarexit | exit ZAR token from Vat contract
[**get_all_ilks**](StableCoinSystemApi.md#get_all_ilks) | **GET** /v2/ilks | Get all Ilks
[**get_collector_data**](StableCoinSystemApi.md#get_collector_data) | **GET** /v2/stats | Get collector data
[**get_ilk_by_name**](StableCoinSystemApi.md#get_ilk_by_name) | **GET** /v2/ilks/{name} | Get Ilk by name
[**get_vault_by_id**](StableCoinSystemApi.md#get_vault_by_id) | **GET** /v2/vaults/{id} | Get a vault by ID
[**get_vault_events_by_id**](StableCoinSystemApi.md#get_vault_events_by_id) | **GET** /v2/vaults/{id}/events | Get vault events by ID
[**get_vaults_by_owner**](StableCoinSystemApi.md#get_vaults_by_owner) | **GET** /v2/vaults | Get vaults by owner query
[**liquidate_vault_transaction**](StableCoinSystemApi.md#liquidate_vault_transaction) | **POST** /v2/stablecoinsystem/tx/bark | liquidate a vault
[**mint_zar_transaction**](StableCoinSystemApi.md#mint_zar_transaction) | **POST** /v2/stablecoinsystem/tx/mintzar | Mint ZAR
[**repay_zar_transaction**](StableCoinSystemApi.md#repay_zar_transaction) | **POST** /v2/stablecoinsystem/tx/repayzar | Repay ZAR
[**reset_auction_transaction**](StableCoinSystemApi.md#reset_auction_transaction) | **POST** /v2/stablecoinsystem/auctions/tx/redo | reset a auction
[**take_auction_transaction**](StableCoinSystemApi.md#take_auction_transaction) | **POST** /v2/stablecoinsystem/auctions/tx/take | take a auction
[**withdraw_collateral_transaction**](StableCoinSystemApi.md#withdraw_collateral_transaction) | **POST** /v2/stablecoinsystem/tx/withdrawcollateral | Withdraw collateral


# **approve_and_join_zar_transaction**
> ChainActivity approve_and_join_zar_transaction(stablecoin_system_zarjoin_tx_request)

approve and join ZAR token into Vat contract

approve and join ZAR token into Vat contract

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
    api_instance = zarban.service.openapi_client.StableCoinSystemApi(api_client)
    stablecoin_system_zarjoin_tx_request = zarban.service.openapi_client.StablecoinSystemZarjoinTxRequest() # StablecoinSystemZarjoinTxRequest | Approve and join Zar token into Vat contract transaction request

    try:
        # approve and join ZAR token into Vat contract
        api_response = api_instance.approve_and_join_zar_transaction(stablecoin_system_zarjoin_tx_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StableCoinSystemApi->approve_and_join_zar_transaction: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stablecoin_system_zarjoin_tx_request** | [**StablecoinSystemZarjoinTxRequest**](StablecoinSystemZarjoinTxRequest.md)| Approve and join Zar token into Vat contract transaction request | 

### Return type

[**ChainActivity**](ChainActivity.md)

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

# **create_stable_coin_vault**
> ChainActivity create_stable_coin_vault(stablecoin_system_create_vault_tx_request)

Create vault

Create vault

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
    api_instance = zarban.service.openapi_client.StableCoinSystemApi(api_client)
    stablecoin_system_create_vault_tx_request = zarban.service.openapi_client.StablecoinSystemCreateVaultTxRequest() # StablecoinSystemCreateVaultTxRequest | Create vault transaction request

    try:
        # Create vault
        api_response = api_instance.create_stable_coin_vault(stablecoin_system_create_vault_tx_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StableCoinSystemApi->create_stable_coin_vault: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stablecoin_system_create_vault_tx_request** | [**StablecoinSystemCreateVaultTxRequest**](StablecoinSystemCreateVaultTxRequest.md)| Create vault transaction request | 

### Return type

[**ChainActivity**](ChainActivity.md)

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

# **deposit_stable_coin_collateral**
> ChainActivity deposit_stable_coin_collateral(stablecoin_system_deposit_collateral_tx_request)

Deposit collateral

Deposit collateral

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
    api_instance = zarban.service.openapi_client.StableCoinSystemApi(api_client)
    stablecoin_system_deposit_collateral_tx_request = zarban.service.openapi_client.StablecoinSystemDepositCollateralTxRequest() # StablecoinSystemDepositCollateralTxRequest | Deposit collateral transaction request

    try:
        # Deposit collateral
        api_response = api_instance.deposit_stable_coin_collateral(stablecoin_system_deposit_collateral_tx_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StableCoinSystemApi->deposit_stable_coin_collateral: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stablecoin_system_deposit_collateral_tx_request** | [**StablecoinSystemDepositCollateralTxRequest**](StablecoinSystemDepositCollateralTxRequest.md)| Deposit collateral transaction request | 

### Return type

[**ChainActivity**](ChainActivity.md)

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

# **exit_gem_transaction**
> ChainActivity exit_gem_transaction(stablecoin_system_gemexit_tx_request)

exit Gem token (which can be used as collateral) from Vat contract

exit Gem token (which can be used as collateral) from Vat contract

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
    api_instance = zarban.service.openapi_client.StableCoinSystemApi(api_client)
    stablecoin_system_gemexit_tx_request = zarban.service.openapi_client.StablecoinSystemGemexitTxRequest() # StablecoinSystemGemexitTxRequest | exit Zar token from Vat contract transaction request

    try:
        # exit Gem token (which can be used as collateral) from Vat contract
        api_response = api_instance.exit_gem_transaction(stablecoin_system_gemexit_tx_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StableCoinSystemApi->exit_gem_transaction: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stablecoin_system_gemexit_tx_request** | [**StablecoinSystemGemexitTxRequest**](StablecoinSystemGemexitTxRequest.md)| exit Zar token from Vat contract transaction request | 

### Return type

[**ChainActivity**](ChainActivity.md)

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

# **exit_zar_transaction**
> ChainActivity exit_zar_transaction(stablecoin_system_zarexit_tx_request)

exit ZAR token from Vat contract

exit ZAR token from Vat contract

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
    api_instance = zarban.service.openapi_client.StableCoinSystemApi(api_client)
    stablecoin_system_zarexit_tx_request = zarban.service.openapi_client.StablecoinSystemZarexitTxRequest() # StablecoinSystemZarexitTxRequest | exit Zar token from Vat contract transaction request

    try:
        # exit ZAR token from Vat contract
        api_response = api_instance.exit_zar_transaction(stablecoin_system_zarexit_tx_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StableCoinSystemApi->exit_zar_transaction: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stablecoin_system_zarexit_tx_request** | [**StablecoinSystemZarexitTxRequest**](StablecoinSystemZarexitTxRequest.md)| exit Zar token from Vat contract transaction request | 

### Return type

[**ChainActivity**](ChainActivity.md)

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

# **get_all_ilks**
> IlksResponse get_all_ilks()

Get all Ilks

Retrieve a list of all available ilks.

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
    api_instance = zarban.service.openapi_client.StableCoinSystemApi(api_client)
    
    try:
        # Get all Ilks
        api_response = api_instance.get_all_ilks()
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StableCoinSystemApi->get_all_ilks: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**IlksResponse**](IlksResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successfully retrieved list of ilks |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_collector_data**
> Stats get_collector_data()

Get collector data

Get collector data

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
    api_instance = zarban.service.openapi_client.StableCoinSystemApi(api_client)
    
    try:
        # Get collector data
        api_response = api_instance.get_collector_data()
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StableCoinSystemApi->get_collector_data: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**Stats**](Stats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A JSON object with collector data |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_ilk_by_name**
> Ilk get_ilk_by_name(name)

Get Ilk by name

Retrieve an Ilk by providing its name.

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
    api_instance = zarban.service.openapi_client.StableCoinSystemApi(api_client)
    name = 'name_example' # str | Name of the ILK

    try:
        # Get Ilk by name
        api_response = api_instance.get_ilk_by_name(name)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StableCoinSystemApi->get_ilk_by_name: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Name of the ILK | 

### Return type

[**Ilk**](Ilk.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successfully retrieved ILK |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_vault_by_id**
> Vault get_vault_by_id(id)

Get a vault by ID

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
    api_instance = zarban.service.openapi_client.StableCoinSystemApi(api_client)
    id = 56 # int | Vault ID

    try:
        # Get a vault by ID
        api_response = api_instance.get_vault_by_id(id)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StableCoinSystemApi->get_vault_by_id: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Vault ID | 

### Return type

[**Vault**](Vault.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Bad request |  -  |
**404** | Not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_vault_events_by_id**
> VaultEventsResponse get_vault_events_by_id(id, type=type)

Get vault events by ID

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
    api_instance = zarban.service.openapi_client.StableCoinSystemApi(api_client)
    id = 56 # int | Vault ID
type = 'type_example' # str | Event type (optional)

    try:
        # Get vault events by ID
        api_response = api_instance.get_vault_events_by_id(id, type=type)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StableCoinSystemApi->get_vault_events_by_id: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Vault ID | 
 **type** | **str**| Event type | [optional] 

### Return type

[**VaultEventsResponse**](VaultEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Bad request |  -  |
**404** | Not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_vaults_by_owner**
> VaultsResponse get_vaults_by_owner(owner=owner)

Get vaults by owner query

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
    api_instance = zarban.service.openapi_client.StableCoinSystemApi(api_client)
    owner = 'owner_example' # str | Ethereum address of the owner (optional)

    try:
        # Get vaults by owner query
        api_response = api_instance.get_vaults_by_owner(owner=owner)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StableCoinSystemApi->get_vaults_by_owner: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **str**| Ethereum address of the owner | [optional] 

### Return type

[**VaultsResponse**](VaultsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Bad request |  -  |
**404** | Not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **liquidate_vault_transaction**
> ChainActivity liquidate_vault_transaction(stablecoin_system_bark_tx_request)

liquidate a vault

liquidate a vault

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
    api_instance = zarban.service.openapi_client.StableCoinSystemApi(api_client)
    stablecoin_system_bark_tx_request = zarban.service.openapi_client.StablecoinSystemBarkTxRequest() # StablecoinSystemBarkTxRequest | Liquidate a Vault transaction request

    try:
        # liquidate a vault
        api_response = api_instance.liquidate_vault_transaction(stablecoin_system_bark_tx_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StableCoinSystemApi->liquidate_vault_transaction: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stablecoin_system_bark_tx_request** | [**StablecoinSystemBarkTxRequest**](StablecoinSystemBarkTxRequest.md)| Liquidate a Vault transaction request | 

### Return type

[**ChainActivity**](ChainActivity.md)

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

# **mint_zar_transaction**
> ChainActivity mint_zar_transaction(stablecoin_system_mint_zar_tx_request)

Mint ZAR

Mint ZAR

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
    api_instance = zarban.service.openapi_client.StableCoinSystemApi(api_client)
    stablecoin_system_mint_zar_tx_request = zarban.service.openapi_client.StablecoinSystemMintZarTxRequest() # StablecoinSystemMintZarTxRequest | Mint ZAR transaction request

    try:
        # Mint ZAR
        api_response = api_instance.mint_zar_transaction(stablecoin_system_mint_zar_tx_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StableCoinSystemApi->mint_zar_transaction: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stablecoin_system_mint_zar_tx_request** | [**StablecoinSystemMintZarTxRequest**](StablecoinSystemMintZarTxRequest.md)| Mint ZAR transaction request | 

### Return type

[**ChainActivity**](ChainActivity.md)

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

# **repay_zar_transaction**
> ChainActivity repay_zar_transaction(stablecoin_system_repay_zar_tx_request)

Repay ZAR

Repay ZAR

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
    api_instance = zarban.service.openapi_client.StableCoinSystemApi(api_client)
    stablecoin_system_repay_zar_tx_request = zarban.service.openapi_client.StablecoinSystemRepayZarTxRequest() # StablecoinSystemRepayZarTxRequest | Repay ZAR transaction request

    try:
        # Repay ZAR
        api_response = api_instance.repay_zar_transaction(stablecoin_system_repay_zar_tx_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StableCoinSystemApi->repay_zar_transaction: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stablecoin_system_repay_zar_tx_request** | [**StablecoinSystemRepayZarTxRequest**](StablecoinSystemRepayZarTxRequest.md)| Repay ZAR transaction request | 

### Return type

[**ChainActivity**](ChainActivity.md)

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

# **reset_auction_transaction**
> ChainActivity reset_auction_transaction(stablecoin_system_redo_tx_request)

reset a auction

reset a auction

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
    api_instance = zarban.service.openapi_client.StableCoinSystemApi(api_client)
    stablecoin_system_redo_tx_request = zarban.service.openapi_client.StablecoinSystemRedoTxRequest() # StablecoinSystemRedoTxRequest | Reset a Auction transaction request

    try:
        # reset a auction
        api_response = api_instance.reset_auction_transaction(stablecoin_system_redo_tx_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StableCoinSystemApi->reset_auction_transaction: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stablecoin_system_redo_tx_request** | [**StablecoinSystemRedoTxRequest**](StablecoinSystemRedoTxRequest.md)| Reset a Auction transaction request | 

### Return type

[**ChainActivity**](ChainActivity.md)

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

# **take_auction_transaction**
> ChainActivity take_auction_transaction(stablecoin_system_take_tx_request)

take a auction

take a auction

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
    api_instance = zarban.service.openapi_client.StableCoinSystemApi(api_client)
    stablecoin_system_take_tx_request = zarban.service.openapi_client.StablecoinSystemTakeTxRequest() # StablecoinSystemTakeTxRequest | Take a Auction transaction request

    try:
        # take a auction
        api_response = api_instance.take_auction_transaction(stablecoin_system_take_tx_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StableCoinSystemApi->take_auction_transaction: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stablecoin_system_take_tx_request** | [**StablecoinSystemTakeTxRequest**](StablecoinSystemTakeTxRequest.md)| Take a Auction transaction request | 

### Return type

[**ChainActivity**](ChainActivity.md)

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

# **withdraw_collateral_transaction**
> ChainActivity withdraw_collateral_transaction(stablecoin_system_withdraw_collateral_tx_request)

Withdraw collateral

Withdraw collateral

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
    api_instance = zarban.service.openapi_client.StableCoinSystemApi(api_client)
    stablecoin_system_withdraw_collateral_tx_request = zarban.service.openapi_client.StablecoinSystemWithdrawCollateralTxRequest() # StablecoinSystemWithdrawCollateralTxRequest | Withdraw collateral transaction request

    try:
        # Withdraw collateral
        api_response = api_instance.withdraw_collateral_transaction(stablecoin_system_withdraw_collateral_tx_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling StableCoinSystemApi->withdraw_collateral_transaction: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stablecoin_system_withdraw_collateral_tx_request** | [**StablecoinSystemWithdrawCollateralTxRequest**](StablecoinSystemWithdrawCollateralTxRequest.md)| Withdraw collateral transaction request | 

### Return type

[**ChainActivity**](ChainActivity.md)

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


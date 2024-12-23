# zarban.service.openapi_client.VaultsApi

All URIs are relative to *https://api.zarban.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_vault_by_id**](VaultsApi.md#get_vault_by_id) | **GET** /v2/vaults/{id} | Get a vault by ID
[**get_vault_events_by_id**](VaultsApi.md#get_vault_events_by_id) | **GET** /v2/vaults/{id}/events | Get vault events by ID
[**get_vaults_by_owner**](VaultsApi.md#get_vaults_by_owner) | **GET** /v2/vaults | Get vaults by owner query


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
    api_instance = zarban.service.openapi_client.VaultsApi(api_client)
    id = 56 # int | Vault ID

    try:
        # Get a vault by ID
        api_response = api_instance.get_vault_by_id(id)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling VaultsApi->get_vault_by_id: %s\n" % e)
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
    api_instance = zarban.service.openapi_client.VaultsApi(api_client)
    id = 56 # int | Vault ID
type = 'type_example' # str | Event type (optional)

    try:
        # Get vault events by ID
        api_response = api_instance.get_vault_events_by_id(id, type=type)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling VaultsApi->get_vault_events_by_id: %s\n" % e)
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
    api_instance = zarban.service.openapi_client.VaultsApi(api_client)
    owner = 'owner_example' # str | Ethereum address of the owner (optional)

    try:
        # Get vaults by owner query
        api_response = api_instance.get_vaults_by_owner(owner=owner)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling VaultsApi->get_vaults_by_owner: %s\n" % e)
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


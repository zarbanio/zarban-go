# zarban.service.openapi_client.LogsApi

All URIs are relative to *https://api.zarban.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_logs_by_transaction_hash**](LogsApi.md#get_logs_by_transaction_hash) | **GET** /v2/logs/{txHash} | Get raw and decoded logs by transaction hash


# **get_logs_by_transaction_hash**
> EventDetailsResponse get_logs_by_transaction_hash(tx_hash)

Get raw and decoded logs by transaction hash

Get raw and decoded logs by transaction hash

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
    api_instance = zarban.service.openapi_client.LogsApi(api_client)
    tx_hash = 'tx_hash_example' # str | Transaction hash

    try:
        # Get raw and decoded logs by transaction hash
        api_response = api_instance.get_logs_by_transaction_hash(tx_hash)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling LogsApi->get_logs_by_transaction_hash: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tx_hash** | **str**| Transaction hash | 

### Return type

[**EventDetailsResponse**](EventDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Event details |  -  |
**400** | Bad request |  -  |
**404** | Not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


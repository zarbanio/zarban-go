# zarban.service.openapi_client.AddressesApi

All URIs are relative to *https://api.zarban.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_all_addresses**](AddressesApi.md#get_all_addresses) | **GET** /v2/addresses | Get all addresses


# **get_all_addresses**
> AddressResponse get_all_addresses(format=format)

Get all addresses

All addresses knows to Zarban.

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
    api_instance = zarban.service.openapi_client.AddressesApi(api_client)
    format = 'format_example' # str | The type of addresses to return (optional)

    try:
        # Get all addresses
        api_response = api_instance.get_all_addresses(format=format)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling AddressesApi->get_all_addresses: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **str**| The type of addresses to return | [optional] 

### Return type

[**AddressResponse**](AddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A JSON object with addresses |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


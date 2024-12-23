# zarban.wallet.openapi_client.SwapApi

All URIs are relative to *https://wapi.zarban.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**swap_coins**](SwapApi.md#swap_coins) | **POST** /swap | Swap coins


# **swap_coins**
> SwapResponse swap_coins(swap_request)

Swap coins

Swap coins in the wallet.

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
from __future__ import print_function
import time
import zarban.wallet.openapi_client
from zarban.wallet.openapi_client.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://wapi.zarban.io
# See configuration.py for a list of all supported configuration parameters.
configuration = zarban.wallet.openapi_client.Configuration(
    host = "https://wapi.zarban.io"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = zarban.wallet.openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with zarban.wallet.openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = zarban.wallet.openapi_client.SwapApi(api_client)
    swap_request = zarban.wallet.openapi_client.SwapRequest() # SwapRequest | 

    try:
        # Swap coins
        api_response = api_instance.swap_coins(swap_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling SwapApi->swap_coins: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **swap_request** | [**SwapRequest**](SwapRequest.md)|  | 

### Return type

[**SwapResponse**](SwapResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**401** | Unauthorized |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


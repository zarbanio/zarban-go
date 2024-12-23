# zarban.wallet.openapi_client.HealthApi

All URIs are relative to *https://wapi.zarban.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**check_api_health**](HealthApi.md#check_api_health) | **GET** /healthz | Health check


# **check_api_health**
> HealthStatus check_api_health()

Health check

Check the health of the API.

### Example

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


# Enter a context with an instance of the API client
with zarban.wallet.openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = zarban.wallet.openapi_client.HealthApi(api_client)
    
    try:
        # Health check
        api_response = api_instance.check_api_health()
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling HealthApi->check_api_health: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**HealthStatus**](HealthStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


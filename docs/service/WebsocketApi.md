# zarban.service.openapi_client.WebsocketApi

All URIs are relative to *https://api.zarban.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_unfilled_orders_websocket**](WebsocketApi.md#get_unfilled_orders_websocket) | **GET** /v2/ws | Websocket Upgrade


# **get_unfilled_orders_websocket**
> get_unfilled_orders_websocket()

Websocket Upgrade

Upgrade to websocket connection

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
    api_instance = zarban.service.openapi_client.WebsocketApi(api_client)
    
    try:
        # Websocket Upgrade
        api_instance.get_unfilled_orders_websocket()
    except ApiException as e:
        print("Exception when calling WebsocketApi->get_unfilled_orders_websocket: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**101** | Switching Protocols |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


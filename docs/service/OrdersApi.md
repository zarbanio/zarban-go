# zarban.service.openapi_client.OrdersApi

All URIs are relative to *https://api.zarban.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_unfilled_orders**](OrdersApi.md#get_unfilled_orders) | **GET** /v2/orders | Fetch Unfilled Orders
[**sync_order**](OrdersApi.md#sync_order) | **POST** /v2/orders/sync | Updates Order Entity


# **get_unfilled_orders**
> OrderResponse get_unfilled_orders(type=type, hash=hash, status=status, offerer=offerer, filler=filler, decay_start_time=decay_start_time, decay_end_time=decay_end_time, deadline=deadline, cursor=cursor, limit=limit)

Fetch Unfilled Orders

Get a list of unfilled orders filtered by different parameters.

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
    api_instance = zarban.service.openapi_client.OrdersApi(api_client)
    type = 'type_example' # str | Type of the order (optional)
hash = 'hash_example' # str | order hash (optional)
status = 'status_example' # str | Status of the order (optional)
offerer = 'offerer_example' # str | Ethereum address of the offerer (optional)
filler = 'filler_example' # str | Ethereum address of the filler (optional)
decay_start_time = zarban.service.openapi_client.TimeRange() # TimeRange | Decay start time (optional)
decay_end_time = zarban.service.openapi_client.TimeRange() # TimeRange | Decay end time (optional)
deadline = zarban.service.openapi_client.TimeRange() # TimeRange | Order deadline (optional)
cursor = 56 # int | Cursor for pagination (optional)
limit = 10 # int | Limit the number of orders returned (default is 10) (optional) (default to 10)

    try:
        # Fetch Unfilled Orders
        api_response = api_instance.get_unfilled_orders(type=type, hash=hash, status=status, offerer=offerer, filler=filler, decay_start_time=decay_start_time, decay_end_time=decay_end_time, deadline=deadline, cursor=cursor, limit=limit)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling OrdersApi->get_unfilled_orders: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type** | **str**| Type of the order | [optional] 
 **hash** | **str**| order hash | [optional] 
 **status** | **str**| Status of the order | [optional] 
 **offerer** | **str**| Ethereum address of the offerer | [optional] 
 **filler** | **str**| Ethereum address of the filler | [optional] 
 **decay_start_time** | [**TimeRange**](.md)| Decay start time | [optional] 
 **decay_end_time** | [**TimeRange**](.md)| Decay end time | [optional] 
 **deadline** | [**TimeRange**](.md)| Order deadline | [optional] 
 **cursor** | **int**| Cursor for pagination | [optional] 
 **limit** | **int**| Limit the number of orders returned (default is 10) | [optional] [default to 10]

### Return type

[**OrderResponse**](OrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response. |  -  |
**400** | Bad request (e.g. invalid filter parameters). |  -  |
**500** | Internal Server Error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **sync_order**
> Error sync_order(update_order_request)

Updates Order Entity

updates an order entity in database

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
    api_instance = zarban.service.openapi_client.OrdersApi(api_client)
    update_order_request = zarban.service.openapi_client.UpdateOrderRequest() # UpdateOrderRequest | 

    try:
        # Updates Order Entity
        api_response = api_instance.sync_order(update_order_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling OrdersApi->sync_order: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **update_order_request** | [**UpdateOrderRequest**](UpdateOrderRequest.md)|  | 

### Return type

[**Error**](Error.md)

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


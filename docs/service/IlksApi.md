# zarban.service.openapi_client.IlksApi

All URIs are relative to *https://api.zarban.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_all_ilks**](IlksApi.md#get_all_ilks) | **GET** /v2/ilks | Get all Ilks
[**get_ilk_by_name**](IlksApi.md#get_ilk_by_name) | **GET** /v2/ilks/{name} | Get Ilk by name


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
    api_instance = zarban.service.openapi_client.IlksApi(api_client)
    
    try:
        # Get all Ilks
        api_response = api_instance.get_all_ilks()
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling IlksApi->get_all_ilks: %s\n" % e)
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
    api_instance = zarban.service.openapi_client.IlksApi(api_client)
    name = 'name_example' # str | Name of the ILK

    try:
        # Get Ilk by name
        api_response = api_instance.get_ilk_by_name(name)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling IlksApi->get_ilk_by_name: %s\n" % e)
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


# zarban.service.openapi_client.PointsApi

All URIs are relative to *https://api.zarban.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_scoreboard**](PointsApi.md#get_scoreboard) | **GET** /v2/points/scoreboard | Get scoreboard


# **get_scoreboard**
> Scoreboard get_scoreboard()

Get scoreboard

Get scoreboard

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
    api_instance = zarban.service.openapi_client.PointsApi(api_client)
    
    try:
        # Get scoreboard
        api_response = api_instance.get_scoreboard()
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling PointsApi->get_scoreboard: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**Scoreboard**](Scoreboard.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A JSON object with scoreboard data |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


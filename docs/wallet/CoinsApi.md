# zarban.wallet.openapi_client.CoinsApi

All URIs are relative to *https://wapi.zarban.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_coin_details**](CoinsApi.md#get_coin_details) | **GET** /coins/{symbol} | Get coin
[**get_supported_coins**](CoinsApi.md#get_supported_coins) | **GET** /coins | Get coins


# **get_coin_details**
> Coin get_coin_details(symbol)

Get coin

Get the details of a coin.

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
    api_instance = zarban.wallet.openapi_client.CoinsApi(api_client)
    symbol = zarban.wallet.openapi_client.Symbol() # Symbol | Coin symbol

    try:
        # Get coin
        api_response = api_instance.get_coin_details(symbol)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling CoinsApi->get_coin_details: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | [**Symbol**](.md)| Coin symbol | 

### Return type

[**Coin**](Coin.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_supported_coins**
> CoinResponse get_supported_coins()

Get coins

Get the list of supported coins.

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
    api_instance = zarban.wallet.openapi_client.CoinsApi(api_client)
    
    try:
        # Get coins
        api_response = api_instance.get_supported_coins()
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling CoinsApi->get_supported_coins: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**CoinResponse**](CoinResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


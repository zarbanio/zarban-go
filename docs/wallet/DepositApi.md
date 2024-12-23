# zarban.wallet.openapi_client.DepositApi

All URIs are relative to *https://wapi.zarban.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**deposit_money**](DepositApi.md#deposit_money) | **GET** /deposit | Deposit money


# **deposit_money**
> DepositResponse deposit_money(network, symbol)

Deposit money

Deposit money to the wallet.

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
    api_instance = zarban.wallet.openapi_client.DepositApi(api_client)
    network = 'arbitrum' # str | Network to deposit
symbol = 'ZAR' # str | Coin symbol

    try:
        # Deposit money
        api_response = api_instance.deposit_money(network, symbol)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling DepositApi->deposit_money: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **network** | **str**| Network to deposit | 
 **symbol** | **str**| Coin symbol | 

### Return type

[**DepositResponse**](DepositResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**401** | Unauthorized |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# zarban.wallet.openapi_client.TransactionsApi

All URIs are relative to *https://wapi.zarban.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_user_transactions**](TransactionsApi.md#get_user_transactions) | **GET** /transactions | Get user transactions


# **get_user_transactions**
> TransactionResponse get_user_transactions(cursor=cursor, limit=limit)

Get user transactions

Get a list of the recent transactions of the user.

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
    api_instance = zarban.wallet.openapi_client.TransactionsApi(api_client)
    cursor = 0 # int | Cursor for pagination (optional) (default to 0)
limit = 100 # int | Limit the number of transactions returned (default is 100) (optional) (default to 100)

    try:
        # Get user transactions
        api_response = api_instance.get_user_transactions(cursor=cursor, limit=limit)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling TransactionsApi->get_user_transactions: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **int**| Cursor for pagination | [optional] [default to 0]
 **limit** | **int**| Limit the number of transactions returned (default is 100) | [optional] [default to 100]

### Return type

[**TransactionResponse**](TransactionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**401** | Unauthorized |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


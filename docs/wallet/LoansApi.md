# zarban.wallet.openapi_client.LoansApi

All URIs are relative to *https://wapi.zarban.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_loan_vault**](LoansApi.md#create_loan_vault) | **POST** /loans/create | Create vault
[**estimate_loan_collateral**](LoansApi.md#estimate_loan_collateral) | **GET** /loans/estimate | Get collateral and loan amount estimation
[**get_all_loan_plans**](LoansApi.md#get_all_loan_plans) | **GET** /loans/plans | Get all plan loans
[**get_loan_details**](LoansApi.md#get_loan_details) | **GET** /loans/{id} | Get loan
[**get_user_loans**](LoansApi.md#get_user_loans) | **GET** /loans | Get user loans
[**repay_loan**](LoansApi.md#repay_loan) | **POST** /loans/repay | Repay Loan


# **create_loan_vault**
> LoansResponse create_loan_vault(loan_create_request)

Create vault

Create a vault

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
    api_instance = zarban.wallet.openapi_client.LoansApi(api_client)
    loan_create_request = zarban.wallet.openapi_client.LoanCreateRequest() # LoanCreateRequest | 

    try:
        # Create vault
        api_response = api_instance.create_loan_vault(loan_create_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling LoansApi->create_loan_vault: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loan_create_request** | [**LoanCreateRequest**](LoanCreateRequest.md)|  | 

### Return type

[**LoansResponse**](LoansResponse.md)

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

# **estimate_loan_collateral**
> Currency estimate_loan_collateral(plan_name, loan_to_value_option, amount, input_type)

Get collateral and loan amount estimation

Get an estimate of required collateral for a specific loan amount, and vice versa.

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
    api_instance = zarban.wallet.openapi_client.LoansApi(api_client)
    plan_name = 'ETH-A' # str | The name of the loan plan
loan_to_value_option = zarban.wallet.openapi_client.LoanToValueOptions() # LoanToValueOptions | The desired loan to value option
amount = '1.234' # str | Loan/Collateral amount
input_type = 'loan' # str | The type of the input amount

    try:
        # Get collateral and loan amount estimation
        api_response = api_instance.estimate_loan_collateral(plan_name, loan_to_value_option, amount, input_type)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling LoansApi->estimate_loan_collateral: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plan_name** | **str**| The name of the loan plan | 
 **loan_to_value_option** | [**LoanToValueOptions**](.md)| The desired loan to value option | 
 **amount** | **str**| Loan/Collateral amount | 
 **input_type** | **str**| The type of the input amount | 

### Return type

[**Currency**](Currency.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | Bad request (e.g. invalid input parameters). |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_all_loan_plans**
> LoanPlanResponse get_all_loan_plans()

Get all plan loans

Get a list of the available plan loans.

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
    api_instance = zarban.wallet.openapi_client.LoansApi(api_client)
    
    try:
        # Get all plan loans
        api_response = api_instance.get_all_loan_plans()
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling LoansApi->get_all_loan_plans: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**LoanPlanResponse**](LoanPlanResponse.md)

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
**401** | Unauthorized |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_loan_details**
> LoansResponse get_loan_details(id)

Get loan

Get the details of a loan.

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
    api_instance = zarban.wallet.openapi_client.LoansApi(api_client)
    id = '1234567890' # str | Loan ID

    try:
        # Get loan
        api_response = api_instance.get_loan_details(id)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling LoansApi->get_loan_details: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| Loan ID | 

### Return type

[**LoansResponse**](LoansResponse.md)

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

# **get_user_loans**
> LoansResponseList get_user_loans(state=state, plan_name=plan_name)

Get user loans

Get a list of the user's loans.

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
    api_instance = zarban.wallet.openapi_client.LoansApi(api_client)
    state = 'active' # str | loan state (optional)
plan_name = 'DAIB' # str | loan plan name (optional)

    try:
        # Get user loans
        api_response = api_instance.get_user_loans(state=state, plan_name=plan_name)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling LoansApi->get_user_loans: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **str**| loan state | [optional] 
 **plan_name** | **str**| loan plan name | [optional] 

### Return type

[**LoansResponseList**](LoansResponseList.md)

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

# **repay_loan**
> LoansResponse repay_loan(repay_loan_request)

Repay Loan

Repay a loan totally

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
    api_instance = zarban.wallet.openapi_client.LoansApi(api_client)
    repay_loan_request = zarban.wallet.openapi_client.RepayLoanRequest() # RepayLoanRequest | 

    try:
        # Repay Loan
        api_response = api_instance.repay_loan(repay_loan_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling LoansApi->repay_loan: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repay_loan_request** | [**RepayLoanRequest**](RepayLoanRequest.md)|  | 

### Return type

[**LoansResponse**](LoansResponse.md)

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


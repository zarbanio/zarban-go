# zarban.wallet.openapi_client.AuthApi

All URIs are relative to *https://wapi.zarban.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**authenticate_with_telegram**](AuthApi.md#authenticate_with_telegram) | **POST** /auth/telegram | Authenticate with Telegram
[**generate_jwt_token**](AuthApi.md#generate_jwt_token) | **GET** /auth/token | Generate a JWT token
[**get_otp**](AuthApi.md#get_otp) | **GET** /auth/otp | Get OTP
[**login_with_email_and_password**](AuthApi.md#login_with_email_and_password) | **POST** /auth/login | Login with email and password
[**signup_with_email_and_password**](AuthApi.md#signup_with_email_and_password) | **POST** /auth/signup | signup with email and password
[**verify_user_email**](AuthApi.md#verify_user_email) | **GET** /verify-email | Verify email


# **authenticate_with_telegram**
> JwtResponse authenticate_with_telegram(auth_telegram_request)

Authenticate with Telegram

Authenticate with Telegram and get a JWT token.

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
    api_instance = zarban.wallet.openapi_client.AuthApi(api_client)
    auth_telegram_request = zarban.wallet.openapi_client.AuthTelegramRequest() # AuthTelegramRequest | 

    try:
        # Authenticate with Telegram
        api_response = api_instance.authenticate_with_telegram(auth_telegram_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling AuthApi->authenticate_with_telegram: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auth_telegram_request** | [**AuthTelegramRequest**](AuthTelegramRequest.md)|  | 

### Return type

[**JwtResponse**](JwtResponse.md)

### Authorization

No authorization required

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

# **generate_jwt_token**
> JwtResponse generate_jwt_token(duration)

Generate a JWT token

Generate a JWT token.

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
    api_instance = zarban.wallet.openapi_client.AuthApi(api_client)
    duration = 56 # int | Token duration in days

    try:
        # Generate a JWT token
        api_response = api_instance.generate_jwt_token(duration)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling AuthApi->generate_jwt_token: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **duration** | **int**| Token duration in days | 

### Return type

[**JwtResponse**](JwtResponse.md)

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

# **get_otp**
> SimpleResponse get_otp(channel)

Get OTP

Get OTP for the user.

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
    api_instance = zarban.wallet.openapi_client.AuthApi(api_client)
    channel = 'phone' # str | Channel to send OTP

    try:
        # Get OTP
        api_response = api_instance.get_otp(channel)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling AuthApi->get_otp: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **str**| Channel to send OTP | 

### Return type

[**SimpleResponse**](SimpleResponse.md)

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

# **login_with_email_and_password**
> JwtResponse login_with_email_and_password(login_request)

Login with email and password

Login with email and password and get a JWT token.

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
    api_instance = zarban.wallet.openapi_client.AuthApi(api_client)
    login_request = zarban.wallet.openapi_client.LoginRequest() # LoginRequest | 

    try:
        # Login with email and password
        api_response = api_instance.login_with_email_and_password(login_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling AuthApi->login_with_email_and_password: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **login_request** | [**LoginRequest**](LoginRequest.md)|  | 

### Return type

[**JwtResponse**](JwtResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**401** | Unauthorized |  -  |
**404** | Not Found |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **signup_with_email_and_password**
> SimpleResponse signup_with_email_and_password(sign_up_request)

signup with email and password

signup with email and password and get a JWT token.

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
    api_instance = zarban.wallet.openapi_client.AuthApi(api_client)
    sign_up_request = zarban.wallet.openapi_client.SignUpRequest() # SignUpRequest | 

    try:
        # signup with email and password
        api_response = api_instance.signup_with_email_and_password(sign_up_request)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling AuthApi->signup_with_email_and_password: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sign_up_request** | [**SignUpRequest**](SignUpRequest.md)|  | 

### Return type

[**SimpleResponse**](SimpleResponse.md)

### Authorization

No authorization required

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

# **verify_user_email**
> JwtResponse verify_user_email(token)

Verify email

Verify the email of the user.

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
    api_instance = zarban.wallet.openapi_client.AuthApi(api_client)
    token = 'token_example' # str | Verification token

    try:
        # Verify email
        api_response = api_instance.verify_user_email(token)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling AuthApi->verify_user_email: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **str**| Verification token | 

### Return type

[**JwtResponse**](JwtResponse.md)

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


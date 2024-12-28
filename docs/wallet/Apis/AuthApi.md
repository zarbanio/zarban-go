# AuthApi

All URIs are relative to *https://wapi.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**authenticateWithTelegram**](AuthApi.md#authenticateWithTelegram) | **POST** /auth/telegram | Authenticate with Telegram |
| [**generateJwtToken**](AuthApi.md#generateJwtToken) | **GET** /auth/token | Generate a JWT token |
| [**getOtp**](AuthApi.md#getOtp) | **GET** /auth/otp | Get OTP |
| [**loginWithEmailAndPassword**](AuthApi.md#loginWithEmailAndPassword) | **POST** /auth/login | Login with email and password |
| [**signupWithEmailAndPassword**](AuthApi.md#signupWithEmailAndPassword) | **POST** /auth/signup | signup with email and password |
| [**verifyUserEmail**](AuthApi.md#verifyUserEmail) | **GET** /verify-email | Verify email |


<a name="authenticateWithTelegram"></a>
# **authenticateWithTelegram**
> JwtResponse authenticateWithTelegram(AuthTelegramRequest)

Authenticate with Telegram

    Authenticate with Telegram and get a JWT token.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **AuthTelegramRequest** | [**AuthTelegramRequest**](../Models/AuthTelegramRequest.md)|  | |

### Return type

[**JwtResponse**](../Models/JwtResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="generateJwtToken"></a>
# **generateJwtToken**
> JwtResponse generateJwtToken(duration)

Generate a JWT token

    Generate a JWT token.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **duration** | **Integer**| Token duration in days | [default to null] [enum: 7, 15, 30, 90] |

### Return type

[**JwtResponse**](../Models/JwtResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getOtp"></a>
# **getOtp**
> SimpleResponse getOtp(channel)

Get OTP

    Get OTP for the user.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **channel** | **String**| Channel to send OTP | [default to null] [enum: phone] |

### Return type

[**SimpleResponse**](../Models/SimpleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="loginWithEmailAndPassword"></a>
# **loginWithEmailAndPassword**
> JwtResponse loginWithEmailAndPassword(LoginRequest)

Login with email and password

    Login with email and password and get a JWT token.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **LoginRequest** | [**LoginRequest**](../Models/LoginRequest.md)|  | |

### Return type

[**JwtResponse**](../Models/JwtResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="signupWithEmailAndPassword"></a>
# **signupWithEmailAndPassword**
> SimpleResponse signupWithEmailAndPassword(SignUpRequest)

signup with email and password

    signup with email and password and get a JWT token.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **SignUpRequest** | [**SignUpRequest**](../Models/SignUpRequest.md)|  | |

### Return type

[**SimpleResponse**](../Models/SimpleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="verifyUserEmail"></a>
# **verifyUserEmail**
> JwtResponse verifyUserEmail(token)

Verify email

    Verify the email of the user.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **token** | **String**| Verification token | [default to null] |

### Return type

[**JwtResponse**](../Models/JwtResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


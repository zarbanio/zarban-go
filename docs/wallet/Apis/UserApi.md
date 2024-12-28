# UserApi

All URIs are relative to *https://wapi.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**confirmKyc**](UserApi.md#confirmKyc) | **POST** /users/kyc/confirm | Confirm KYC |
| [**confirmPhoneNumber**](UserApi.md#confirmPhoneNumber) | **POST** /users/phone/confirm | Confirm phone number |
| [**createChildUser**](UserApi.md#createChildUser) | **POST** /users/children | create a child user |
| [**getUserProfile**](UserApi.md#getUserProfile) | **GET** /profile | Get profile |
| [**submitEmailConfirmationOtp**](UserApi.md#submitEmailConfirmationOtp) | **POST** /users/email/confirm | Submit email confirmation OTP |
| [**submitKyc**](UserApi.md#submitKyc) | **POST** /users/kyc | Submit KYC |
| [**verifyPhoneNumber**](UserApi.md#verifyPhoneNumber) | **POST** /users/phone | Verify phone number |
| [**verifyUserEmailAddress**](UserApi.md#verifyUserEmailAddress) | **POST** /users/email | Verify email |


<a name="confirmKyc"></a>
# **confirmKyc**
> SimpleResponse confirmKyc(KycConfirmRequest)

Confirm KYC

    Confirm KYC for the user.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **KycConfirmRequest** | [**KycConfirmRequest**](../Models/KycConfirmRequest.md)|  | |

### Return type

[**SimpleResponse**](../Models/SimpleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="confirmPhoneNumber"></a>
# **confirmPhoneNumber**
> confirmPhoneNumber(PhoneOtpSubmitRequest)

Confirm phone number

    Confirm the phone number of the user.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **PhoneOtpSubmitRequest** | [**PhoneOtpSubmitRequest**](../Models/PhoneOtpSubmitRequest.md)|  | |

### Return type

null (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createChildUser"></a>
# **createChildUser**
> User createChildUser(CreateChildUserRequest)

create a child user

    create a child user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CreateChildUserRequest** | [**CreateChildUserRequest**](../Models/CreateChildUserRequest.md)|  | |

### Return type

[**User**](../Models/User.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="getUserProfile"></a>
# **getUserProfile**
> ProfileResponse getUserProfile()

Get profile

    Get the profile of the user.

### Parameters
This endpoint does not need any parameter.

### Return type

[**ProfileResponse**](../Models/ProfileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="submitEmailConfirmationOtp"></a>
# **submitEmailConfirmationOtp**
> SimpleResponse submitEmailConfirmationOtp(EmailOtpSubmitRequest)

Submit email confirmation OTP

    Submit email confirmation OTP for the user.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **EmailOtpSubmitRequest** | [**EmailOtpSubmitRequest**](../Models/EmailOtpSubmitRequest.md)|  | |

### Return type

[**SimpleResponse**](../Models/SimpleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="submitKyc"></a>
# **submitKyc**
> KycResponse submitKyc(KycRequest)

Submit KYC

    Submit KYC for the user.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **KycRequest** | [**KycRequest**](../Models/KycRequest.md)|  | |

### Return type

[**KycResponse**](../Models/KycResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="verifyPhoneNumber"></a>
# **verifyPhoneNumber**
> verifyPhoneNumber(UpdatePhoneRequest)

Verify phone number

    Verify the phone number of the user.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdatePhoneRequest** | [**UpdatePhoneRequest**](../Models/UpdatePhoneRequest.md)|  | |

### Return type

null (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="verifyUserEmailAddress"></a>
# **verifyUserEmailAddress**
> SimpleResponse verifyUserEmailAddress(UpdateEmailRequest)

Verify email

    Verify the email of the user.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdateEmailRequest** | [**UpdateEmailRequest**](../Models/UpdateEmailRequest.md)|  | |

### Return type

[**SimpleResponse**](../Models/SimpleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


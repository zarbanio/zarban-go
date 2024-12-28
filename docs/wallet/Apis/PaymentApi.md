# PaymentApi

All URIs are relative to *https://wapi.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createPayment**](PaymentApi.md#createPayment) | **POST** /payments | Create a payment |


<a name="createPayment"></a>
# **createPayment**
> Payment createPayment(PaymentRequest)

Create a payment

    Create a new payment.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **PaymentRequest** | [**PaymentRequest**](../Models/PaymentRequest.md)|  | |

### Return type

[**Payment**](../Models/Payment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


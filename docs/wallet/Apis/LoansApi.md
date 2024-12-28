# LoansApi

All URIs are relative to *https://wapi.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createLoanVault**](LoansApi.md#createLoanVault) | **POST** /loans/create | Create vault |
| [**estimateLoanCollateral**](LoansApi.md#estimateLoanCollateral) | **GET** /loans/estimate | Get collateral and loan amount estimation |
| [**getAllLoanPlans**](LoansApi.md#getAllLoanPlans) | **GET** /loans/plans | Get all plan loans |
| [**getLoanDetails**](LoansApi.md#getLoanDetails) | **GET** /loans/{id} | Get loan |
| [**getUserLoans**](LoansApi.md#getUserLoans) | **GET** /loans | Get user loans |
| [**repayLoan**](LoansApi.md#repayLoan) | **POST** /loans/repay | Repay Loan |


<a name="createLoanVault"></a>
# **createLoanVault**
> LoansResponse createLoanVault(LoanCreateRequest)

Create vault

    Create a vault

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **LoanCreateRequest** | [**LoanCreateRequest**](../Models/LoanCreateRequest.md)|  | |

### Return type

[**LoansResponse**](../Models/LoansResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="estimateLoanCollateral"></a>
# **estimateLoanCollateral**
> Currency estimateLoanCollateral(planName, loanToValueOption, amount, inputType)

Get collateral and loan amount estimation

    Get an estimate of required collateral for a specific loan amount, and vice versa.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **planName** | **String**| The name of the loan plan | [default to null] |
| **loanToValueOption** | [**LoanToValueOptions**](../Models/.md)| The desired loan to value option | [default to null] [enum: Risky, Normal, Safe] |
| **amount** | **String**| Loan/Collateral amount | [default to null] |
| **inputType** | **String**| The type of the input amount | [default to null] [enum: loan, collateral] |

### Return type

[**Currency**](../Models/Currency.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAllLoanPlans"></a>
# **getAllLoanPlans**
> LoanPlanResponse getAllLoanPlans()

Get all plan loans

    Get a list of the available plan loans.

### Parameters
This endpoint does not need any parameter.

### Return type

[**LoanPlanResponse**](../Models/LoanPlanResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getLoanDetails"></a>
# **getLoanDetails**
> LoansResponse getLoanDetails(id)

Get loan

    Get the details of a loan.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | **String**| Loan ID | [default to null] |

### Return type

[**LoansResponse**](../Models/LoansResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserLoans"></a>
# **getUserLoans**
> LoansResponseList getUserLoans(state, planName)

Get user loans

    Get a list of the user&#39;s loans.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **state** | **String**| loan state | [optional] [default to null] [enum: pending, active, repayment-ongoing, settled, creation-failed, settlement-failed] |
| **planName** | **String**| loan plan name | [optional] [default to null] |

### Return type

[**LoansResponseList**](../Models/LoansResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="repayLoan"></a>
# **repayLoan**
> LoansResponse repayLoan(RepayLoanRequest)

Repay Loan

    Repay a loan totally

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **RepayLoanRequest** | [**RepayLoanRequest**](../Models/RepayLoanRequest.md)|  | |

### Return type

[**LoansResponse**](../Models/LoansResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


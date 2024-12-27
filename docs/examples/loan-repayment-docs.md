# Loan Repayment Documentation

## Overview

This document describes how to use the Zarban SDK to repay loans. The SDK provides functionality for previewing repayments, executing repayments, and monitoring their status through a simple go interface.

## Prerequisites

- go >= 1.21.0
- Zarban SDK (`zarban.wallet`)
- Valid API access token
- Child user credentials (if applicable)

## Installation

```bash
go get github.com/zarbanio/zarban-go
```

## Authentication

The SDK requires two forms of authentication:

1. An API access token
2. A child user header (optional, depending on use case)

```go
const ACCESS_TOKEN = "your_access_token_here"

// Setup API client
// Set the X-Child-User header
headers := map[string]string{
    "Authorization": "Bearer " + ACCESS_TOKEN,
    "X-Child-User":  "child_user_test",
}
client, err := wallet.NewClient(
    "https://testwapi.zarban.io",
    wallet.WithRequestEditorFn(wallet.AddHeaders(headers)),
)
if err != nil {
    log.Fatalf("Failed to create wallet client: %v", err)
    return
}
```

## Core Functions

### Repaying a Loan

#### `func repayLoan(client, loanId, intent) (wallet.LoansResponse, error)`

Handles loan repayment operations, including preview and actual repayment.

**Parameters:**

- `client`: wallet.Client instance
- `loanId`: string - Unique identifier for the loan
- `intent`: wallet.RepayLoanRequestIntent - Either "Preview" or "Repay"

**Returns:**

- API response object containing repayment details if successful
- None if operation fails

**Example:**

```go
previewLoanResponse, err := repayLoan(*client, LOAN_ID, wallet.RepayLoanRequestIntentPreview)
if err != nil {
    // you can do some addition works with error here!
    return
}
```

### Checking Loan Status

#### `func getLoanStatus(client, loanId string) (wallet.LoansResponse, error)`

Retrieves and displays detailed information about a loan.

**Parameters:**

- `client`: wallet.Client
- `loanId`: string - Unique identifier for the loan

**Returns:**
Loan details object containing:

- Loan status
- User ID
- Liquidation price
- Collateral amount
- Collateralization ratio
- Loan to value
- Debt amount
- Loan plan

**Example:**

```go
loanDetails, err := getLoanStatus(
    *client,
    repaymentLoanResponse.Id,
)
if err != nil {
    // you can do some addition works with error here!
    return
}
```

## API Endpoints

### POST /loans/repay

Previews or executes a loan repayment.

**Request Body:**

```json
{
  "intent": "Preview", // or "Repay"
  "loan_id": "loan123"
}
```

**Response (200 OK):**

```json
{
  "id": "vault123",
  "userId": 12345,
  "liquidationPrice": {
    "USD": "1500.00",
    "ETH": "0.75"
  },
  "collateral": {
    "ETH": "1.0",
    "USD": "2000.00"
  },
  "collateralizationRatio": "1.5",
  "loanToValue": "0.66",
  "debt": {
    "DAI": "1000.00",
    "USD": "1000.00"
  }
  // other fields
}
```

### GET /loans/{id}

Retrieves loan details.

**Path Parameters:**

- `id`: Loan identifier

**Response (200 OK):**
Returns loan details including state, collateral, debt, and other relevant information.

## Error Handling

The SDK uses `withErrorHandler` for error handling. Common errors include:

- 400: Bad Request
- 401: Unauthorized
- 500: Internal Server Error

Example error handling:

```go
if err != nil {
    // you can do some addition works with error here!
    return
}
```

## Complete Usage Example

```go
// Replace with your actual access token
const ACCESS_TOKEN = "your_access_token_here"

// Setup API client
// Set the X-Child-User header
headers := map[string]string{
    "Authorization": "Bearer " + ACCESS_TOKEN,
    "X-Child-User":  "child_user_test",
}
client, err := wallet.NewClient(
    "https://testwapi.zarban.io",
    wallet.WithRequestEditorFn(wallet.AddHeaders(headers)),
)
if err != nil {
    log.Fatalf("Failed to create wallet client: %v", err)
    return
}

// Loan ID to repay, replace with actual loan ID
const LOAN_ID = "DAIA#2825"

// Preview repayment
fmt.Println("Previewing loan repayment...")

previewLoanResponse, err := repayLoan(*client, LOAN_ID, wallet.RepayLoanRequestIntentPreview)
if err != nil {
    // you can do some addition works with error here!
    return
}

fmt.Println("\nRepayment preview details:")
fmt.Println("Collateral to be returned: ,", previewLoanResponse.Collateral)
fmt.Println("Debt to be repaid: ", previewLoanResponse.Debt)

var confirm string
fmt.Print("\nDo you want to proceed with the repayment? (y/n): ")
fmt.Scanln(&confirm)

if confirm == "y" {
    // Proceed with actual repayment
    repaymentLoanResponse, err := repayLoan(*client, LOAN_ID, wallet.RepayLoanRequestIntentRepay)
    if err != nil {
        // you can do some addition works with error here!
        return
    }
    fmt.Print("repayment in progress...")
    for {
        loanDetails, err := getLoanStatus(
            *client,
            repaymentLoanResponse.Id,
        )
        if err != nil {
            // you can do some addition works with error here!
            return
        }
        if loanDetails.State["LocaleEn"] == "Loan settled" {
            fmt.Print("\nLoan repayment successful!")
            getLoanStatus(*client, repaymentLoanResponse.Id)
            break
        } else if loanDetails.State["LocaleEn"] == "Loan settlement failed" {
            fmt.Print(loanDetails.State["LocaleEn"])
            break
        }
        time.Sleep(1 * time.Second)
    }
} else {
    fmt.Print("Repayment cancelled.")
}
```

## Best Practices

1. Always preview repayment before executing
2. Implement proper error handling for all API calls
3. Monitor repayment status after execution
4. Remove the X-Child-User header after use
5. Store sensitive information (like access tokens) in environment variables
6. Implement proper timeout handling for status monitoring

## Limitations

- Repayment preview is recommended before actual repayment
- Status monitoring may require multiple API calls
- API access tokens should be kept secure and not hardcoded

## Support

For additional support or bug reports, please contact the Zarban support team.

## See Also

- [API Reference Documentation](../wallet)
- [Loan Creation Documentation](./loan-creation-docs.md)

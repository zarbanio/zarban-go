# Loan Creation Documentation

## Overview

This document describes how to use the Zarban SDK to create loans. The SDK provides functionality for creating new loans and monitoring their status through a simple go interface.

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
// Replace with your actual access token
const ACCESS_TOKEN = "your_access_token_here"

// Setup API client
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

### Creating a Loan

#### `func createLoan(client, planName, collateral, debt, symbol , loanToValueOption) (walletLoansResponse, error)`

Creates a new loan using the specified parameters.

**Parameters:**

- `client`: wallet.Client instance
- `planName`: string - Name of the loan plan (Currently supports "DAIA" and "DAIB")
- `collateral`: string - Amount of collateral
- `debt`: string - Amount of debt
- `symbol`: string - Coin symbol (e.g., "DAI")
- `loanToValueOption`: wallet.LoanToValueOptions - Risk level ("Safe", "Normal", or "Risky")

**Important Notes:**

- Either `collateral` or `debt` must be empty
- Returns the loan ID if successful, None if failed

**Example:**

```go
createLoanResponse, err := createLoan(*client, PLAN_NAME, COLLATERAL, DEBT, SYMBOL, LOAN_TO_VALUE_OPTIONS)
if err != nil {
    return
}
```

### Checking Loan Status

#### `func loanStatus(client , loanId)`

Retrieves and displays the current state of a loan.

**Parameters:**

- `client`: wallet.Client instance
- `loanId`: string - The ID of the loan to check

**Returns:**
Loan details object containing:

- Status
- Collateral amount
- Debt amount
- Interest rate
- Creation date

**Example:**

```go
loanDetails, err := loanStatus(
    *client,
    createLoanResponse.Id,
)
if err != nil {
    // you can do some addition works with error here!
    return
}
```

## API Endpoints

### POST /loans/create

Creates a new vault/loan.

**Request Body:**

```json
{
  "intent": "Create",
  "planName": "DAIA",
  "collateral": "1000",
  "debt": "",
  "symbol": "USDT",
  "loanToValueOption": "Safe"
}
```

**Response (200 OK):**

```json
{
  "id": "1234567890"
}
```

### GET /loans/{id}

Retrieves loan details.

**Path Parameters:**

- `id`: Loan identifier

**Response (200 OK):**
Returns loan details including status, collateral, debt, and other relevant information.

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

// Loan creation parameters, Replace them with yout actual data
// *** either collateral or debt must be empty ***
const PLAN_NAME = "DAIA"             // Only DAIA and DAIB supported
const COLLATERAL = "1000"            // Amount of collateral
const DEBT = ""                      // Amount of debt
const SYMBOL = "DAI"                 // Coin symbol
const LOAN_TO_VALUE_OPTIONS = "Safe" // Risky - Normal - Safe

createLoanResponse, err := createLoan(*client, PLAN_NAME, COLLATERAL, DEBT, SYMBOL, LOAN_TO_VALUE_OPTIONS)
if err != nil {
    return
}

// Remove the X-Child-User header after use
headers = map[string]string{
    "Authorization": "Bearer " + ACCESS_TOKEN,
}
client, err = wallet.NewClient(
    "https://testwapi.zarban.io",
    wallet.WithRequestEditorFn(wallet.AddHeaders(headers)),
)
if err != nil {
    log.Fatalf("Failed to create wallet client: %v", err)
    return
}

// Track loan status
fmt.Println("\nTracking loan status...")
loanDetails, err := loanStatus(
    *client,
    createLoanResponse.Id,
)
if err != nil {
    // you can do some addition works with error here!
    return
}
// You can add more specific actions based on the loan state
fmt.Println("Loan status: ", loanDetails.State)
```

## Best Practices

1. Always remove the X-Child-User header after use:
   ```go
    headers = map[string]string{
        "Authorization": "Bearer " + ACCESS_TOKEN,
    }
    client, err = wallet.NewClient(
        "https://testwapi.zarban.io",
        wallet.WithRequestEditorFn(wallet.AddHeaders(headers)),
    )
   ```
2. Implement proper error handling for all API calls
3. Validate input parameters before making API calls
4. Store sensitive information (like access tokens) in environment variables

## Limitations

- Only "DAIA" and "DAIB" plans are currently supported
- Either collateral or debt must be empty when creating a loan
- API access tokens should be kept secure and not hardcoded

## Support

For additional support or bug reports, please contact the Zarban support team.

## See Also

- [API Reference Documentation](../wallet)

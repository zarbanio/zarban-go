# Stablecoin System Vault Creation Documentation

## Overview

This documentation covers the implementation of vault creation in the Zarban Stablecoin System. The SDK provides functionality for creating vaults, handling transactions, and managing collateral using go-ethereum and the Zarban API.

## Prerequisites

- go > 1.21.0
- go-ethereum
- Zarban SDK
- Ethereum node access (RPC URL)
- Private key with sufficient funds
- Required node packages:
  ```bash
  go get github.com/zarbanio/zarban-go
  ```

## Core Components

### 1. Collateral Type Management

#### `func getIlksSymbol(client service.Client) ([]service.Symbol, error)`

Retrieves available collateral types from the system.

**Parameters:**

- `client (service.Client)`: The Client Api instance

**Returns:**

- `[]service.Symbol`: Unique collateral type symbols

**Example:**

```go
symbols, err := getIlksSymbol(client)
if err != nil {
    return "", fmt.Errorf("failed to get symbols: %w", err)
}
```

### 2. Amount Conversion

#### `func toNative(client, symbol, amount) (string, error)`

Converts human-readable amounts to blockchain native format.

**Parameters:**

- `client (service.Client)`: The Client Api instance
- `symbol (service.Symbol)`: Asset symbol (e.g., "ETH", "ZAR")
- `amount (float64)`: Human-readable amount

**Returns:**

- `(string, error)`: Amount in native blockchain format (wei), error

**Example:**

```go
nativeCollateralAmount, err := toNative(client, "ETH", 0.01)
if err != nil {
    log.Fatalf("Error converting collateral amount: %v", err)
    return service.ChainActivity{}, err
}
```

### 3. Vault Transaction Processing

#### `func getVaultTxSteps(client, ilkName, symbol, walletAddress, collateralAmount, loanAmount) (service.ChainActivity, error)`

Retrieves transaction steps for vault creation.

**Parameters:**

- `client (service.Client)`: The Client Api instance
- `ilkName (string)`: Collateral type name
- `symbol ( service.Symbol)`: Asset symbol
- `walletAddress (string)`: User's wallet address
- `collateralAmount (float64)`: Collateral amount
- `loanAmount (float64)`: Loan amount in stablecoin

**Returns:**

- `service.ChainActivity, error`: (ChainActivity, error)

### 4. Transaction Management

#### `func waitForTransactionReceipt(client, txHash, maxWaitTime, checkInterval) (*types.Receipt, error)`

Waits for transaction confirmation.

**Parameters:**

- `client (*ethclient.Client)`: ethclient instance
- `txHash (common.Hash)`: Transaction hash
- `maxWaitTime (time.Duration)`: Maximum wait time in seconds
- `checkInterval (time.Duration)`: Check interval in seconds

**Returns:**

- (Transaction receipt or None, error)

### 5. Transaction Logging

#### `func SaveTransactionDetails(tx, txHash, vaultID)`

Saves transaction details to a JSON file.

**Parameters:**

- `tx (EthereumTransaction)`: Transaction object
- `txHash (string)`: Transaction hash
- `vaultId (int)`: Vault identifier (can be None)

## Complete Implementation Example

```go
// Configuration
const HTTPS_RPC_URL = "Replace with your Ethereum node URL"
const PRIVATE_KEY = "Replace with your Private key"
privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
if err != nil {
    log.Fatalf("Failed to load private key: %v", err)
}
WALLET_ADDRESS, err := GetAddressFromPrivateKey(PRIVATE_KEY)
if err != nil {
    log.Fatalf("Failed to get address from private key: %v", err)
    return
}

// Create and configure the client
client, err := service.NewClient("https://testapi.zarban.io")
if err != nil {
    log.Fatalf("Failed to create wallet client: %v", err)
    return
}

// Connect to an Ethereum client
ethClient, err := ethclient.Dial(HTTPS_RPC_URL)
if err != nil {
    log.Fatalf("Failed to connect to Ethereum client: %v", err)
    return
}

account := WALLET_ADDRESS

// Get the chain ID
chainID, err := ethClient.NetworkID(context.Background())
if err != nil {
    log.Fatalf("Failed to get chain ID: %v", err)
    return
}

// Define vault creation parameters
const ILK_NAME = "ETHA"        // Replace with your desired ilk
const SYMBOL = "ETH"           // Replace with the symbol associated with your ilk
const COLLATERAL_AMOUNT = .001 // Replace with your desired amount
const LOAN_AMOUNT = 100        // Replace with your desired amount

vaultSteps, err := getVaultTxSteps(
    *client,
    ILK_NAME,
    SYMBOL,
    WALLET_ADDRESS,
    COLLATERAL_AMOUNT,
    LOAN_AMOUNT)
if err != nil {
    log.Fatalf("Failed to get vault creartion steps: %v", err)
    return
}

numOfSteps := vaultSteps.NumberOfSteps
stepNumber := vaultSteps.StepNumber
steps := vaultSteps.Steps
```

## Transaction Flow

1. **Initialization**

   - Set up API client and go-ethereum connection
   - Configure vault parameters

2. **Transaction Steps**

   - Retrieve vault creation steps
   - Process each step sequentially
   - Handle transaction signing and submission

3. **Transaction Monitoring**

   - Wait for transaction confirmation
   - Log transaction details
   - Verify vault creation

4. **Result Verification**
   - Get transaction logs
   - Extract vault ID
   - Update transaction records

## Error Handling

```go
logs, err := getLogs(*client, txHash)
if err != nil {
    // handle errors here
    log.Fatalf("Failed to get logs: %v", err)
    return
}
```

## Best Practices

1. **Security**

   - Never hardcode private keys
   - Use environment variables for sensitive data
   - Verify all transaction parameters before signing

2. **Transaction Management**

   - Always wait for transaction confirmations
   - Implement proper error handling
   - Log all transactions for audit purposes

3. **Gas Management**
   - Use appropriate gas limits
   - Monitor gas prices
   - Implement retry mechanisms for failed transactions

## Logging and Monitoring

The system maintains a transaction log file (`transaction_log.json`) with the following information:

- Timestamp
- Transaction details
- Transaction hash
- Vault ID (when available)

Example log entry:

```json
{
  "timestamp": "2024-03-13T10:00:00",
  "tx": {
    "from": "0x...",
    "to": "0x...",
    "value": "1000000000000000",
    "gas": 200000
  },
  "txHash": "0x...",
  "vaultId": "123"
}
```

## Limitations and Considerations

1. **Network Dependencies**

   - Requires stable network connection
   - Dependent on Ethereum network status
   - May be affected by network congestion

2. **Resource Requirements**

   - Sufficient ETH for gas fees
   - Adequate collateral for vault creation
   - Valid API access

3. **Transaction Timing**
   - Transaction confirmation times vary
   - Maximum wait time configurable
   - May require multiple steps

## Troubleshooting

1. **Connection Issues**

```go
ethClient, err := ethclient.Dial(HTTPS_RPC_URL)
if err != nil {
    log.Fatalf("Failed to connect to Ethereum client: %v", err)
    return
}
```

2. **Transaction Failures**

   - Check gas prices and limits
   - Verify account balance
   - Confirm network status

3. **API Errors**
   - Validate API credentials
   - Check request parameters
   - Review error responses

## Support

For additional support or bug reports, please contact the Zarban support team or refer to the API documentation.

## See Also

- [API Reference Documentation](../service/Apis/StableCoinSystemApi.md)

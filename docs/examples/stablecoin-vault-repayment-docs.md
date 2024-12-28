# Stablecoin System Vault Repayment Documentation

## Overview

This documentation covers the implementation of vault repayment in the Zarban Stablecoin System. The SDK provides functionality for repaying existing vaults, handling transactions, and logging repayment details using go-ethereum and the Zarban API.

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

### 1. Amount Conversion

#### `func toNative(amount) (*string, error)`

Converts a human-readable amount to its native blockchain format (wei).

**Parameters:**

- `amount (float64)`: The amount to be converted

**Returns:**

- `(*string, error)`: The amount in native blockchain format 

**Example:**

```go
nativeAmount, err = toNative(amount)
if err != nil {
    log.Fatalf("Error converting collateral amount: %v", err)
    return service.ChainActivity{}, err
}
```

### 2. Vault Repayment Transaction Processing

#### `func getVaultTxSteps(client, walletAddress, vaultId, amount) (service.ChainActivity, error)`

Retrieves transaction steps for vault repayment.

**Parameters:**

- `client (service.Client)`: The StableCoinSystem Api instance
- `walletAddress (string)`: User's wallet address
- `vaultId (int)`: The ID of the vault to be repaid
- `amount (float64)`: The repayment amount

**Returns:**

- `(service.ChainActivity, error)`: (ChainActivity, error)

### 3. Transaction Management

#### `func waitForTransactionReceipt(client, txHash, maxWaitTime, checkInterval) (*types.Receipt, error)`

Waits for transaction confirmation.

**Parameters:**

- `client (*ethclient.Client)`: ethclient instance
- `txHash (common.Hash)`: Transaction hash
- `maxWaitTime (time.Duration)`: Maximum wait time in seconds
- `checkInterval (time.Duration)`: Check interval in seconds

**Returns:**

- (Transaction receipt or None, error)

### 4. Transaction Logging

#### `func SaveTransactionDetails(tx, txHash, vaultID)`

Saves transaction details to a JSON file.

**Parameters:**

- `tx (EthereumTransaction)`: Transaction object
- `txHash (string)`: Transaction hash
- `vaultId (int)`: Vault identifier (can be None)

## Complete Implementation Example

```go
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

const VAULT_ID int = 0   // Update with the actual vault ID
const AMOUNT float64 = 0 // Update with the amount to repay

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
vaultSteps, err := getVaultTxSteps(
    *client,
    WALLET_ADDRESS,
    VAULT_ID,
    AMOUNT)
if err != nil {
    log.Fatalf("Failed to get vault repayment steps: %v", err)
    return
}

numOfSteps := vaultSteps.NumberOfSteps
stepNumber := vaultSteps.StepNumber
steps := vaultSteps.Steps
```

## Transaction Flow

1. **Initialization**

   - Set up API client and go-ethereum connection
   - Configure vault repayment parameters

2. **Transaction Steps**

   - Retrieve vault repayment steps
   - Process each step sequentially
   - Handle transaction signing and submission

3. **Transaction Monitoring**

   - Wait for transaction confirmation
   - Log transaction details

4. **Result Verification**
   - Review transaction logs for successful repayment

## Error Handling

```go
vaultSteps, err = getVaultTxSteps(
    *client,
    WALLET_ADDRESS,
    VAULT_ID,
    AMOUNT)
if err != nil {
    log.Fatalf("Failed to get vault repayment steps: %v", err)
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

The system maintains a transaction log file (`repay_transaction_log.json`) with the following information:

- Timestamp
- Transaction details
- Transaction hash

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
  "vaultId": "ETH#43"
}
```

## Limitations and Considerations

1. **Network Dependencies**

   - Requires stable network connection
   - Dependent on Ethereum network status
   - May be affected by network congestion

2. **Resource Requirements**

   - Sufficient ETH for gas fees
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

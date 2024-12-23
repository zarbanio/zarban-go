# Zarban SDK

<p align="center">
  <img src="https://zarban.io/favicon.ico" width="400" alt="Logo">
</p>

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Zarban SDK is a Go interface for interacting with the Zarban DeFi protocol, enabling developers to seamlessly integrate lending and borrowing functionalities into their applications. This SDK simplifies complex DeFi operations by providing easy-to-use methods for lending assets, managing collateral, borrowing funds, and monitoring positions in the Zarban protocol.

## Features

- **Automated API Client Generation**: Built using OpenAPI specification, ensuring type safety and up-to-date API compatibility
- **Lending Operations**: Easily deposit assets, view lending rates, and manage lending positions
- **Borrowing Management**: Streamlined methods for borrowing assets, managing collateral, and monitoring loan health
- **Position Tracking**: Real-time access to user positions, including borrowed amounts, collateral ratios, and liquidation thresholds
- **Market Data**: Simple methods to fetch current interest rates, available liquidity, and market statistics
- **Type Safety**: Full type hints support for Go static type checking
- **Error Handling**: Comprehensive error handling with detailed exceptions for DeFi operations
- **Async Support**: Asynchronous methods for improved performance in high-throughput applications

## Environments

Zarban SDK supports two distinct environments:

1. **Mainnet**: The production environment for the Zarban DeFi protocol.

   - Wallet API: `https://wapi.zarban.io`
   - Service API: `https://api.zarban.io`

2. **Testnet**: A separate testing environment for the Zarban protocol.
   - Wallet API: `https://testwapi.zarban.io`
   - Service API: `https://testapi.zarban.io`

Be sure to use the appropriate environment configuration when interacting with the Zarban SDK.

## Installation

```bash
go get github.com/zarbanio/zarban-go
```

## Quick Start

Zarban SDK provides access to two distinct APIs:

### 1. Wallet API (`zarban.wallet`)

The Wallet API handles user authentication and wallet management operations.

### 2. Service API(`zarban.service`)

The Zarban Service API provides access to core DeFi protocol operations.

```go
import (
	"context"
	"log"

	"github.com/zarbanio/zarban-go/wallet"
)

client, err := wallet.NewClient("https://testwapi.zarban.io")
if err != nil {
    log.Fatalf("Failed to create wallet client: %v", err)
    return
}

httpResponse, err := client.someMethod(context.Background())
if err != nil {
    log.Fatalf("Error during API call: %v", err)
    return
}
```

## Usage Examples

For detailed usage examples, see our [Examples Documentation](docs/examples).

### Advanced Usage

Here's a simple example to sign up and get started with Zarban:

```go
import (
	"context"
	"fmt"
	"log"

	"github.com/zarbanio/zarban-go/wallet"
)

func SignupExample() {
	// Create and configure the client
	client, err := wallet.NewClient("https://testwapi.zarban.io")
	if err != nil {
		log.Fatalf("Failed to create wallet client: %v", err)
		return
	}
	// Prepare the signup request data
	signUpRequest := wallet.SignUpRequest{
		Email:    "user@example.com",
		Password: "yourSecurePassword",
	}

	httpResponse, err := client.SignupWithEmailAndPassword(context.Background(), signUpRequest)
	if err != nil {
		log.Fatalf("Error during API call: %v", err)
		return
	}

	var successResponse wallet.SimpleResponse
	err = wallet.HandleAPIResponse(context.Background(), httpResponse, &successResponse)
	if err != nil {
		if apiErr, ok := err.(*wallet.APIError); ok {
			fmt.Println(wallet.PrettyPrintError(apiErr))
		} else {
			log.Printf("Unexpected error: %v", err)
		}
		return
	}

	fmt.Printf("Signup successful: %+v\n", successResponse.Messages)
}
```

## Configuration

The SDK can be configured with various options to customize its behavior and authentication methods.

### Basic Configuration

```Go
import "github.com/zarbanio/zarban-go/wallet"

// Basic configuration with just the host URL
client, err := wallet.NewClient("https://testwapi.zarban.io")
if err != nil {
    log.Fatalf("Failed to create wallet client: %v", err)
    return
}
```

### Authentication Options

The SDK supports multiple authentication methods:

1. API Key Authentication:

```Go
// Define headers to be added
headers := map[string]string{
    "Authorization": "Bearer " + loginResponse.Token,
}

// configure it with the header editing function
client, err = wallet.NewClient(
    "https://testwapi.zarban.io",
    wallet.WithRequestEditorFn(wallet.AddHeaders(headers)),
)
if err != nil {
    log.Fatalf("Failed to create client: %v", err)
}
```

## Error Handling

To make error handling easier, we provide a utility function named HandleAPIResponse. This function simplifies the process of managing errors and helps avoid repetitive if/else(or switch/case) blocks in your code.

While using HandleAPIResponse is not mandatory, we highly recommend it for cleaner and more maintainable code. If you prefer, you can always handle errors manually using traditional if/else(or switch/case) blocks.

### Usage example:

Using HandleAPIResponse

```go
httpResponse, err = client.CreateChildUser(context.Background(), createChildUserRequest)
if err != nil {
    log.Fatalf("Error during API call -> CreateChildUser: %v", err)
    return
}

var createChildResponse wallet.User
err = wallet.HandleAPIResponse(context.Background(), httpResponse, &createChildResponse)
if err != nil {
    if apiErr, ok := err.(*wallet.APIError); ok {
        fmt.Println(wallet.PrettyPrintError(apiErr))
    } else {
        log.Printf("Unexpected error: %v", err)
    }
    return
}
```

Manual Error Handling

```go
httpResponse, err = client.CreateChildUser(context.Background(), createChildUserRequest)
if err != nil {
    log.Fatalf("Error during API call: %v", err)
    return
}

createChildResponse, err := wallet.ParseCreateChildUserResponse(httpResponse)
if err != nil {
    log.Fatalf("Error while parsing http response: %v", err)
    return
}

switch c.StatusCode() {
case 200:
    fmt.Printf("Child user created successfully. User: %s\n", *c.JSON200.Username)
    return c.JSON200, nil
case 400:
    return c.JSON400, fmt.Errorf("bad request: %s", c.JSON400.Msg)
case 500:
    return c.JSON500, fmt.Errorf("internal server error: %s", c.JSON500.Msg)
default:
    return nil, fmt.Errorf("unexpected status code: %d", c.StatusCode())
}
```

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

1. Fork the repository
2. Create a new branch
3. Make your changes
4. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- Create an issue on GitHub
- Email: info@zarban.io
- Documentation: [https://docs.zarban.io](https://docs.zarban.io)

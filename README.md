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

   - Service API: `https://api.zarban.io`

2. **Testnet**: A separate testing environment for the Zarban protocol.
   - Service API: `https://testapi.zarban.io`

Be sure to use the appropriate environment configuration when interacting with the Zarban SDK.

## Installation

```bash
go get github.com/zarbanio/zarban-go
```

## Quick Start

The Zarban Service API provides access to core DeFi protocol operations.

```go
import (
	"context"
	"log"

	"github.com/zarbanio/zarban-go/service"
)

client, err := service.NewClient("https://testapi.zarban.io")
if err != nil {
    log.Fatalf("Failed to create service client: %v", err)
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

## Configuration

The SDK can be configured with various options to customize its behavior and authentication methods.

### Basic Configuration

```Go
import "github.com/zarbanio/zarban-go/service"

// Basic configuration with just the host URL
client, err := service.NewClient("https://testapi.zarban.io")
if err != nil {
    log.Fatalf("Failed to create service client: %v", err)
    return
}
```

### Authentication Options

The SDK supports multiple authentication methods:

1. API Key Authentication:

```Go
// Define headers to be added
headers := map[string]string{
    "Authorization": "Bearer " + token,
}

// configure it with the header editing function
client, err = service.NewClient(
    "https://testapi.zarban.io",
    service.WithRequestEditorFn(service.AddHeaders(headers)),
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
httpResponse, err = client.GetAllIlks(context.Background())
if err != nil {
    log.Fatalf("Error during API call -> GetAllIlks: %v", err)
    return
}

var ilksResponse service.IlksResponse
err = service.HandleAPIResponse(context.Background(), httpResponse, &ilksResponse)
if err != nil {
    if apiErr, ok := err.(*service.APIError); ok {
        fmt.Println(service.PrettyPrintError(apiErr))
    } else {
        log.Printf("Unexpected error: %v", err)
    }
    return
}
```

Manual Error Handling

```go
httpResponse, err = client.GetAllIlks(context.Background())
if err != nil {
    log.Fatalf("Error during API call: %v", err)
    return
}

ilksResponse, err := service.ParseGetAllIlksResponse(httpResponse)
if err != nil {
    log.Fatalf("Error while parsing http response: %v", err)
    return
}

switch i.StatusCode() {
case 200:
    fmt.Printf("Ilks fetched successfully. Count: %d\n", len(i.JSON200.Data))
    return i.JSON200, nil
case 400:
    return i.JSON400, fmt.Errorf("bad request: %s", i.JSON400.Msg)
case 500:
    return i.JSON500, fmt.Errorf("internal server error: %s", i.JSON500.Msg)
default:
    return nil, fmt.Errorf("unexpected status code: %d", i.StatusCode())
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

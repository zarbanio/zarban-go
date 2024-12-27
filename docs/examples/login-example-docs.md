# User Login Example

This example demonstrates how to implement user authentication using the Zarban SDK. It shows how to log in a user, handle the authentication token.

## Prerequisites

Before running this example, ensure you have:

1. Installed the Zarban SDK:

```bash
go get github.com/zarbanio/zarban-go
```

2. Access to the Zarban API (test environment)

## API Specification

### Endpoint: `/auth/login`

- **Method**: POST
- **Description**: Login with email and password and get a JWT token.

### Request Format

```json
{
  "email": "example@domain.com",
  "password": "password"
}
```

#### Required Fields

| Field    | Type   | Description     | Example            |
| -------- | ------ | --------------- | ------------------ |
| email    | string | User's email    | example@domain.com |
| password | string | User's password | password           |

### Response Format

#### Success Response (200 OK)

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR..."
}
```

#### Error Responses

- **400 Bad Request**

```json
{
  "msg": "Bad request",
  "reasons": ["Invalid address"]
}
```

- **401 Unauthorized**

```json
{
  "msg": "Unauthorized",
  "reasons": ["Invalid credentials"]
}
```

- **404 Not Found**

```json
{
  "msg": "Not Found",
  "reasons": ["User not found"]
}
```

- **500 Internal Server Error**

```json
{
  "msg": "Internal Server Error",
  "reasons": ["Server error occurred"]
}
```

## Code Example

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zarbanio/zarban-go/wallet"
)

func LoginExample() {
	// Create and configure the client
	client, err := wallet.NewClient("https://testwapi.zarban.io")
	if err != nil {
		log.Fatalf("Failed to create wallet client: %v", err)
		return
	}
	// Prepare the signup request data
	loginRequest := wallet.LoginRequest{
		Email:    "user@example.com",
		Password: "your_secure_password",
	}

	httpResponse, err := client.LoginWithEmailAndPassword(context.Background(), loginRequest)
	if err != nil {
		log.Fatalf("Error during API call: %v", err)
		return
	}

	var successResponse wallet.JwtResponse
	err = wallet.HandleAPIResponse(context.Background(), httpResponse, &successResponse)
	if err != nil {
		if apiErr, ok := err.(*wallet.APIError); ok {
			fmt.Println(wallet.PrettyPrintError(apiErr))
		} else {
			log.Printf("Unexpected error: %v", err)
		}
		return
	}

	fmt.Println("login successful!")
	fmt.Println("Token: ", successResponse.Token)
}

func main() {
    LoginExample()
}
```

## Step-by-Step Explanation

1. **Configure API Client**

   ```go
    client, err := wallet.NewClient("https://testwapi.zarban.io")
	if err != nil {
		log.Fatalf("Failed to create wallet client: %v", err)
		return
	}
   ```

   Creates and configures the API client with the test environment endpoint.

2. **Prepare Login Request**

   ```go
    loginRequest := wallet.LoginRequest{
        Email:    "user@example.com",
        Password: "your_secure_password",
    }
   ```

   Creates a login request object with user credentials.

3. **Handle Authentication**

   ```go
   httpResponse, err := client.LoginWithEmailAndPassword(context.Background(), loginRequest)
   if err != nil {
        log.Fatalf("Error during API call: %v", err)
        return
    }

	var successResponse wallet.JwtResponse
	err = wallet.HandleAPIResponse(context.Background(), httpResponse, &successResponse)
	if err != nil {
		if apiErr, ok := err.(*wallet.APIError); ok {
			fmt.Println(wallet.PrettyPrintError(apiErr))
		} else {
			log.Printf("Unexpected error: %v", err)
		}
		return
	}
   ```

   Sends login request and you can store the authentication token for future requests.

## Error Handling

The example includes comprehensive error handling based on the API specification:

### Error Status Codes

1. **400 Bad Request**

   - Invalid email format
   - Missing required fields
   - Password format validation failure

2. **401 Unauthorized**

   - Invalid credentials
   - Account locked
   - Too many failed attempts

3. **404 Not Found**

   - User account doesn't exist
   - Account deleted or deactivated

4. **500 Internal Server Error**
   - Database connection issues
   - Server configuration problems
   - Internal service failures

### Error Handling Example

Error Handling is already done by HandleAPIResponse method but you can do more with error if you want

```go
if err != nil {
    if apiErr, ok := err.(*wallet.APIError); ok {
        fmt.Println(wallet.PrettyPrintError(apiErr))
    } else {
        log.Printf("Unexpected error: %v", err)
    }
    return
}
```

## Best Practices

1. **Credential Management**

   ```go
   // DON'T store credentials in code
    const email = "user@example.com"  // Incorrect

    // DO use environment variables
    import (
        "os"
        "github.com/joho/godotenv"
    )

    func main() {
        // Load .env file
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Error loading .env file")
        }
        
        // Get environment variable
        email := os.Getenv("ZARBAN_USER_EMAIL")
    }
   ```

2. **Token Storage**

   ```go
   func saveToken(token string) {
    // You could store in secure storage like:
    // - Encrypted file
    // - System keychain
    // - Secure database
    }

    func logout() {
        // Clear the stored token
        // Set token to empty string or nil
        // Clear any session data
    }
   ```

3. **Security Considerations**
   - Always use HTTPS for API calls
   - Implement rate limiting
   - Monitor for suspicious login attempts
   - Store tokens securely
   - Implement token refresh mechanisms
   - Clear tokens on logout
   - Handle token expiration

## See Also

- [API Reference Documentation](../wallet)
- [Security Best Practices](security-best-practices.md)

# Child User Creation and Management

This example demonstrates how to create and manage child users using the Zarban SDK. It includes the process of authenticating as a superuser, creating a child user, and accessing the child user's profile.

## Prerequisites

Before running this example, ensure you have:

1. Installed the Zarban SDK:

```bash
go get github.com/zarbanio/zarban-go
```

2. Superuser credentials with appropriate permissions
3. Access to the Zarban API (test environment)

## API Specification

### Endpoint: `/users/children`

- **Method**: POST
- **Description**: Create a child user
- **Authentication**: Required (Bearer Token)

### Request Format

```json
{
  "username": "john"
}
```

#### Required Fields

| Field    | Type   | Description         | Example |
| -------- | ------ | ------------------- | ------- |
| username | string | Child user username | "john"  |

### Response Format

#### Success Response (200 OK)

```json
{
  "username": "john"
  // Additional user properties
}
```

#### Error Responses

- **400 Bad Request**

```json
{
  "msg": "Bad request",
  "reasons": ["Invalid username format"]
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

func CreateChildExample() {
	// Create and configure the client
	client, err := wallet.NewClient("https://testwapi.zarban.io")
	if err != nil {
		log.Fatalf("Failed to create wallet client: %v", err)
		return
	}

	// Constant superuser email and password
	const SUPERUSER_EMAIL = "user@example.com"
	const SUPERUSER_PASSWORD = "your_secure_password"
	// Prepare the signup request data
	loginRequest := wallet.LoginRequest{
		Email:    SUPERUSER_EMAIL,
		Password: SUPERUSER_PASSWORD,
	}

	// Call the login API
	httpResponse, err := client.LoginWithEmailAndPassword(context.Background(), loginRequest)
	if err != nil {
		log.Fatalf("Error during API call: %v", err)
		return
	}

	var loginResponse wallet.JwtResponse

	err = wallet.HandleAPIResponse(context.Background(), httpResponse, &loginResponse)
	if err != nil {
		if apiErr, ok := err.(*wallet.APIError); ok {
			fmt.Println(wallet.PrettyPrintError(apiErr))
		} else {
			log.Printf("Unexpected error: %v", err)
		}
		return
	}

	fmt.Println("Superuser login successful")
	fmt.Println("Token: ", loginResponse.Token)

	// Define headers to be added
	headers := map[string]string{
		"Authorization": "Bearer " + loginResponse.Token,
	}

	// re-configure it with the header editing function
	client, err = wallet.NewClient(
		"https://testwapi.zarban.io",
		wallet.WithRequestEditorFn(wallet.AddHeaders(headers)),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Create a child user
	const childUsername = "child_user_test"
	createChildUserRequest := wallet.CreateChildUserRequest{
		Username: childUsername,
	}

	// Call the login API
	httpResponse, err = client.CreateChildUser(context.Background(), createChildUserRequest)
	if err != nil {
		log.Fatalf("Error during API call: %v", err)
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

	fmt.Println("Child user created. Username: ", *createChildResponse.Username)

	headers = map[string]string{
		"Authorization": "Bearer " + loginResponse.Token,
		"X-Child-User":  *createChildResponse.Username,
	}
	// re-configure it with the header editing function
	client, err = wallet.NewClient(
		"https://testwapi.zarban.io",
		wallet.WithRequestEditorFn(wallet.AddHeaders(headers)),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Make the profile request
	httpResponse, err = client.GetUserProfile(context.Background())
	if err != nil {
		log.Fatalf("Error during API call: %v", err)
		return
	}

	var getUserProfileResponse wallet.User
	err = wallet.HandleAPIResponse(context.Background(), httpResponse, &getUserProfileResponse)
	if err != nil {
		if apiErr, ok := err.(*wallet.APIError); ok {
			fmt.Println(wallet.PrettyPrintError(apiErr))
		} else {
			log.Printf("Unexpected error: %v", err)
		}
		return
	}

	fmt.Println("Child user profile:")
	fmt.Println(getUserProfileResponse)
}

func main () {
    CreateChildExample()
}
```

## Step-by-Step Explanation

1. **Initialize API Client**

   ```go
    client, err := wallet.NewClient("https://testwapi.zarban.io")
	if err != nil {
		log.Fatalf("Failed to create wallet client: %v", err)
		return
	}
   ```

   Sets up the API client with the test environment endpoint.

2. **Superuser Authentication**

   ```go
    // Constant superuser email and password
	const SUPERUSER_EMAIL = "user@example.com"
	const SUPERUSER_PASSWORD = "your_secure_password"
	// Prepare the signup request data
	loginRequest := wallet.LoginRequest{
		Email:    SUPERUSER_EMAIL,
		Password: SUPERUSER_PASSWORD,
	}

	// Call the login API
	httpResponse, err := client.LoginWithEmailAndPassword(context.Background(), loginRequest)
	if err != nil {
		log.Fatalf("Error during API call: %v", err)
		return
	}

	var loginResponse wallet.JwtResponse

	err = wallet.HandleAPIResponse(context.Background(), httpResponse, &loginResponse)
	if err != nil {
		if apiErr, ok := err.(*wallet.APIError); ok {
			fmt.Println(wallet.PrettyPrintError(apiErr))
		} else {
			log.Printf("Unexpected error: %v", err)
		}
		return
	}

	fmt.Println("Superuser login successful")
	fmt.Println("Token: ", loginResponse.Token)

	// Define headers to be added
	headers := map[string]string{
		"Authorization": "Bearer " + loginResponse.Token,
	}

	// re-configure it with the header editing function
	client, err = wallet.NewClient(
		"https://testwapi.zarban.io",
		wallet.WithRequestEditorFn(wallet.AddHeaders(headers)),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
   ```

Authenticates the superuser and stores the access token.

3. **Create Child User**

    ```go
    // Create a child user
	const childUsername = "child_user_test"
	createChildUserRequest := wallet.CreateChildUserRequest{
		Username: childUsername,
	}

	// Call the login API
	httpResponse, err = client.CreateChildUser(context.Background(), createChildUserRequest)
	if err != nil {
		log.Fatalf("Error during API call: %v", err)
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

	fmt.Println("Child user created. Username: ", *createChildResponse.Username)
    ```

Creates a new child user account.

4. **Access Child User Profile**

   ```go
   headers = map[string]string{
		"Authorization": "Bearer " + loginResponse.Token,
		"X-Child-User":  *createChildResponse.Username,
	}
	// re-configure it with the header editing function
	client, err = wallet.NewClient(
		"https://testwapi.zarban.io",
		wallet.WithRequestEditorFn(wallet.AddHeaders(headers)),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Make the profile request
	httpResponse, err = client.GetUserProfile(context.Background())
	if err != nil {
		log.Fatalf("Error during API call: %v", err)
		return
	}

	var getUserProfileResponse wallet.User
	err = wallet.HandleAPIResponse(context.Background(), httpResponse, &getUserProfileResponse)
	if err != nil {
		if apiErr, ok := err.(*wallet.APIError); ok {
			fmt.Println(wallet.PrettyPrintError(apiErr))
		} else {
			log.Printf("Unexpected error: %v", err)
		}
		return
	}

	fmt.Println("Child user profile:")
	fmt.Println(getUserProfileResponse)
   ```

   Sets the required header and retrieves the child user's profile.

## Important Headers

| Header Name   | Description                     | Example Value    |
| ------------- | ------------------------------- | ---------------- |
| Authorization | Bearer token for authentication | Bearer eyJhbG... |
| X-Child-User  | Username of the child user      | child_user_test  |

## Error Handling

### Common Error Scenarios

1. **400 Bad Request**

   - Invalid username format
   - Username already exists
   - Missing required fields

2. **500 Internal Server Error**
   - Database connection issues
   - Internal service failures

### Error Handling Example

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

1. **Header Management**

   ```go
   // DO clean up headers after use
    client, err = wallet.NewClient(
		"https://testwapi.zarban.io",
	)

   // DON'T leave sensitive headers in place
   // Incorrect: Not removing headers after use
   ```

2. **Username Generation**

   ```go
    // DO use safe username generation
    import (
        "fmt"
        "github.com/google/uuid"
        "strings"
    )

    // DO use random UUIDs for usernames
    func generateSafeUsername() string {
        // Generate UUID and remove hyphens
        id := strings.Replace(uuid.New().String(), "-", "", -1)
        // Take first 8 characters and create username
        safeUsername := fmt.Sprintf("child_%s", id[:8])
        return safeUsername
    }

    // DON'T use predictable usernames
    // Incorrect: sequential usernames
    // func generateUsername(seq int) string {
    //     return fmt.Sprintf("child_%d", seq)  // Avoid using sequential numbers
    // }
   ```

3. **Security Considerations**
   - Implement proper authentication checks
   - Use secure random username generation
   - Clean up sensitive headers after use
   - Implement proper error handling
   - Monitor child user creation activities
   - Implement rate limiting for creation requests

## See Also

- [API Reference Documentation](../wallet)
- [Security Best Practices](security-best-practices.md)

# User Signup Example

This example demonstrates how to use the Zarban SDK to implement a user signup process. The example shows how to create a new user account by sending a signup request to the Zarban API.

## Prerequisites

Before running this example, make sure you have:

1. Installed the Zarban SDK:

```bash
go get github.com/zarbanio/zarban-go
```

2. Access to the Zarban API (test environment)

## Code Example

```go
package main

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

func main(){
    SignupExample()
}
```

## Step-by-Step Explanation

1. **Import Required Modules**

```go
import (
    "context"
    "fmt"
    "log"

    "github.com/zarbanio/zarban-go/wallet"
)
   ```

   These imports provide the necessary classes and functions to interact with the Zarban API.

2. **Configure the API Client**

   ```go
    client, err := wallet.NewClient("https://testwapi.zarban.io")
	if err != nil {
		log.Fatalf("Failed to create wallet client: %v", err)
		return
	}
   ```

   Creates a client object with the API endpoint

3. **Prepare Signup Request**

   ```go
    signUpRequest := wallet.SignUpRequest{
		Email:    "user@example.com",
		Password: "yourSecurePassword",
	}
   ```

   Creates a signup request object with user credentials.

4. **Make the API Call**
   ```go
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
   ```
   Sends the signup request to the API and handles the response.

## Response Handling

The example includes error handling using HandleAPIResponse:

- On success: Prints a confirmation message and the response messages
- On failure: Catches `error` and prints the error details

## Expected Output

On successful signup:

```
Confirmation link sent successful!
Message: [Confirmation email details...]
```

On error:

```
Exception when calling DefaultApi->auth_signup_post: [Error details]
Error message: [Detailed error message]
```

## Important Notes

1. Replace `"user@example.com"` and `"yourSecuredPassword"` with actual user credentials
2. The example uses the test API endpoint (`testwapi.zarban.io`). For production use, update the host accordingly
3. Ensure proper password security practices when implementing in production
4. The API will send a confirmation email to the provided email address

## Error Handling

Common errors that might occur:

- Invalid email format
- Password doesn't meet security requirements
- Email already registered
- Network connectivity issues
- API server errors

## See Also

- [API Reference Documentation](../wallet)
- [Security Best Practices](security-best-practices.md)

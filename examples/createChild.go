package examples

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

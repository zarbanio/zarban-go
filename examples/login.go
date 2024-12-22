package examples

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
	err = HandleAPIResponse(httpResponse, &successResponse)
	if err != nil {
		if apiErr, ok := err.(*APIError); ok {
			fmt.Println(PrettyPrintError(apiErr))
		} else {
			log.Printf("Unexpected error: %v", err)
		}
		return
	}

	fmt.Println("login successful!")
	fmt.Println("Token: ", successResponse.Token)
}

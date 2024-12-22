package examples

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
	err = HandleAPIResponse(httpResponse, &successResponse)
	if err != nil {
		if apiErr, ok := err.(*APIError); ok {
			fmt.Println(PrettyPrintError(apiErr))
		} else {
			log.Printf("Unexpected error: %v", err)
		}
		return
	}

	fmt.Printf("Signup successful: %+v\n", successResponse.Messages)
}

package examples

import (
	"context"
	"fmt"
	"log"

	"github.com/zarbanio/zarban-go/wallet"
)

func createLoan(
	client wallet.Client,
	planName string,
	collateral string,
	debt string,
	symbol string,
	loanToValueOption wallet.LoanToValueOptions,
) (wallet.LoansResponse, error) {
	loanCreateRequest := wallet.LoanCreateRequest{
		Intent:            wallet.LoanCreateRequestIntentCreate,
		PlanName:          planName,
		Collateral:        &collateral,
		Debt:              &debt,
		Symbol:            symbol,
		LoanToValueOption: loanToValueOption,
	}
	// Call the CreateLoanVault API
	httpResponse, err := client.CreateLoanVault(context.Background(), loanCreateRequest)
	if err != nil {
		log.Fatalf("Error during API call -> CreateLoanVault: %v", err)
		return wallet.LoansResponse{}, err
	}

	var loansResponse wallet.LoansResponse
	err = wallet.HandleAPIResponse(context.Background(), httpResponse, &loansResponse)
	if err != nil {
		if apiErr, ok := err.(*wallet.APIError); ok {
			fmt.Println(wallet.PrettyPrintError(apiErr))
		} else {
			log.Printf("Unexpected error: %v", err)
		}
		return wallet.LoansResponse{}, err
	}

	fmt.Println("Loan created successfully. Loan ID: ", loansResponse.Id)

	return loansResponse, nil
}

func loanStatus(
	client wallet.Client,
	loanId string,
) (wallet.LoansResponse, error) {
	// Call the GetLoanDetails API
	httpResponse, err := client.GetLoanDetails(context.Background(), loanId)
	if err != nil {
		log.Fatalf("Error during API call -> CreateLoanVault: %v", err)
		return wallet.LoansResponse{}, err
	}

	var loansResponse wallet.LoansResponse
	err = wallet.HandleAPIResponse(context.Background(), httpResponse, &loansResponse)
	if err != nil {
		if apiErr, ok := err.(*wallet.APIError); ok {
			fmt.Println(wallet.PrettyPrintError(apiErr))
		} else {
			log.Printf("Unexpected error: %v", err)
		}
		return wallet.LoansResponse{}, err
	}

	fmt.Println("Loan Details for ID: ", loanId)
	fmt.Println("State: ", loansResponse.State)
	fmt.Println("Collateral: ", loansResponse.Collateral)
	fmt.Println("Debt: ", loansResponse.Debt)
	fmt.Println("Liquidation Price: ", loansResponse.LiquidationPrice.Values)
	fmt.Println("Loan To Value: ", loansResponse.LoanToValue)

	return loansResponse, nil
}

func CreateLoanExample() {
	// Replace with your actual access token
	const ACCESS_TOKEN = "your_access_token_here"

	// Setup API client
	headers := map[string]string{
		"Authorization": "Bearer " + ACCESS_TOKEN,
		"X-Child-User":  "child_user_test",
	}
	client, err := wallet.NewClient(
		"https://testwapi.zarban.io",
		wallet.WithRequestEditorFn(wallet.AddHeaders(headers)),
	)
	if err != nil {
		log.Fatalf("Failed to create wallet client: %v", err)
		return
	}

	// Loan creation parameters, Replace them with yout actual data
	// *** either collateral or debt must be empty ***
	const PLAN_NAME = "DAIA"             // Only DAIA and DAIB supported
	const COLLATERAL = "1000"            // Amount of collateral
	const DEBT = ""                      // Amount of debt
	const SYMBOL = "DAI"                 // Coin symbol
	const LOAN_TO_VALUE_OPTIONS = "Safe" // Risky - Normal - Safe

	createLoanResponse, err := createLoan(*client, PLAN_NAME, COLLATERAL, DEBT, SYMBOL, LOAN_TO_VALUE_OPTIONS)
	if err != nil {
		return
	}

	// Remove the X-Child-User header after use
	headers = map[string]string{
		"Authorization": "Bearer " + ACCESS_TOKEN,
	}
	client, err = wallet.NewClient(
		"https://testwapi.zarban.io",
		wallet.WithRequestEditorFn(wallet.AddHeaders(headers)),
	)
	if err != nil {
		log.Fatalf("Failed to create wallet client: %v", err)
		return
	}

	// Track loan status
	fmt.Println("\nTracking loan status...")
	loanDetails, err := loanStatus(
		*client,
		createLoanResponse.Id,
	)
	if err != nil {
		// you can do some addition works with error here!
		return
	}
	// You can add more specific actions based on the loan state
	fmt.Println("Loan status: ", loanDetails.State)
}

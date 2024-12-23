package examples

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/zarbanio/zarban-go/wallet"
)

func repayLoan(
	client wallet.Client,
	loanId string,
	intent wallet.RepayLoanRequestIntent,
) (wallet.LoansResponse, error) {
	repayLoanRequest := wallet.RepayLoanRequest{
		LoanId: loanId,
		Intent: intent,
	}
	// Call the CreateLoanVault API
	httpResponse, err := client.RepayLoan(context.Background(), repayLoanRequest)
	if err != nil {
		log.Fatalf("Error during API call -> RepayLoan: %v", err)
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

	fmt.Printf("Loan repayment successful. Loan ID: %s", loansResponse.Id)

	return loansResponse, nil
}

func getLoanStatus(
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
	fmt.Println("Collateralization Ratio: ", loansResponse.CollateralizationRatio)
	fmt.Println("Loan To Value: ", loansResponse.LoanToValue)
	fmt.Println("Debt: ", loansResponse.Debt)
	fmt.Println("Liquidation Price: ", loansResponse.LiquidationPrice.Values)
	fmt.Println("Loan Plan: ", loansResponse.Plan)

	return loansResponse, nil
}

func RepayLoanExample() {
	// Replace with your actual access token
	const ACCESS_TOKEN = "your_access_token_here"

	// Setup API client
	// Set the X-Child-User header
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

	// Loan ID to repay, replace with actual loan ID
	const LOAN_ID = "DAIA#2825"

	// Preview repayment
	fmt.Println("Previewing loan repayment...")

	previewLoanResponse, err := repayLoan(*client, LOAN_ID, wallet.RepayLoanRequestIntentPreview)
	if err != nil {
		// you can do some addition works with error here!
		return
	}

	fmt.Println("\nRepayment preview details:")
	fmt.Println("Collateral to be returned: ,", previewLoanResponse.Collateral)
	fmt.Println("Debt to be repaid: ", previewLoanResponse.Debt)

	var confirm string
	fmt.Print("\nDo you want to proceed with the repayment? (y/n): ")
	fmt.Scanln(&confirm)

	if confirm == "y" {
		// Proceed with actual repayment
		repaymentLoanResponse, err := repayLoan(*client, LOAN_ID, wallet.RepayLoanRequestIntentRepay)
		if err != nil {
			// you can do some addition works with error here!
			return
		}
		fmt.Print("repayment in progress...")
		for {
			loanDetails, err := getLoanStatus(
				*client,
				repaymentLoanResponse.Id,
			)
			if err != nil {
				// you can do some addition works with error here!
				return
			}
			if loanDetails.State["LocaleEn"] == "Loan settled" {
				fmt.Print("\nLoan repayment successful!")
				getLoanStatus(*client, repaymentLoanResponse.Id)
				break
			} else if loanDetails.State["LocaleEn"] == "Loan settlement failed" {
				fmt.Print(loanDetails.State["LocaleEn"])
				break
			}
			time.Sleep(1 * time.Second)
		}
	} else {
		fmt.Print("Repayment cancelled.")
	}
}

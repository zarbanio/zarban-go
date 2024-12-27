/*
Vault Creation Script using Web3 and Zarban API

This script demonstrates the process of creating a vault in a stablecoin system
using the Web3 library for Ethereum interactions and the Zarban API for
stablecoin-specific operations.

Key components and functionality:

1. Imports and Setup:
- Web3 for Ethereum blockchain interactions
- Zarban API client for stablecoin system operations
- Custom functions for data conversion and API interactions

2. Helper Functions:
- get_ilks_symbol: Retrieves available collateral types (ilks) from the API
- to_native: Converts human-readable amounts to blockchain-native amounts
- get_vault_tx_steps: Obtains transaction steps for vault creation

3. Main Execution:
- Sets up Web3 connection and Zarban API client
- Defines vault creation parameters (collateral type, amounts)
- Retrieves vault creation transaction steps
- Iterates through steps, creating and sending Ethereum transactions

Usage:
1. Replace placeholder values (RPC URL, private key, wallet address)
2. Set desired ILK_NAME, COLLATERAL_AMOUNT, and LOAN_AMOUNT
3. Run the script to create a vault with specified parameters

Note: This script interacts with real blockchain networks and APIs. Use with
caution and ensure you understand the implications of each transaction.

Security Warning: Never hardcode or commit private keys. Use secure methods
for managing sensitive information in production environments.
*/
package examples

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/zarban-go/service"
)

func getIlksSymbol(client service.Client) ([]service.Symbol, error) {
	httpResponse, err := client.GetAllIlks(context.Background())
	if err != nil {
		fmt.Printf("Error during API call -> GetAllIkls: %v", err)
	}
	var ilks service.IlksResponse
	err = service.HandleAPIResponse(context.Background(), httpResponse, &ilks)
	if err != nil {
		if apiErr, ok := err.(*service.APIError); ok {
			fmt.Println(service.PrettyPrintError(apiErr))
		} else {
			log.Printf("Unexpected error: %v", err)
		}
		return nil, err
	}

	// Use a map to ensure unique symbols
	symbolSet := make(map[service.Symbol]struct{})
	for _, ilk := range ilks.Data {
		symbolSet[ilk.Symbol] = struct{}{}
	}

	// Convert the map keys back to a slice
	symbols := make([]service.Symbol, 0, len(symbolSet))
	for symbol := range symbolSet {
		symbols = append(symbols, symbol)
	}

	return symbols, nil
}

func toNative(client service.Client, symbol service.Symbol, amount float64) (string, error) {
	// Get symbols
	symbols, err := getIlksSymbol(client)
	if err != nil {
		return "", fmt.Errorf("failed to get symbols: %w", err)
	}

	// Define precision map
	precision := map[service.Symbol]int{
		"ZAR": 18,
	}
	for _, s := range symbols {
		precision[s] = 18
	}

	// Validate symbol
	p, ok := precision[symbol]
	if !ok {
		return "", fmt.Errorf("unknown symbol: %s", symbol)
	}

	// Convert amount to native units using big.Int
	scaledAmount := new(big.Int)
	floatAmount := new(big.Float).SetFloat64(amount)
	scale := new(big.Float).SetFloat64(math.Pow10(p))

	// Multiply amount by scale factor
	floatAmount.Mul(floatAmount, scale)

	// Convert to integer, checking for accuracy
	if _, accuracy := floatAmount.Int(scaledAmount); accuracy != big.Exact {
		return "", fmt.Errorf("lost precision during conversion")
	}

	return scaledAmount.String(), nil
}

func getVaultTxSteps(
	client service.Client,
	ilkName string,
	symbol service.Symbol,
	walletAddress string,
	collateralAmount float64,
	loanAmount float64,
) (service.ChainActivity, error) {
	nativeCollateralAmount, err := toNative(client, symbol, collateralAmount)
	if err != nil {
		log.Fatalf("Error converting collateral amount: %v", err)
		return service.ChainActivity{}, err
	}

	nativeLoanAmount, err := toNative(client, "ZAR", loanAmount)
	if err != nil {
		log.Fatalf("Error converting loan amount: %v", err)
		return service.ChainActivity{}, err
	}

	request := service.StablecoinSystemCreateVaultTxRequest{
		CollateralAmount: &nativeCollateralAmount,
		MintAmount:       nativeLoanAmount,
		User:             walletAddress,
		IlkName:          ilkName,
	}

	httpResponse, err := client.CreateStableCoinVault(context.Background(), request)
	if err != nil {
		log.Fatalf("Error during API call -> CreateStableCoinVault: %v", err)
		return service.ChainActivity{}, err
	}

	var txSteps service.ChainActivity
	err = service.HandleAPIResponse(context.Background(), httpResponse, &txSteps)
	if err != nil {
		if apiErr, ok := err.(*service.APIError); ok {
			log.Printf("API Error: %v", service.PrettyPrintError(apiErr))
		} else {
			log.Printf("Unexpected error: %v", err)
		}
		return service.ChainActivity{}, err
	}

	return txSteps, nil

}

func GetAddressFromPrivateKey(privateKeyHex string) (string, error) {
	// Decode the private key from hex
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode private key: %v", err)
	}

	// Derive the public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("invalid public key type")
	}

	// Get the address
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return address, nil
}

func getLogs(client service.Client, txHash string) ([]service.Log, error) {
	httpResponse, err := client.GetLogsByTransactionHash(context.Background(), txHash)
	if err != nil {
		log.Fatalf("Error during API call -> GetLogsByTransactionHash: %v", err)
		return nil, err
	}
	var logs service.EventDetailsResponse
	err = service.HandleAPIResponse(context.Background(), httpResponse, &logs)
	if err != nil {
		if apiErr, ok := err.(*service.APIError); ok {
			log.Printf("API Error: %v", service.PrettyPrintError(apiErr))
		} else {
			log.Printf("Unexpected error: %v", err)
		}
		return nil, err
	}

	return logs.Data, nil
}

func getVaultId(logs []service.Log) (int, error) {
	for _, log := range logs {
		if log.Contract == "Cdpmanager" {
			payload := *log.Decoded
			vaultIdStr, ok := payload["Cdp"]
			if !ok {
				return 0, fmt.Errorf("failed to get vault id from log")
			}
			vaultId, err := strconv.Atoi(vaultIdStr)
			if err != nil {
				return 0, fmt.Errorf("failed to convert vault id to int")
			}
			return vaultId, nil
		}
	}
	return 0, nil
}

type EthereumTransaction struct {
	From     string   `json:"from"`     // Sender address
	To       string   `json:"to"`       // Recipient address
	Value    int      `json:"value"`    // Amount of ETH to send (in Wei)
	Gas      int      `json:"gas"`      // Gas limit for the transaction
	GasPrice *big.Int `json:"gasPrice"` // Gas price (in Wei)
	Nonce    uint64   `json:"nonce"`    // Nonce (transaction count)
	ChainID  *big.Int `json:"chainId"`  // Chain ID (mainnet: 1)
	Data     string   `json:"data"`     // Optional data for the transaction
}

type TransactionDetails struct {
	Timestamp string              `json:"timestamp"`
	Tx        EthereumTransaction `json:"tx"`
	TxHash    string              `json:"tx_hash"`
	VaultID   int                 `json:"vault_id"`
}

func SaveTransactionDetails(tx EthereumTransaction, txHash string, vaultID int) error {
	// Create transaction details
	transactionData := TransactionDetails{
		Timestamp: time.Now().Format(time.RFC3339),
		Tx:        tx,
		TxHash:    txHash,
		VaultID:   vaultID,
	}

	// Read existing data
	var existingData []TransactionDetails
	file, err := os.Open("transaction_log.json")
	if err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&existingData); err != nil {
			return fmt.Errorf("failed to decode existing data: %v", err)
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("failed to open transaction_log.json: %v", err)
	}

	// Append new transaction data
	existingData = append(existingData, transactionData)

	// Write updated data back to the file
	file, err = os.Create("transaction_log.json")
	if err != nil {
		return fmt.Errorf("failed to create transaction_log.json: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(existingData); err != nil {
		return fmt.Errorf("failed to write transaction data: %v", err)
	}

	fmt.Println("Transaction details saved to transaction_log.json")
	return nil
}

func waitForTransactionReceipt(client *ethclient.Client, txHash common.Hash, maxWaitTime time.Duration, checkInterval time.Duration) (*types.Receipt, error) {
	startTime := time.Now()

	for time.Since(startTime) < maxWaitTime {
		// Attempt to get the transaction receipt
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			// Receipt found
			return receipt, nil
		}

		// Handle specific error types
		if err != ethereum.NotFound {
			log.Printf("Error checking transaction receipt: %v", err)
		}

		fmt.Printf("Waiting for transaction %s to be mined...\n", txHash.Hex())
		time.Sleep(checkInterval)
	}

	fmt.Printf("Transaction not mined after %v seconds\n", maxWaitTime.Seconds())
	return nil, fmt.Errorf("transaction not mined after %v seconds", maxWaitTime.Seconds())
}

func CreateVaultExample() {
	// Configuration
	const HTTPS_RPC_URL = "Replace with your Ethereum node URL"
	const PRIVATE_KEY = "Replace with your Private key"
	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}
	WALLET_ADDRESS, err := GetAddressFromPrivateKey(PRIVATE_KEY)
	if err != nil {
		log.Fatalf("Failed to get address from private key: %v", err)
		return
	}

	// Create and configure the client
	client, err := service.NewClient("https://testapi.zarban.io")
	if err != nil {
		log.Fatalf("Failed to create wallet client: %v", err)
		return
	}

	// Connect to an Ethereum client
	ethClient, err := ethclient.Dial(HTTPS_RPC_URL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
		return
	}

	account := WALLET_ADDRESS

	// Get the chain ID
	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
		return
	}

	// Define vault creation parameters
	const ILK_NAME = "ETHA"        // Replace with your desired ilk
	const SYMBOL = "ETH"           // Replace with the symbol associated with your ilk
	const COLLATERAL_AMOUNT = .001 // Replace with your desired amount
	const LOAN_AMOUNT = 100        // Replace with your desired amount

	vaultSteps, err := getVaultTxSteps(
		*client,
		ILK_NAME,
		SYMBOL,
		WALLET_ADDRESS,
		COLLATERAL_AMOUNT,
		LOAN_AMOUNT)
	if err != nil {
		log.Fatalf("Failed to get vault creartion steps: %v", err)
		return
	}

	numOfSteps := vaultSteps.NumberOfSteps
	stepNumber := vaultSteps.StepNumber
	steps := vaultSteps.Steps

	if len(steps) > 0 {
		for s := 0; s <= (numOfSteps - stepNumber); s++ {
			vaultSteps, err := getVaultTxSteps(
				*client,
				ILK_NAME,
				SYMBOL,
				WALLET_ADDRESS,
				COLLATERAL_AMOUNT,
				LOAN_AMOUNT)
			if err != nil {
				log.Fatalf("Failed to get vault creartion steps: %v", err)
				return
			}
			numberOf := vaultSteps.NumberOfSteps
			stepNumber := vaultSteps.StepNumber
			steps := vaultSteps.Steps

			var txHash string
			for number, step := range steps {
				// Do something with number and step
				data, err := step.Data.AsPreparedTx()
				if err != nil {
					log.Fatalf("Failed to convert step to prepared tx: %v", err)
					return
				}
				label := data.Label
				fmt.Printf("steps %d: %s\n", number+1, label["en-US"])
				if (stepNumber - 1) == number {
					fmt.Println("processing...")
					methodParams := data.MethodParameters
					addressTo := methodParams.To
					calldata := methodParams.Calldata
					valueStr := methodParams.Value
					value, err := strconv.Atoi(valueStr)
					if err != nil {
						log.Fatalf("failed to convert vault id to int")
						return
					}
					// Get the transaction count (nonce)
					nonce, err := ethClient.PendingNonceAt(context.Background(), common.HexToAddress(account))
					if err != nil {
						log.Fatalf("Failed to get transaction count: %v", err)
						return
					}
					gas := data.GasUseEstimate
					// Get the current gas price
					gasPrice, err := ethClient.SuggestGasPrice(context.Background())
					if err != nil {
						log.Fatalf("Failed to get gas price: %v", err)
						return
					}

					// Prepare transaction
					tx := types.NewTransaction(nonce, common.HexToAddress(addressTo), big.NewInt(int64(value)), uint64(gas), gasPrice, []byte(calldata))

					txToSave := EthereumTransaction{
						From:     WALLET_ADDRESS,
						To:       addressTo,
						Value:    value,
						Gas:      gas,
						GasPrice: gasPrice,
						Nonce:    nonce,
						ChainID:  chainID,
						Data:     methodParams.Calldata,
					}
					// Sign the transaction
					signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
					if err != nil {
						log.Fatalf("Failed to sign the transaction: %v", err)
						return
					}

					// Send the signed transaction
					err = ethClient.SendTransaction(context.Background(), signedTx)
					if err != nil {
						log.Fatalf("Failed to send transaction: %v", err)
						return
					}

					// Get the transaction hash
					txHash = signedTx.Hash().Hex()
					fmt.Printf("Transaction sent: %s\n", txHash)

					// Save transaction details (function needs to be implemented)
					err = SaveTransactionDetails(txToSave, txHash, -1) // vault_id is nil at this point
					if err != nil {
						log.Fatalf("Failed to save transaction details: %v", err)
						return
					}
					// Wait for the transaction to be mined
					receipt, err := waitForTransactionReceipt(ethClient, common.HexToHash(txHash), 120*time.Second, 15*time.Second)
					if err != nil {
						log.Fatalf("Transaction %s was not mined within the timeout period.\n", txHash)
						return
					}
					if receipt.Status == 0 {
						log.Fatalf("Transaction %s failed: %v", txHash, receipt.Status)
						return
					}
					// Print the block number
					fmt.Printf("Transaction %s was mined in block %d\n", txHash, receipt.BlockNumber)
				}
			}

			if numberOf == stepNumber {
				logs, err := getLogs(*client, txHash)
				if err != nil {
					log.Fatalf("Failed to get logs: %v", err)
					return
				}

				vaultId, err := getVaultId(logs)
				if err != nil {
					log.Fatalf("Failed to vault id: %v", err)
					return
				}
				// Print success message
				fmt.Println("Vault was created successfully.")
				fmt.Printf("TX HASH: 0x%s\nVAULT ID: %d\n", txHash, vaultId)

				// Read existing data from the file
				data, err := os.ReadFile("transaction_log.json")
				if err != nil {
					log.Fatalf("Error reading file: %v", err)
				}

				// Unmarshal the JSON data into a slice of TransactionData
				var transactions []TransactionDetails
				err = json.Unmarshal(data, &transactions)
				if err != nil {
					log.Fatalf("Error unmarshalling JSON: %v", err)
				}

				// Update the last transaction's vault_id
				if len(transactions) > 0 {
					transactions[len(transactions)-1].VaultID = vaultId
				} else {
					log.Println("No transactions found to update.")
				}

				// Marshal the updated data back to JSON
				updatedData, err := json.MarshalIndent(transactions, "", "  ")
				if err != nil {
					log.Fatalf("Error marshalling updated data: %v", err)
				}

				// Write the updated data back to the file
				err = os.WriteFile("transaction_log.json", updatedData, 0644)
				if err != nil {
					log.Fatalf("Error writing to file: %v", err)
				}

				// Print success message
				fmt.Println("Transaction details updated successfully.")
			}

		}
	} else {
		fmt.Println("\nNo steps found in the response.")
	}

}

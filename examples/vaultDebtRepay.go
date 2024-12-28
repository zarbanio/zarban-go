/*
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
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/zarban-go/service"
)

func toNative(amount float64) (*string, error) {
	// Convert amount to native units using big.Int
	scaledAmount := new(big.Int)
	floatAmount := new(big.Float).SetFloat64(amount)
	scale := new(big.Float).SetFloat64(math.Pow10(18))

	// Multiply amount by scale factor
	floatAmount.Mul(floatAmount, scale)

	// Convert to integer, checking for accuracy
	if _, accuracy := floatAmount.Int(scaledAmount); accuracy != big.Exact {
		return nil, fmt.Errorf("lost precision during conversion")
	}

	str := scaledAmount.String()
	return &str, nil
}

func getVaultTxSteps(
	client service.Client,
	walletAddress string,
	vaultId int,
	amount float64,
) (service.ChainActivity, error) {
	var nativeAmount *string
	var err error
	if amount > 0 {
		nativeAmount, err = toNative(amount)
		if err != nil {
			log.Fatalf("Error converting collateral amount: %v", err)
			return service.ChainActivity{}, err
		}

	}

	request := service.StablecoinSystemRepayZarTxRequest{
		Amount:  nativeAmount,
		User:    walletAddress,
		VaultId: vaultId,
	}

	httpResponse, err := client.RepayZarTransaction(context.Background(), request)
	if err != nil {
		log.Fatalf("Error during API call -> RepayZarTransaction: %v", err)
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

type EthereumTransaction struct {
	From     string   `json:"from"`     // Sender address
	To       string   `json:"to"`       // Recipient address
	Value    *big.Int `json:"value"`    // Amount of ETH to send (in Wei)
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
	file, err := os.Open("repay_transaction_log.json")
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
	file, err = os.Create("repay_transaction_log.json")
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
func VaultDebtRepayExample() {
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

	const VAULT_ID int = 0   // Update with the actual vault ID
	const AMOUNT float64 = 0 // Update with the amount to repay

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
	vaultSteps, err := getVaultTxSteps(
		*client,
		WALLET_ADDRESS,
		VAULT_ID,
		AMOUNT)
	if err != nil {
		log.Fatalf("Failed to get vault repayment steps: %v", err)
		return
	}

	numOfSteps := vaultSteps.NumberOfSteps
	stepNumber := vaultSteps.StepNumber
	steps := vaultSteps.Steps

	if len(steps) > 0 {
		for s := 0; s <= (numOfSteps - stepNumber); s++ {
			vaultSteps, err = getVaultTxSteps(
				*client,
				WALLET_ADDRESS,
				VAULT_ID,
				AMOUNT)
			if err != nil {
				log.Fatalf("Failed to get vault repayment steps: %v", err)
				return
			}
			numOfSteps = vaultSteps.NumberOfSteps
			stepNumber = vaultSteps.StepNumber
			steps = vaultSteps.Steps

			var txHash string
			for i, step := range steps {
				data, err := step.Data.AsPreparedTx()
				if err != nil {
					log.Fatalf("Failed to convert step to prepared tx: %v", err)
					return
				}
				label := data.Label["en-US"]
				fmt.Printf("Step %d: %s\n", i+1, label)
				if stepNumber-1 == i {
					fmt.Println("Processing...")
					methodParams := data.MethodParameters
					addressTo := methodParams.To
					calldata := methodParams.Calldata
					valueStr := methodParams.Value
					value, ok := new(big.Int).SetString(valueStr, 10)
					if !ok {
						log.Fatalf("failed to convert value to bigint")
						return
					}
					nonce, err := ethClient.PendingNonceAt(context.Background(), common.HexToAddress(account))
					if err != nil {
						log.Fatalf("Failed to get nonce: %v", err)
						return
					}
					gas := data.GasUseEstimate
					gasPrice, err := ethClient.SuggestGasPrice(context.Background())
					if err != nil {
						log.Fatalf("Failed to get gas price: %v", err)
						return
					}
					gasPrice = gasPrice.Mul(gasPrice, big.NewInt(110)) // Increase gas price
					gasPrice = gasPrice.Div(gasPrice, big.NewInt(100)) // by 10%
					calldataBytes, err := hexutil.Decode(calldata)
					if err != nil {
						log.Fatalf("Failed to decode calldata: %v", err)
						return
					}
					tx := types.NewTransaction(nonce, common.HexToAddress(addressTo), value, uint64(gas), gasPrice, calldataBytes)
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
					if err != nil {
						log.Fatalf("Failed to sign transaction: %v", err)
						return
					}

					signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
					if err != nil {
						log.Fatalf("Failed to sign the transaction: %v", err)
						return
					}

					err = ethClient.SendTransaction(context.Background(), signedTx)
					if err != nil {
						log.Fatalf("Failed to send transaction: %v", err)
						return
					}
					txHash = signedTx.Hash().Hex()
					err = SaveTransactionDetails(txToSave, txHash, VAULT_ID)
					if err != nil {
						log.Fatalf("Failed to save transaction details: %v", err)
						return
					}
					// Wait for the transaction to be mined
					receipt, err := waitForTransactionReceipt(ethClient, signedTx.Hash(), 1*time.Minute, 5*time.Second)
					if err != nil {
						log.Printf("Transaction %s was not mined: %v", txHash, err)
						continue
					}
					fmt.Printf("Transaction %s was mined in block %d\n", txHash, receipt.BlockNumber)
				}
			}
		}
	} else {
		log.Fatalf("No steps found in vault repayment transaction")
		return
	}
}

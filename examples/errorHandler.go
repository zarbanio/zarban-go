package examples

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/zarbanio/zarban-go/wallet"
)

// APIError represents a generic API error with dynamic details
type APIError struct {
	StatusCode int
	Message    string
	Details    interface{}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("APIError: status %d, message: %s, details: %+v", e.StatusCode, e.Message, e.Details)
}

// HandleAPIResponse processes API responses and distinguishes between Error and UserError types
func HandleAPIResponse[T any](resp *http.Response, successResponse *T) error {
	defer resp.Body.Close()

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Handle successful status codes
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		if err := json.Unmarshal(bodyBytes, successResponse); err != nil {
			return fmt.Errorf("failed to parse success response: %w", err)
		}
		return nil
	}

	// Check if response matches UserError structure
	var userError wallet.UserError
	if err := json.Unmarshal(bodyBytes, &userError); err == nil && len(userError.Messages) > 0 {
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    "User error",
			Details:    userError,
		}
	}

	// Check if response matches Error structure
	var genericError wallet.Error
	if err := json.Unmarshal(bodyBytes, &genericError); err == nil && genericError.Msg != "" {
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    "API error",
			Details:    genericError,
		}
	}

	// Fallback for unhandled responses
	return &APIError{
		StatusCode: resp.StatusCode,
		Message:    "Unhandled error",
		Details:    string(bodyBytes),
	}
}

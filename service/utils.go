package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// APIError represents a generic API error with dynamic details and stack trace support
type APIError struct {
	StatusCode   int
	Message      string
	Details      interface{}
	RequestID    string
	Path         string
	Method       string
	ErrorContext map[string]interface{}
}

// Error implements the error interface
func (e *APIError) Error() string {
	return fmt.Sprintf("APIError[%s]: status %d, path: %s, message: %s, details: %+v",
		e.RequestID, e.StatusCode, e.Path, e.Message, e.Details)
}

// WithContext adds additional context to the APIError
func (e *APIError) WithContext(key string, value interface{}) *APIError {
	if e.ErrorContext == nil {
		e.ErrorContext = make(map[string]interface{})
	}
	e.ErrorContext[key] = value
	return e
}

// IsNotFound returns true if the error represents a 404 status
func (e *APIError) IsNotFound() bool {
	return e.StatusCode == http.StatusNotFound
}

// IsUnauthorized returns true if the error represents a 401 status
func (e *APIError) IsUnauthorized() bool {
	return e.StatusCode == http.StatusUnauthorized
}

// IsForbidden returns true if the error represents a 403 status
func (e *APIError) IsForbidden() bool {
	return e.StatusCode == http.StatusForbidden
}

// IsClientError returns true if the error is in the 4xx range
func (e *APIError) IsClientError() bool {
	return e.StatusCode >= 400 && e.StatusCode < 500
}

// IsServerError returns true if the error is in the 5xx range
func (e *APIError) IsServerError() bool {
	return e.StatusCode >= 500 && e.StatusCode < 600
}

// HandleAPIResponse processes API responses with improved error handling and context support
func HandleAPIResponse[T any](ctx context.Context, resp *http.Response, successResponse *T) error {
	if resp == nil {
		return &APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    "nil response received",
		}
	}

	defer resp.Body.Close()

	// Extract request information
	requestID := resp.Header.Get("X-Request-ID")
	path := resp.Request.URL.Path
	method := resp.Request.Method

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return &APIError{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to read response body",
			RequestID:  requestID,
			Path:       path,
			Method:     method,
			Details:    err.Error(),
		}
	}

	// Handle successful status codes
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		if err := json.Unmarshal(bodyBytes, successResponse); err != nil {
			return &APIError{
				StatusCode: resp.StatusCode,
				Message:    "failed to parse success response",
				RequestID:  requestID,
				Path:       path,
				Method:     method,
				Details:    err.Error(),
			}
		}
		return nil
	}

	// Attempt to parse different error types
	apiError := &APIError{
		StatusCode: resp.StatusCode,
		RequestID:  requestID,
		Path:       path,
		Method:     method,
	}

	// Try UserError first
	var userError UserError
	if err := json.Unmarshal(bodyBytes, &userError); err == nil && len(userError.Messages) > 0 {
		apiError.Message = "User error"
		apiError.Details = userError
		return apiError
	}

	// Try generic Error
	var genericError Error
	if err := json.Unmarshal(bodyBytes, &genericError); err == nil && genericError.Msg != "" {
		apiError.Message = genericError.Msg
		apiError.Details = genericError
		return apiError
	}

	// Fallback for unhandled responses
	apiError.Message = "Unhandled error"
	apiError.Details = string(bodyBytes)
	return apiError
}

// PrettyPrintError formats the APIError with improved readability
func PrettyPrintError(err *APIError) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("API Error Report\n"))
	sb.WriteString(fmt.Sprintf("---------------\n"))
	sb.WriteString(fmt.Sprintf("Status:     %d\n", err.StatusCode))
	sb.WriteString(fmt.Sprintf("Request ID: %s\n", err.RequestID))
	sb.WriteString(fmt.Sprintf("Path:       %s %s\n", err.Method, err.Path))
	sb.WriteString(fmt.Sprintf("Message:    %s\n", err.Message))

	if err.ErrorContext != nil && len(err.ErrorContext) > 0 {
		sb.WriteString("\nContext:\n")
		for k, v := range err.ErrorContext {
			sb.WriteString(fmt.Sprintf("- %s: %v\n", k, v))
		}
	}

	if userError, ok := err.Details.(UserError); ok {
		sb.WriteString("\nUser Error Details:\n")
		for lang, detail := range userError.Messages {
			sb.WriteString(fmt.Sprintf("[%s]\n", lang))
			sb.WriteString(fmt.Sprintf("Message:   %s\n", detail.UserMessage))
			if len(detail.Solutions) > 0 {
				sb.WriteString("Solutions:\n")
				for _, solution := range detail.Solutions {
					sb.WriteString(fmt.Sprintf("- %s\n", solution))
				}
			}
		}
		return sb.String()
	}

	if genericError, ok := err.Details.(Error); ok {
		sb.WriteString("\nError Details:\n")
		sb.WriteString(fmt.Sprintf("Message: %s\n", genericError.Msg))
		if len(genericError.Reasons) > 0 {
			sb.WriteString("Reasons:\n")
			for _, reason := range genericError.Reasons {
				sb.WriteString(fmt.Sprintf("- %s\n", reason))
			}
		}
		return sb.String()
	}

	sb.WriteString("\nRaw Details:\n")
	sb.WriteString(fmt.Sprintf("%v\n", err.Details))

	return sb.String()
}
func AddHeaders(headers map[string]string) func(ctx context.Context, req *http.Request) error {
	return func(ctx context.Context, req *http.Request) error {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
		return nil
	}
}

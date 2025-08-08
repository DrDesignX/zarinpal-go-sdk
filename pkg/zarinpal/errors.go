package zarinpal

import (
	"fmt"
)

// Error represents a ZarinPal API error
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("zarinpal error (code: %d): %s - %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("zarinpal error (code: %d): %s", e.Code, e.Message)
}

func (e *Error) Unwrap() error {
	return e.Err
}

// HTTPError represents an HTTP-related error
type HTTPError struct {
	StatusCode int
	Body       string
	Err        error
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP error (status: %d): %s", e.StatusCode, e.Body)
}

func (e *HTTPError) Unwrap() error {
	return e.Err
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error for field '%s': %s", e.Field, e.Message)
}

// NewValidationError creates a new validation error
func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}

// NewHTTPError creates a new HTTP error
func NewHTTPError(statusCode int, body string, err error) *HTTPError {
	return &HTTPError{
		StatusCode: statusCode,
		Body:       body,
		Err:        err,
	}
}

// NewError creates a new ZarinPal API error
func NewError(code int, message string, err error) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// IsHTTPError checks if an error is an HTTP error
func IsHTTPError(err error) bool {
	_, ok := err.(*HTTPError)
	return ok
}

// IsValidationError checks if an error is a validation error
func IsValidationError(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}

// GetHTTPStatusCode extracts HTTP status code from error if it's an HTTP error
func GetHTTPStatusCode(err error) int {
	if httpErr, ok := err.(*HTTPError); ok {
		return httpErr.StatusCode
	}
	return 0
}

// Common error codes from ZarinPal API
const (
	ErrorCodeSuccess                = 100
	ErrorCodeInvalidMerchant        = -1
	ErrorCodeInvalidParameters      = -2
	ErrorCodeMerchantNotFound       = -3
	ErrorCodeMerchantNotActive      = -4
	ErrorCodeDuplicateReference     = -5
	ErrorCodeInvalidAmount          = -6
	ErrorCodeTransactionNotFound    = -7
	ErrorCodeTransactionNotVerified = -8
	ErrorCodeInvalidAuthority       = -9
	ErrorCodeInvalidIP              = -10
)

// Predefined errors for common cases
var (
	ErrInvalidMerchantID = NewValidationError("merchant_id", "invalid merchant ID format")
	ErrInvalidAuthority  = NewValidationError("authority", "invalid authority format")
	ErrInvalidAmount     = NewValidationError("amount", "amount must be at least 10000")
	ErrInvalidURL        = NewValidationError("callback_url", "invalid callback URL format")
	ErrMissingParameter  = NewValidationError("", "missing required parameter")
)
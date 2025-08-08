package resources

import (
	"context"
)

// ClientInterface defines the interface that clients must implement to work with resources
type ClientInterface interface {
	Request(ctx context.Context, method, endpoint string, data interface{}) ([]byte, error)
	GraphQLRequest(ctx context.Context, query string, variables map[string]interface{}) ([]byte, error)
	GetConfig() ConfigInterface
	GetValidator() ValidatorInterface
	GetBaseURL() string
}

// ConfigInterface defines the interface for configuration
type ConfigInterface interface {
	BaseURL() string
	GraphQLURL() string
}

// ValidatorInterface defines the interface for validation
type ValidatorInterface interface {
	ValidateAmount(amount int64, minAmount int64) error
	ValidateCallbackURL(callbackURL string) error
	ValidateAuthority(authority string) error
	ValidateTerminalID(terminalID string) error
	ValidateFilter(filter string) error
	ValidateLimit(limit int) error
	ValidateOffset(offset int) error
	ValidateSessionID(sessionID string) error
	ValidateMethod(method string) error
	ValidateReason(reason string) error
	ValidateCurrency(currency string) error
}

// PaymentService defines the interface for payment operations
type PaymentService interface {
	Create(ctx context.Context, request *PaymentRequest) (*PaymentResponse, error)
	GeneratePaymentURL(authority string) string
}

// VerificationService defines the interface for payment verification
type VerificationService interface {
	Verify(ctx context.Context, request *VerificationRequest) (*VerificationResponse, error)
}

// InquiryService defines the interface for transaction inquiries
type InquiryService interface {
	Inquire(ctx context.Context, request *InquiryRequest) (*InquiryResponse, error)
}

// RefundService defines the interface for refund operations
type RefundService interface {
	Create(ctx context.Context, request *RefundRequest) (*RefundResponse, error)
}

// ReversalService defines the interface for transaction reversals
type ReversalService interface {
	Reverse(ctx context.Context, request *ReversalRequest) (*ReversalResponse, error)
}

// TransactionService defines the interface for transaction listing
type TransactionService interface {
	List(ctx context.Context, request *TransactionRequest) (*TransactionResponse, error)
}

// UnverifiedService defines the interface for unverified payment listing
type UnverifiedService interface {
	List(ctx context.Context) (*UnverifiedResponse, error)
}

// FeeService defines the interface for fee calculation
type FeeService interface {
	Calculate(ctx context.Context, request *FeeRequest) (*FeeResponse, error)
}
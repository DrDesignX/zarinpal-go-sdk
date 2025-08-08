package zarinpal

import (
	"github.com/DrDesignX/zarinpal-go-sdk/pkg/zarinpal/resources"
)

// ClientWithResources extends the basic client with resource services
type ClientWithResources struct {
	*Client
	
	// Resources
	Payments      resources.PaymentService
	Verifications resources.VerificationService
	Inquiries     resources.InquiryService
	Refunds       resources.RefundService
	Reversals     resources.ReversalService
	Transactions  resources.TransactionService
	Unverified    resources.UnverifiedService
	Fee           resources.FeeService
}

// NewClientWithResources creates a new ZarinPal client with all resource services
func NewClientWithResources(config *Config, opts ...ClientOption) *ClientWithResources {
	client := NewClient(config, opts...)
	
	clientWithResources := &ClientWithResources{
		Client: client,
	}

	// Initialize resources
	clientWithResources.Payments = resources.NewPaymentService(client)
	clientWithResources.Verifications = resources.NewVerificationService(client)
	clientWithResources.Inquiries = resources.NewInquiryService(client)
	clientWithResources.Refunds = resources.NewRefundService(client)
	clientWithResources.Reversals = resources.NewReversalService(client)
	clientWithResources.Transactions = resources.NewTransactionService(client)
	clientWithResources.Unverified = resources.NewUnverifiedService(client)
	clientWithResources.Fee = resources.NewFeeService(client)

	return clientWithResources
}
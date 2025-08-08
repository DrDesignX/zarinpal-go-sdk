package services

import (
	"context"
	"fmt"
	"log"

	"github.com/DrDesignX/zarinpal-go-sdk/pkg/zarinpal"
	"github.com/DrDesignX/zarinpal-go-sdk/pkg/zarinpal/resources"
)

// ZarinPalService handles ZarinPal payment gateway integration
type ZarinPalService struct {
	client *zarinpal.ClientWithResources
	config *zarinpal.Config
}

// NewZarinPalService creates a new ZarinPal service instance
func NewZarinPalService(merchantID string, sandbox bool) (*ZarinPalService, error) {
	config, err := zarinpal.NewConfig(
		zarinpal.WithMerchantID(merchantID),
		zarinpal.WithSandbox(sandbox),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create ZarinPal config: %w", err)
	}

	client := zarinpal.NewClientWithResources(config)

	return &ZarinPalService{
		client: client,
		config: config,
	}, nil
}

// PaymentRequest represents a payment request
type PaymentRequest struct {
	Amount      int64  // Amount in Rials (changed from int to int64)
	UserID      int64  // User ID for tracking
	Description string // Payment description
	CallbackURL string // Callback URL
}

// PaymentResponse represents a payment response
type PaymentResponse struct {
	Authority  string
	PaymentURL string
	Status     bool
	Message    string
}

// CreatePayment creates a new payment request
func (z *ZarinPalService) CreatePayment(ctx context.Context, req PaymentRequest) (*PaymentResponse, error) {
	log.Printf("Creating ZarinPal payment: Amount=%d, UserID=%d", req.Amount, req.UserID)

	paymentRequest := &resources.PaymentRequest{
		Amount:      req.Amount, // Now int64, no conversion needed
		Description: req.Description,
		CallbackURL: req.CallbackURL,
		// Note: Currency and Metadata fields don't exist in the actual PaymentRequest struct
		// You can store user_id in description or use ReferrerID field if needed
		ReferrerID: fmt.Sprintf("%d", req.UserID), // Store user ID as referrer
	}

	response, err := z.client.Payments.Create(ctx, paymentRequest)
	if err != nil {
		log.Printf("ZarinPal payment creation failed: %v", err)
		return &PaymentResponse{
			Status:  false,
			Message: "خطا در ایجاد پرداخت",
		}, err
	}

	if response.Data.Code != 100 {
		log.Printf("ZarinPal payment creation failed with code: %d, message: %s", response.Data.Code, response.Data.Message)
		return &PaymentResponse{
			Status:  false,
			Message: response.Data.Message,
		}, fmt.Errorf("payment creation failed: %s", response.Data.Message)
	}

	authority := response.Data.Authority
	paymentURL := z.client.Payments.GeneratePaymentURL(authority)

	log.Printf("ZarinPal payment created successfully: Authority=%s", authority)

	return &PaymentResponse{
		Authority:  authority,
		PaymentURL: paymentURL,
		Status:     true,
		Message:    "پرداخت با موفقیت ایجاد شد",
	}, nil
}

// VerificationRequest represents a verification request
type VerificationRequest struct {
	Authority string
	Amount    int64 // Changed from int to int64
}

// VerificationResponse represents a verification response
type VerificationResponse struct {
	Status  bool
	RefID   int64
	Message string
	Code    int
}

// VerifyPayment verifies a payment
func (z *ZarinPalService) VerifyPayment(ctx context.Context, req VerificationRequest) (*VerificationResponse, error) {
	log.Printf("Verifying ZarinPal payment: Authority=%s, Amount=%d", req.Authority, req.Amount)

	verifyRequest := &resources.VerificationRequest{
		Amount:    req.Amount, // Now int64, no conversion needed
		Authority: req.Authority,
	}

	response, err := z.client.Verifications.Verify(ctx, verifyRequest)
	if err != nil {
		log.Printf("ZarinPal verification failed: %v", err)
		return &VerificationResponse{
			Status:  false,
			Message: "خطا در تایید پرداخت",
		}, err
	}

	if response.Data.Code == 100 {
		log.Printf("ZarinPal payment verified successfully: RefID=%d", response.Data.RefID)
		return &VerificationResponse{
			Status:  true,
			RefID:   response.Data.RefID, // Already int64, no conversion needed
			Message: "پرداخت با موفقیت تایید شد",
			Code:    response.Data.Code,
		}, nil
	} else if response.Data.Code == 101 {
		log.Printf("ZarinPal payment already verified: RefID=%d", response.Data.RefID)
		return &VerificationResponse{
			Status:  true,
			RefID:   response.Data.RefID, // Already int64, no conversion needed
			Message: "پرداخت قبلاً تایید شده است",
			Code:    response.Data.Code,
		}, nil
	} else {
		log.Printf("ZarinPal verification failed with code: %d, message: %s", response.Data.Code, response.Data.Message)
		return &VerificationResponse{
			Status:  false,
			Message: response.Data.Message,
			Code:    response.Data.Code,
		}, nil
	}
}

// CalculateFee calculates transaction fee
func (z *ZarinPalService) CalculateFee(ctx context.Context, amount int64) (int64, error) { // Changed both params to int64
	feeRequest := &resources.FeeRequest{
		Amount: amount, // Now int64, no conversion needed
		// Note: Currency field exists in FeeRequest, but it's optional
		Currency: "IRR",
	}

	response, err := z.client.Fee.Calculate(ctx, feeRequest)
	if err != nil {
		return 0, err
	}

	return response.Data.Fee, nil // Already int64, no conversion needed
}

package resources

import (
	"context"
	"encoding/json"
	"fmt"
)

// paymentService implements the PaymentService interface
type paymentService struct {
	client ClientInterface
}

// NewPaymentService creates a new payment service
func NewPaymentService(client ClientInterface) PaymentService {
	return &paymentService{
		client: client,
	}
}

// Create creates a new payment request
func (s *paymentService) Create(ctx context.Context, request *PaymentRequest) (*PaymentResponse, error) {
	// Validate required fields
	if request.Amount == 0 {
		return nil, fmt.Errorf("amount is required")
	}
	if request.CallbackURL == "" {
		return nil, fmt.Errorf("callback_url is required")
	}

	// Validate input data using validator
	validator := s.client.GetValidator()
	
	if err := validator.ValidateAmount(request.Amount, 10000); err != nil {
		return nil, err
	}
	
	if err := validator.ValidateCallbackURL(request.CallbackURL); err != nil {
		return nil, err
	}

	// Convert request to map for API call
	requestData := map[string]interface{}{
		"amount":       request.Amount,
		"callback_url": request.CallbackURL,
	}

	if request.Description != "" {
		requestData["description"] = request.Description
	}
	if request.Mobile != "" {
		requestData["mobile"] = request.Mobile
	}
	if request.Email != "" {
		requestData["email"] = request.Email
	}
	if len(request.CardPAN) > 0 {
		requestData["card_pan"] = request.CardPAN
	}
	if request.ReferrerID != "" {
		requestData["referrer_id"] = request.ReferrerID
	}

	// Make API request
	endpoint := "/pg/v4/payment/request.json"
	responseBody, err := s.client.Request(ctx, "POST", endpoint, requestData)
	if err != nil {
		return nil, fmt.Errorf("payment request failed: %w", err)
	}

	// Parse response
	var response PaymentResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse payment response: %w", err)
	}

	return &response, nil
}

// GeneratePaymentURL generates the payment URL using the authority code
func (s *paymentService) GeneratePaymentURL(authority string) string {
	baseURL := s.client.GetBaseURL()
	return fmt.Sprintf("%s/pg/StartPay/%s", baseURL, authority)
}
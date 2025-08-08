package resources

import (
	"context"
	"encoding/json"
	"fmt"
)

// verificationService implements the VerificationService interface
type verificationService struct {
	client ClientInterface
}

// NewVerificationService creates a new verification service
func NewVerificationService(client ClientInterface) VerificationService {
	return &verificationService{
		client: client,
	}
}

// Verify verifies a payment transaction
func (s *verificationService) Verify(ctx context.Context, request *VerificationRequest) (*VerificationResponse, error) {
	// Validate required fields
	if request.Amount == 0 {
		return nil, fmt.Errorf("amount is required")
	}
	if request.Authority == "" {
		return nil, fmt.Errorf("authority is required")
	}

	// Validate input data using validator
	validator := s.client.GetValidator()
	
	if err := validator.ValidateAmount(request.Amount, 10000); err != nil {
		return nil, err
	}
	
	if err := validator.ValidateAuthority(request.Authority); err != nil {
		return nil, err
	}

	// Convert request to map for API call
	requestData := map[string]interface{}{
		"amount":    request.Amount,
		"authority": request.Authority,
	}

	// Make API request
	endpoint := "/pg/v4/payment/verify.json"
	responseBody, err := s.client.Request(ctx, "POST", endpoint, requestData)
	if err != nil {
		return nil, fmt.Errorf("verification request failed: %w", err)
	}

	// Parse response
	var response VerificationResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse verification response: %w", err)
	}

	return &response, nil
}
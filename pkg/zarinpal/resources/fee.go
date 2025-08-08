package resources

import (
	"context"
	"encoding/json"
	"fmt"
)

// feeService implements the FeeService interface
type feeService struct {
	client ClientInterface
}

// NewFeeService creates a new fee service
func NewFeeService(client ClientInterface) FeeService {
	return &feeService{
		client: client,
	}
}

// Calculate calculates the transaction fee
func (s *feeService) Calculate(ctx context.Context, request *FeeRequest) (*FeeResponse, error) {
	// Validate required fields
	if request.Amount == 0 {
		return nil, fmt.Errorf("amount is required")
	}

	// Validate input data using validator
	validator := s.client.GetValidator()
	
	if err := validator.ValidateAmount(request.Amount, 10000); err != nil {
		return nil, err
	}

	if request.Currency != "" {
		if err := validator.ValidateCurrency(request.Currency); err != nil {
			return nil, err
		}
	}

	// Convert request to map for API call
	requestData := map[string]interface{}{
		"amount": request.Amount,
	}

	if request.Currency != "" {
		requestData["currency"] = request.Currency
	}

	// Make API request
	endpoint := "/pg/v4/payment/feeCalculation.json"
	responseBody, err := s.client.Request(ctx, "POST", endpoint, requestData)
	if err != nil {
		return nil, fmt.Errorf("fee calculation request failed: %w", err)
	}

	// Parse response
	var response FeeResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse fee response: %w", err)
	}

	return &response, nil
}
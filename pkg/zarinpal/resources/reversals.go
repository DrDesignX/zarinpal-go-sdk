package resources

import (
	"context"
	"encoding/json"
	"fmt"
)

// reversalService implements the ReversalService interface
type reversalService struct {
	client ClientInterface
}

// NewReversalService creates a new reversal service
func NewReversalService(client ClientInterface) ReversalService {
	return &reversalService{
		client: client,
	}
}

// Reverse reverses a transaction
func (s *reversalService) Reverse(ctx context.Context, request *ReversalRequest) (*ReversalResponse, error) {
	// Validate required fields
	if request.Authority == "" {
		return nil, fmt.Errorf("authority is required")
	}

	// Validate input data using validator
	validator := s.client.GetValidator()
	
	if err := validator.ValidateAuthority(request.Authority); err != nil {
		return nil, err
	}

	// Convert request to map for API call
	requestData := map[string]interface{}{
		"authority": request.Authority,
	}

	// Make API request
	endpoint := "/pg/v4/payment/reverse.json"
	responseBody, err := s.client.Request(ctx, "POST", endpoint, requestData)
	if err != nil {
		return nil, fmt.Errorf("reversal request failed: %w", err)
	}

	// Parse response
	var response ReversalResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse reversal response: %w", err)
	}

	return &response, nil
}
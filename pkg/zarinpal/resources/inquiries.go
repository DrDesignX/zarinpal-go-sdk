package resources

import (
	"context"
	"encoding/json"
	"fmt"
)

// inquiryService implements the InquiryService interface
type inquiryService struct {
	client ClientInterface
}

// NewInquiryService creates a new inquiry service
func NewInquiryService(client ClientInterface) InquiryService {
	return &inquiryService{
		client: client,
	}
}

// Inquire inquires about the status of a transaction
func (s *inquiryService) Inquire(ctx context.Context, request *InquiryRequest) (*InquiryResponse, error) {
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
	endpoint := "/pg/v4/payment/inquiry.json"
	responseBody, err := s.client.Request(ctx, "POST", endpoint, requestData)
	if err != nil {
		return nil, fmt.Errorf("inquiry request failed: %w", err)
	}

	// Parse response
	var response InquiryResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse inquiry response: %w", err)
	}

	return &response, nil
}
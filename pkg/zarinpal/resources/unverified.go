package resources

import (
	"context"
	"encoding/json"
	"fmt"
)

// unverifiedService implements the UnverifiedService interface
type unverifiedService struct {
	client ClientInterface
}

// NewUnverifiedService creates a new unverified service
func NewUnverifiedService(client ClientInterface) UnverifiedService {
	return &unverifiedService{
		client: client,
	}
}

// List retrieves a list of unverified payments
func (s *unverifiedService) List(ctx context.Context) (*UnverifiedResponse, error) {
	// Make API request
	endpoint := "/pg/v4/payment/unVerified.json"
	responseBody, err := s.client.Request(ctx, "POST", endpoint, map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("unverified list request failed: %w", err)
	}

	// Parse response
	var response UnverifiedResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse unverified response: %w", err)
	}

	return &response, nil
}
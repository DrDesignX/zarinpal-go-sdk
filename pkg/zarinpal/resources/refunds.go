package resources

import (
	"context"
	"encoding/json"
	"fmt"
)

// refundService implements the RefundService interface
type refundService struct {
	client ClientInterface
}

// NewRefundService creates a new refund service
func NewRefundService(client ClientInterface) RefundService {
	return &refundService{
		client: client,
	}
}

// Create creates a new refund request
func (s *refundService) Create(ctx context.Context, request *RefundRequest) (*RefundResponse, error) {
	// Validate required fields
	if request.SessionID == "" {
		return nil, fmt.Errorf("session_id is required")
	}
	if request.Amount == 0 {
		return nil, fmt.Errorf("amount is required")
	}

	// Validate input data using validator
	validator := s.client.GetValidator()
	
	if err := validator.ValidateSessionID(request.SessionID); err != nil {
		return nil, err
	}
	
	if err := validator.ValidateAmount(request.Amount, 10000); err != nil {
		return nil, err
	}

	if request.Method != "" {
		if err := validator.ValidateMethod(request.Method); err != nil {
			return nil, err
		}
	}

	if request.Reason != "" {
		if err := validator.ValidateReason(request.Reason); err != nil {
			return nil, err
		}
	}

	// Prepare GraphQL query
	query := `
		mutation AddRefund($session_id: ID!, $amount: BigInteger!, $description: String, $method: InstantPayoutActionTypeEnum, $reason: RefundReasonEnum) {
			resource: AddRefund(
				session_id: $session_id,
				amount: $amount,
				description: $description,
				method: $method,
				reason: $reason
			) {
				terminal_id,
				id,
				amount,
				timeline {
					refund_amount,
					refund_time,
					refund_status
				}
			}
		}
	`

	// Prepare variables
	variables := map[string]interface{}{
		"session_id": request.SessionID,
		"amount":     request.Amount,
	}

	if request.Description != "" {
		variables["description"] = request.Description
	}
	if request.Method != "" {
		variables["method"] = request.Method
	}
	if request.Reason != "" {
		variables["reason"] = request.Reason
	}

	// Make GraphQL request
	responseBody, err := s.client.GraphQLRequest(ctx, query, variables)
	if err != nil {
		return nil, fmt.Errorf("refund request failed: %w", err)
	}

	// Parse response
	var response RefundResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse refund response: %w", err)
	}

	return &response, nil
}
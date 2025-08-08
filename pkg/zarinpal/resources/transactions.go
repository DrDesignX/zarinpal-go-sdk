package resources

import (
	"context"
	"encoding/json"
	"fmt"
)

// transactionService implements the TransactionService interface
type transactionService struct {
	client ClientInterface
}

// NewTransactionService creates a new transaction service
func NewTransactionService(client ClientInterface) TransactionService {
	return &transactionService{
		client: client,
	}
}

// List retrieves a list of transactions via GraphQL
func (s *transactionService) List(ctx context.Context, request *TransactionRequest) (*TransactionResponse, error) {
	// Validate required fields
	if request.TerminalID == "" {
		return nil, fmt.Errorf("terminal_id is required")
	}

	// Validate input data using validator
	validator := s.client.GetValidator()
	
	if err := validator.ValidateTerminalID(request.TerminalID); err != nil {
		return nil, err
	}

	if request.Filter != "" {
		if err := validator.ValidateFilter(request.Filter); err != nil {
			return nil, err
		}
	}

	if request.Limit > 0 {
		if err := validator.ValidateLimit(request.Limit); err != nil {
			return nil, err
		}
	}

	if request.Offset > 0 {
		if err := validator.ValidateOffset(request.Offset); err != nil {
			return nil, err
		}
	}

	// Prepare GraphQL query
	query := `
		query GetTransactions($terminal_id: ID!, $filter: FilterEnum, $limit: Int, $offset: Int) {
			Session: Session(
				terminal_id: $terminal_id,
				filter: $filter,
				limit: $limit,
				offset: $offset
			) {
				id,
				status,
				amount,
				description,
				created_at
			}
		}
	`

	// Prepare variables
	variables := map[string]interface{}{
		"terminal_id": request.TerminalID,
	}

	if request.Filter != "" {
		variables["filter"] = request.Filter
	}
	if request.Limit > 0 {
		variables["limit"] = request.Limit
	}
	if request.Offset > 0 {
		variables["offset"] = request.Offset
	}

	// Make GraphQL request
	responseBody, err := s.client.GraphQLRequest(ctx, query, variables)
	if err != nil {
		return nil, fmt.Errorf("transaction list request failed: %w", err)
	}

	// Parse response
	var response TransactionResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse transaction response: %w", err)
	}

	return &response, nil
}
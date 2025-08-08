package zarinpal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/DrDesignX/zarinpal-go-sdk/pkg/zarinpal/resources"
)

// Client represents the main ZarinPal SDK client
type Client struct {
	config     *Config
	httpClient *http.Client
	validator  *Validator
}

// ClientOption represents a functional option for configuring the client
type ClientOption func(*Client)

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithTimeout sets request timeout
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// NewClient creates a new ZarinPal client
func NewClient(config *Config, opts ...ClientOption) *Client {
	client := &Client{
		config: config,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		validator: NewValidator(),
	}

	// Apply options
	for _, opt := range opts {
		opt(client)
	}

	return client
}

// Request makes a generic HTTP request to the ZarinPal API
func (c *Client) Request(ctx context.Context, method, endpoint string, data interface{}) ([]byte, error) {
	url := c.config.BaseURL() + endpoint
	
	// Prepare request body
	var requestBody io.Reader
	if data != nil {
		// Add merchant_id to the request data
		if dataMap, ok := data.(map[string]interface{}); ok {
			dataMap["merchant_id"] = c.config.MerchantID
		} else {
			// Convert to map and add merchant_id
			dataBytes, err := json.Marshal(data)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal request data: %w", err)
			}
			
			var dataMap map[string]interface{}
			if err := json.Unmarshal(dataBytes, &dataMap); err != nil {
				return nil, fmt.Errorf("failed to unmarshal request data: %w", err)
			}
			
			dataMap["merchant_id"] = c.config.MerchantID
			data = dataMap
		}
		
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request data: %w", err)
		}
		requestBody = bytes.NewReader(jsonData)
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, method, url, requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("User-Agent", "ZarinPalSdk/v1 (Go)")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Make request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode >= 400 {
		return nil, NewHTTPError(resp.StatusCode, string(body), fmt.Errorf("HTTP %d", resp.StatusCode))
	}

	return body, nil
}

// GraphQLRequest makes a GraphQL request to the ZarinPal API
func (c *Client) GraphQLRequest(ctx context.Context, query string, variables map[string]interface{}) ([]byte, error) {
	url := c.config.GraphQLURL()
	
	// Prepare GraphQL request
	requestData := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}
	
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal GraphQL request: %w", err)
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create GraphQL request: %w", err)
	}

	// Set headers
	req.Header.Set("User-Agent", "ZarinPalSdk/v1 (Go)")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if c.config.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.AccessToken)
	}

	// Make request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GraphQL request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read GraphQL response body: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode >= 400 {
		return nil, NewHTTPError(resp.StatusCode, string(body), fmt.Errorf("GraphQL HTTP %d", resp.StatusCode))
	}

	return body, nil
}

// GetConfig returns the client configuration
func (c *Client) GetConfig() resources.ConfigInterface {
	return c.config
}

// GetValidator returns the validator instance
func (c *Client) GetValidator() resources.ValidatorInterface {
	return c.validator
}

// GetBaseURL returns the base URL for API requests
func (c *Client) GetBaseURL() string {
	return c.config.BaseURL()
}
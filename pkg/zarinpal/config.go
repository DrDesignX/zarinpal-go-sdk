package zarinpal

import (
	"errors"
	"os"
)

// Config represents the configuration for the ZarinPal SDK
type Config struct {
	MerchantID  string
	AccessToken string
	Sandbox     bool
}

// ConfigOption represents a functional option for configuring the client
type ConfigOption func(*Config) error

// NewConfig creates a new Config with the provided options
func NewConfig(opts ...ConfigOption) (*Config, error) {
	config := &Config{
		Sandbox: false, // Default to production
	}

	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, err
		}
	}

	return config, nil
}

// WithMerchantID sets the merchant ID
func WithMerchantID(merchantID string) ConfigOption {
	return func(c *Config) error {
		if merchantID == "" {
			return errors.New("merchant ID cannot be empty")
		}
		c.MerchantID = merchantID
		return nil
	}
}

// WithAccessToken sets the access token
func WithAccessToken(accessToken string) ConfigOption {
	return func(c *Config) error {
		if accessToken == "" {
			return errors.New("access token cannot be empty")
		}
		c.AccessToken = accessToken
		return nil
	}
}

// WithSandbox enables or disables sandbox mode
func WithSandbox(sandbox bool) ConfigOption {
	return func(c *Config) error {
		c.Sandbox = sandbox
		return nil
	}
}

// WithEnvVars loads configuration from environment variables
func WithEnvVars() ConfigOption {
	return func(c *Config) error {
		if merchantID := os.Getenv("ZARINPAL_MERCHANT_ID"); merchantID != "" {
			c.MerchantID = merchantID
		}
		if accessToken := os.Getenv("ZARINPAL_ACCESS_TOKEN"); accessToken != "" {
			c.AccessToken = accessToken
		}
		if sandbox := os.Getenv("ZARINPAL_SANDBOX"); sandbox == "true" {
			c.Sandbox = true
		}
		return nil
	}
}

// BaseURL returns the appropriate base URL based on sandbox mode
func (c *Config) BaseURL() string {
	if c.Sandbox {
		return "https://sandbox.zarinpal.com"
	}
	return "https://payment.zarinpal.com"
}

// GraphQLURL returns the GraphQL endpoint URL
func (c *Config) GraphQLURL() string {
	return "https://next.zarinpal.com/api/v4/graphql/"
}
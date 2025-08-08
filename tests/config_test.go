package zarinpal

import (
	"testing"

	"github.com/DrDesignX/zarinpal-go-sdk/pkg/zarinpal"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name        string
		opts        []zarinpal.ConfigOption
		expectError bool
	}{
		{
			name: "valid config with merchant ID and access token",
			opts: []zarinpal.ConfigOption{
				zarinpal.WithMerchantID("123e4567-e89b-12d3-a456-426614174000"),
				zarinpal.WithAccessToken("test-token"),
				zarinpal.WithSandbox(true),
			},
			expectError: false,
		},
		{
			name: "empty merchant ID should fail",
			opts: []zarinpal.ConfigOption{
				zarinpal.WithMerchantID(""),
			},
			expectError: true,
		},
		{
			name: "empty access token should fail",
			opts: []zarinpal.ConfigOption{
				zarinpal.WithAccessToken(""),
			},
			expectError: true,
		},
		{
			name:        "default config should work",
			opts:        []zarinpal.ConfigOption{},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := zarinpal.NewConfig(tt.opts...)
			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("expected no error but got: %v", err)
			}
			if !tt.expectError && config == nil {
				t.Errorf("expected config but got nil")
			}
		})
	}
}

func TestConfigBaseURL(t *testing.T) {
	tests := []struct {
		name     string
		sandbox  bool
		expected string
	}{
		{
			name:     "sandbox mode",
			sandbox:  true,
			expected: "https://sandbox.zarinpal.com",
		},
		{
			name:     "production mode",
			sandbox:  false,
			expected: "https://payment.zarinpal.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &zarinpal.Config{Sandbox: tt.sandbox}
			if got := config.BaseURL(); got != tt.expected {
				t.Errorf("BaseURL() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestConfigGraphQLURL(t *testing.T) {
	config := &zarinpal.Config{}
	expected := "https://next.zarinpal.com/api/v4/graphql/"
	if got := config.GraphQLURL(); got != expected {
		t.Errorf("GraphQLURL() = %v, want %v", got, expected)
	}
}
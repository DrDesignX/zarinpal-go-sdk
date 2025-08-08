package zarinpal

import (
	"testing"

	"github.com/DrDesignX/zarinpal-go-sdk/pkg/zarinpal"
)

func TestValidateAmount(t *testing.T) {
	validator := zarinpal.NewValidator()

	tests := []struct {
		name        string
		amount      int64
		minAmount   int64
		expectError bool
	}{
		{
			name:        "valid amount",
			amount:      15000,
			minAmount:   10000,
			expectError: false,
		},
		{
			name:        "amount below minimum",
			amount:      5000,
			minAmount:   10000,
			expectError: true,
		},
		{
			name:        "amount equal to minimum",
			amount:      10000,
			minAmount:   10000,
			expectError: false,
		},
		{
			name:        "use default minimum when zero",
			amount:      15000,
			minAmount:   0,
			expectError: false,
		},
		{
			name:        "amount below default minimum",
			amount:      5000,
			minAmount:   0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateAmount(tt.amount, tt.minAmount)
			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("expected no error but got: %v", err)
			}
		})
	}
}

func TestValidateAuthority(t *testing.T) {
	validator := zarinpal.NewValidator()

	tests := []struct {
		name        string
		authority   string
		expectError bool
	}{
		{
			name:        "valid authority with A",
			authority:   "A0000000000000000000000000006qpmlj8d",
			expectError: false,
		},
		{
			name:        "valid authority with S",
			authority:   "S0000000000000000000000000006qpmlj8d",
			expectError: false,
		},
		{
			name:        "invalid authority - wrong prefix",
			authority:   "B0000000000000000000000000006qpmlj8d",
			expectError: true,
		},
		{
			name:        "invalid authority - too short",
			authority:   "A00000000000000000000000000",
			expectError: true,
		},
		{
			name:        "invalid authority - too long",
			authority:   "A0000000000000000000000000006qpmlj8dd",
			expectError: true,
		},
		{
			name:        "empty authority",
			authority:   "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateAuthority(tt.authority)
			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("expected no error but got: %v", err)
			}
		})
	}
}

func TestValidateCallbackURL(t *testing.T) {
	validator := zarinpal.NewValidator()

	tests := []struct {
		name        string
		url         string
		expectError bool
	}{
		{
			name:        "valid https URL",
			url:         "https://example.com/callback",
			expectError: false,
		},
		{
			name:        "valid http URL",
			url:         "http://example.com/callback",
			expectError: false,
		},
		{
			name:        "invalid URL - no protocol",
			url:         "example.com/callback",
			expectError: true,
		},
		{
			name:        "invalid URL - wrong protocol",
			url:         "ftp://example.com/callback",
			expectError: true,
		},
		{
			name:        "empty URL",
			url:         "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateCallbackURL(tt.url)
			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("expected no error but got: %v", err)
			}
		})
	}
}

func TestValidateMobile(t *testing.T) {
	validator := zarinpal.NewValidator()

	tests := []struct {
		name        string
		mobile      string
		expectError bool
	}{
		{
			name:        "valid mobile",
			mobile:      "09123456789",
			expectError: false,
		},
		{
			name:        "empty mobile (optional)",
			mobile:      "",
			expectError: false,
		},
		{
			name:        "invalid mobile - wrong prefix",
			mobile:      "08123456789",
			expectError: true,
		},
		{
			name:        "invalid mobile - too short",
			mobile:      "091234567",
			expectError: true,
		},
		{
			name:        "invalid mobile - too long",
			mobile:      "091234567890",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateMobile(tt.mobile)
			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("expected no error but got: %v", err)
			}
		})
	}
}

func TestValidateEmail(t *testing.T) {
	validator := zarinpal.NewValidator()

	tests := []struct {
		name        string
		email       string
		expectError bool
	}{
		{
			name:        "valid email",
			email:       "test@example.com",
			expectError: false,
		},
		{
			name:        "empty email (optional)",
			email:       "",
			expectError: false,
		},
		{
			name:        "invalid email - no @",
			email:       "testexample.com",
			expectError: true,
		},
		{
			name:        "invalid email - no domain",
			email:       "test@",
			expectError: true,
		},
		{
			name:        "invalid email - no TLD",
			email:       "test@example",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateEmail(tt.email)
			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("expected no error but got: %v", err)
			}
		})
	}
}

func TestValidateCurrency(t *testing.T) {
	validator := zarinpal.NewValidator()

	tests := []struct {
		name        string
		currency    string
		expectError bool
	}{
		{
			name:        "valid currency IRR",
			currency:    "IRR",
			expectError: false,
		},
		{
			name:        "valid currency IRT",
			currency:    "IRT",
			expectError: false,
		},
		{
			name:        "empty currency (optional)",
			currency:    "",
			expectError: false,
		},
		{
			name:        "invalid currency",
			currency:    "USD",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateCurrency(tt.currency)
			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("expected no error but got: %v", err)
			}
		})
	}
}
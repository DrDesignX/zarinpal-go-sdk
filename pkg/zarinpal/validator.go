package zarinpal

import (
	"fmt"
	"regexp"
	"strings"
)

// Validator provides validation utilities for ZarinPal API parameters
type Validator struct{}

// NewValidator creates a new validator instance
func NewValidator() *Validator {
	return &Validator{}
}

// ValidateCurrency validates currency code
func (v *Validator) ValidateCurrency(currency string) error {
	if currency == "" {
		return nil // Currency is optional
	}
	
	validCurrencies := []string{"IRR", "IRT"}
	for _, valid := range validCurrencies {
		if currency == valid {
			return nil
		}
	}
	
	return NewValidationError("currency", "invalid currency format. Allowed values are 'IRR' or 'IRT'")
}

// ValidateMerchantID validates merchant ID format (UUID)
func (v *Validator) ValidateMerchantID(merchantID string) error {
	if merchantID == "" {
		return NewValidationError("merchant_id", "merchant ID is required")
	}
	
	// UUID regex pattern
	uuidPattern := `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`
	matched, err := regexp.MatchString(uuidPattern, merchantID)
	if err != nil {
		return fmt.Errorf("error validating merchant ID: %w", err)
	}
	
	if !matched {
		return NewValidationError("merchant_id", "invalid merchant_id format. It should be a valid UUID")
	}
	
	return nil
}

// ValidateAuthority validates authority code format
func (v *Validator) ValidateAuthority(authority string) error {
	if authority == "" {
		return NewValidationError("authority", "authority is required")
	}
	
	// Authority should start with 'A' or 'S' followed by 35 alphanumeric characters
	authorityPattern := `^[AS][0-9a-zA-Z]{35}$`
	matched, err := regexp.MatchString(authorityPattern, authority)
	if err != nil {
		return fmt.Errorf("error validating authority: %w", err)
	}
	
	if !matched {
		return NewValidationError("authority", "invalid authority format. It should be a string starting with 'A' or 'S' followed by 35 alphanumeric characters")
	}
	
	return nil
}

// ValidateAmount validates transaction amount
func (v *Validator) ValidateAmount(amount int64, minAmount int64) error {
	if minAmount == 0 {
		minAmount = 10000 // Default minimum amount
	}
	
	if amount < minAmount {
		return NewValidationError("amount", fmt.Sprintf("amount must be at least %d", minAmount))
	}
	
	return nil
}

// ValidateCallbackURL validates callback URL format
func (v *Validator) ValidateCallbackURL(callbackURL string) error {
	if callbackURL == "" {
		return NewValidationError("callback_url", "callback URL is required")
	}
	
	if !strings.HasPrefix(callbackURL, "http://") && !strings.HasPrefix(callbackURL, "https://") {
		return NewValidationError("callback_url", "invalid callback URL format. It should start with http:// or https://")
	}
	
	return nil
}

// ValidateMobile validates Iranian mobile number format
func (v *Validator) ValidateMobile(mobile string) error {
	if mobile == "" {
		return nil // Mobile is optional
	}
	
	// Iranian mobile number pattern
	mobilePattern := `^09[0-9]{9}$`
	matched, err := regexp.MatchString(mobilePattern, mobile)
	if err != nil {
		return fmt.Errorf("error validating mobile: %w", err)
	}
	
	if !matched {
		return NewValidationError("mobile", "invalid mobile number format")
	}
	
	return nil
}

// ValidateEmail validates email format
func (v *Validator) ValidateEmail(email string) error {
	if email == "" {
		return nil // Email is optional
	}
	
	// Basic email pattern
	emailPattern := `^[^\s@]+@[^\s@]+\.[^\s@]+$`
	matched, err := regexp.MatchString(emailPattern, email)
	if err != nil {
		return fmt.Errorf("error validating email: %w", err)
	}
	
	if !matched {
		return NewValidationError("email", "invalid email format")
	}
	
	return nil
}

// ValidateTerminalID validates terminal ID
func (v *Validator) ValidateTerminalID(terminalID string) error {
	if terminalID == "" {
		return NewValidationError("terminal_id", "terminal ID is required")
	}
	
	return nil
}

// ValidateFilter validates transaction filter
func (v *Validator) ValidateFilter(filter string) error {
	if filter == "" {
		return nil // Filter is optional
	}
	
	validFilters := []string{"PAID", "VERIFIED", "TRASH", "ACTIVE", "REFUNDED"}
	for _, valid := range validFilters {
		if filter == valid {
			return nil
		}
	}
	
	return NewValidationError("filter", "invalid filter value")
}

// ValidateLimit validates pagination limit
func (v *Validator) ValidateLimit(limit int) error {
	if limit <= 0 {
		return NewValidationError("limit", "limit must be a positive integer")
	}
	
	return nil
}

// ValidateOffset validates pagination offset
func (v *Validator) ValidateOffset(offset int) error {
	if offset < 0 {
		return NewValidationError("offset", "offset must be a non-negative integer")
	}
	
	return nil
}

// ValidateCardPAN validates card PAN format
func (v *Validator) ValidateCardPAN(cardPAN string) error {
	if cardPAN == "" {
		return nil // Card PAN is optional
	}
	
	// 16-digit card number pattern
	cardPattern := `^[0-9]{16}$`
	matched, err := regexp.MatchString(cardPattern, cardPAN)
	if err != nil {
		return fmt.Errorf("error validating card PAN: %w", err)
	}
	
	if !matched {
		return NewValidationError("card_pan", "invalid card PAN format. It should be a 16-digit number")
	}
	
	return nil
}

// ValidateSessionID validates session ID
func (v *Validator) ValidateSessionID(sessionID string) error {
	if sessionID == "" {
		return NewValidationError("session_id", "session ID is required")
	}
	
	return nil
}

// ValidateMethod validates refund method
func (v *Validator) ValidateMethod(method string) error {
	if method == "" {
		return nil // Method is optional
	}
	
	validMethods := []string{"PAYA", "CARD"}
	for _, valid := range validMethods {
		if method == valid {
			return nil
		}
	}
	
	return NewValidationError("method", "invalid method. Allowed values are 'PAYA' or 'CARD'")
}

// ValidateReason validates refund reason
func (v *Validator) ValidateReason(reason string) error {
	if reason == "" {
		return nil // Reason is optional
	}
	
	validReasons := []string{
		"CUSTOMER_REQUEST",
		"DUPLICATE_TRANSACTION",
		"SUSPICIOUS_TRANSACTION",
		"OTHER",
	}
	
	for _, valid := range validReasons {
		if reason == valid {
			return nil
		}
	}
	
	return NewValidationError("reason", "invalid reason. Allowed values are 'CUSTOMER_REQUEST', 'DUPLICATE_TRANSACTION', 'SUSPICIOUS_TRANSACTION', or 'OTHER'")
}

// ValidateIBAN validates IBAN format (basic check)
func (v *Validator) ValidateIBAN(iban string) error {
	if iban == "" {
		return NewValidationError("iban", "IBAN is required")
	}
	
	// Basic IBAN pattern: 2 letters + 2 digits + up to 30 alphanumeric
	ibanPattern := `^[A-Z]{2}[0-9]{2}[0-9A-Z]{1,30}$`
	matched, err := regexp.MatchString(ibanPattern, iban)
	if err != nil {
		return fmt.Errorf("error validating IBAN: %w", err)
	}
	
	if !matched {
		return NewValidationError("iban", "invalid IBAN format")
	}
	
	return nil
}
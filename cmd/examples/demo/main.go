package main

import (
	"fmt"
	"log"

	"github.com/DrDesignX/zarinpal-go-sdk/pkg/zarinpal"
	"github.com/DrDesignX/zarinpal-go-sdk/pkg/zarinpal/resources"
)

func main() {
	fmt.Println("🚀 ZarinPal Go SDK Demo")
	fmt.Println("========================")

	// Demo configuration validation
	configDemo()

	// Demo validation
	validatorDemo()

	// Demo client creation (without real API calls)
	clientDemo()
}

func configDemo() {
	fmt.Println("\n📋 Configuration Demo:")

	// Test different configuration options
	config1, err := zarinpal.NewConfig(
		zarinpal.WithMerchantID("123e4567-e89b-12d3-a456-426614174000"),
		zarinpal.WithSandbox(true),
	)
	if err != nil {
		fmt.Printf("❌ Config creation failed: %v\n", err)
	} else {
		fmt.Printf("✅ Config created successfully\n")
		fmt.Printf("   Sandbox mode: %v\n", config1.Sandbox)
		fmt.Printf("   Base URL: %s\n", config1.BaseURL())
	}

	// Test validation error
	_, err = zarinpal.NewConfig(
		zarinpal.WithMerchantID(""), // This should fail
	)
	if err != nil {
		fmt.Printf("✅ Validation working: %v\n", err)
	}
}

func validatorDemo() {
	fmt.Println("\n🔍 Validation Demo:")

	validator := zarinpal.NewValidator()

	// Test amount validation
	err := validator.ValidateAmount(15000, 10000)
	if err != nil {
		fmt.Printf("❌ Amount validation failed: %v\n", err)
	} else {
		fmt.Printf("✅ Amount validation passed (15000 >= 10000)\n")
	}

	// Test invalid amount
	err = validator.ValidateAmount(5000, 10000)
	if err != nil {
		fmt.Printf("✅ Amount validation correctly failed: %v\n", err)
	}

	// Test authority validation
	err = validator.ValidateAuthority("A0000000000000000000000000006qpmlj8d")
	if err != nil {
		fmt.Printf("❌ Authority validation failed: %v\n", err)
	} else {
		fmt.Printf("✅ Authority validation passed\n")
	}

	// Test mobile validation
	err = validator.ValidateMobile("09123456789")
	if err != nil {
		fmt.Printf("❌ Mobile validation failed: %v\n", err)
	} else {
		fmt.Printf("✅ Mobile validation passed\n")
	}

	// Test email validation
	err = validator.ValidateEmail("test@example.com")
	if err != nil {
		fmt.Printf("❌ Email validation failed: %v\n", err)
	} else {
		fmt.Printf("✅ Email validation passed\n")
	}
}

func clientDemo() {
	fmt.Println("\n🌐 Client Demo:")

	config, err := zarinpal.NewConfig(
		zarinpal.WithMerchantID("123e4567-e89b-12d3-a456-426614174000"),
		zarinpal.WithSandbox(true),
	)
	if err != nil {
		log.Fatalf("Failed to create config: %v", err)
	}

	// Create client with resources
	client := zarinpal.NewClientWithResources(config)
	fmt.Printf("✅ Client created successfully\n")

	// Demo payment request structure (no actual API call)
	paymentRequest := &resources.PaymentRequest{
		Amount:      20000,
		CallbackURL: "https://example.com/callback",
		Description: "Test payment from Go SDK",
		Mobile:      "09123456789",
		Email:       "customer@example.com",
	}

	fmt.Printf("✅ Payment request structure created:\n")
	fmt.Printf("   Amount: %d IRR\n", paymentRequest.Amount)
	fmt.Printf("   Callback URL: %s\n", paymentRequest.CallbackURL)
	fmt.Printf("   Description: %s\n", paymentRequest.Description)

	// Demo other request structures
	verifyRequest := &resources.VerificationRequest{
		Amount:    20000,
		Authority: "A0000000000000000000000000006qpmlj8d",
	}
	fmt.Printf("✅ Verification request structure created for authority: %s\n", verifyRequest.Authority)

	feeRequest := &resources.FeeRequest{
		Amount:   90000,
		Currency: "IRR",
	}
	fmt.Printf("✅ Fee calculation request structure created for amount: %d %s\n", 
		feeRequest.Amount, feeRequest.Currency)

	// Show that we have all the resources available
	fmt.Printf("✅ All resources initialized:\n")
	fmt.Printf("   - Payments: %T\n", client.Payments)
	fmt.Printf("   - Verifications: %T\n", client.Verifications)
	fmt.Printf("   - Inquiries: %T\n", client.Inquiries)
	fmt.Printf("   - Refunds: %T\n", client.Refunds)
	fmt.Printf("   - Reversals: %T\n", client.Reversals)
	fmt.Printf("   - Transactions: %T\n", client.Transactions)
	fmt.Printf("   - Unverified: %T\n", client.Unverified)
	fmt.Printf("   - Fee: %T\n", client.Fee)

	fmt.Println("\n🎉 Demo completed successfully!")
	fmt.Println("ℹ️  Note: No actual API calls were made in this demo.")
	fmt.Println("ℹ️  To make real API calls, provide valid merchant credentials and use the appropriate methods.")
}
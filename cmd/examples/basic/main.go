package main

import (
	"context"
	"fmt"
	"log"

	"github.com/DrDesignX/zarinpal-go-sdk/pkg/zarinpal"
	"github.com/DrDesignX/zarinpal-go-sdk/pkg/zarinpal/resources"
)

func main() {
	// Create configuration
	config, err := zarinpal.NewConfig(
		zarinpal.WithMerchantID("your-merchant-id"),
		zarinpal.WithSandbox(true),
	)
	if err != nil {
		log.Fatalf("Failed to create config: %v", err)
	}

	// Create client with resources
	client := zarinpal.NewClientWithResources(config)

	// Example: Create a payment
	paymentExample(client)

	// Example: Calculate fee
	feeExample(client)
}

func paymentExample(client *zarinpal.ClientWithResources) {
	fmt.Println("=== Payment Example ===")

	// Create payment request
	paymentRequest := &resources.PaymentRequest{
		Amount:      20000,
		CallbackURL: "https://example.com/callback",
		Description: "Test payment from Go SDK",
		Mobile:      "09123456789",
		Email:       "customer@example.com",
	}

	// Create payment
	response, err := client.Payments.Create(context.Background(), paymentRequest)
	if err != nil {
		fmt.Printf("Payment creation failed: %v\n", err)
		return
	}

	fmt.Printf("Payment created successfully!\n")
	fmt.Printf("Authority: %s\n", response.Data.Authority)
	fmt.Printf("Code: %d\n", response.Data.Code)
	fmt.Printf("Message: %s\n", response.Data.Message)

	// Generate payment URL
	if response.Data.Authority != "" {
		paymentURL := client.Payments.GeneratePaymentURL(response.Data.Authority)
		fmt.Printf("Payment URL: %s\n", paymentURL)
	}

	// Example: Verify payment (normally done after user returns from payment gateway)
	verifyPayment(client, response.Data.Authority, paymentRequest.Amount)
}

func verifyPayment(client *zarinpal.ClientWithResources, authority string, amount int64) {
	fmt.Println("\n=== Verification Example ===")

	// Create verification request
	verifyRequest := &resources.VerificationRequest{
		Amount:    amount,
		Authority: authority,
	}

	// Verify payment
	response, err := client.Verifications.Verify(context.Background(), verifyRequest)
	if err != nil {
		fmt.Printf("Payment verification failed: %v\n", err)
		return
	}

	fmt.Printf("Payment verified!\n")
	fmt.Printf("Code: %d\n", response.Data.Code)
	fmt.Printf("Message: %s\n", response.Data.Message)
	if response.Data.RefID != 0 {
		fmt.Printf("Reference ID: %d\n", response.Data.RefID)
	}
}

func feeExample(client *zarinpal.ClientWithResources) {
	fmt.Println("\n=== Fee Calculation Example ===")

	// Create fee calculation request
	feeRequest := &resources.FeeRequest{
		Amount:   90000,
		Currency: "IRR",
	}

	// Calculate fee
	response, err := client.Fee.Calculate(context.Background(), feeRequest)
	if err != nil {
		fmt.Printf("Fee calculation failed: %v\n", err)
		return
	}

	fmt.Printf("Fee calculation successful!\n")
	fmt.Printf("Amount: %d\n", response.Data.Amount)
	fmt.Printf("Fee: %d\n", response.Data.Fee)
	fmt.Printf("Fee Type: %s\n", response.Data.FeeType)
	fmt.Printf("Suggested Amount: %d\n", response.Data.SuggestedAmount)
}
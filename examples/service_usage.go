package main

import (
	"context"
	"fmt"
	"log"

	"github.com/DrDesignX/zarinpal-go-sdk/pkg/services"
)

func main() {
	// Initialize ZarinPal service
	zarinpalService, err := services.NewZarinPalService("your-merchant-id", true) // true for sandbox
	if err != nil {
		log.Fatalf("Failed to create ZarinPal service: %v", err)
	}

	ctx := context.Background()

	// Example 1: Create payment
	paymentReq := services.PaymentRequest{
		Amount:      100000, // 100,000 Rials (now int64)
		UserID:      12345,
		Description: "Test payment",
		CallbackURL: "https://yoursite.com/callback",
	}

	paymentResp, err := zarinpalService.CreatePayment(ctx, paymentReq)
	if err != nil {
		log.Printf("Payment creation failed: %v", err)
		return
	}

	fmt.Printf("Payment created: %+v\n", paymentResp)

	// Example 2: Verify payment (after user returns from payment gateway)
	verificationReq := services.VerificationRequest{
		Authority: paymentResp.Authority,
		Amount:    100000, // Must match the original amount (now int64)
	}

	verificationResp, err := zarinpalService.VerifyPayment(ctx, verificationReq)
	if err != nil {
		log.Printf("Payment verification failed: %v", err)
		return
	}

	fmt.Printf("Payment verification: %+v\n", verificationResp)

	// Example 3: Calculate fee
	fee, err := zarinpalService.CalculateFee(ctx, 100000) // Amount as int64
	if err != nil {
		log.Printf("Fee calculation failed: %v", err)
		return
	}

	fmt.Printf("Transaction fee: %d Rials\n", fee) // fee is now int64
}

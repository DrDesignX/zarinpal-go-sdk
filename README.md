# ZarinPal Go SDK

The **ZarinPal Go SDK** provides an easy and flexible way to interact with ZarinPal's payment gateway APIs. This SDK is a complete port from Python to Go, featuring payment initiation, transaction management, refunds, reversals, and fee calculations for both live and sandbox environments.

## **Features**

- ✅ **Payment Management**: Create and verify payments with secure APIs
- ✅ **Transaction Queries**: Fetch transaction details with filtering and pagination via GraphQL
- ✅ **Refunds & Reversals**: Process refunds and reverse transactions
- ✅ **Fee Calculation**: Calculate transaction fees
- ✅ **Sandbox Support**: Easy switching between sandbox and production environments
- ✅ **Type Safety**: Strong typing with custom types and comprehensive error handling
- ✅ **Context Support**: Full context.Context support for timeouts and cancellation
- ✅ **Go Modules**: Compatible with Go modules and modern Go practices

## **Installation**

### **Requirements**:
- Go >= 1.18

### **Install**:

```bash
go get github.com/DrDesignX/zarinpal-go-sdk
```

## **Quick Start**

```go
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
        zarinpal.WithAccessToken("your-access-token"), // For GraphQL operations
        zarinpal.WithSandbox(true), // Use sandbox environment
    )
    if err != nil {
        log.Fatalf("Failed to create config: %v", err)
    }

    // Create client with all resources
    client := zarinpal.NewClientWithResources(config)

    // Create a payment
    paymentRequest := &resources.PaymentRequest{
        Amount:      20000,
        CallbackURL: "https://yourwebsite.com/callback",
        Description: "Order #1234",
        Mobile:      "09123456789",
        Email:       "customer@example.com",
    }

    response, err := client.Payments.Create(context.Background(), paymentRequest)
    if err != nil {
        log.Fatalf("Payment creation failed: %v", err)
    }

    fmt.Printf("Payment Authority: %s\n", response.Data.Authority)
    
    // Generate payment URL
    paymentURL := client.Payments.GeneratePaymentURL(response.Data.Authority)
    fmt.Printf("Payment URL: %s\n", paymentURL)
}
```

## **Configuration Options**

### **Using Functional Options**:

```go
config, err := zarinpal.NewConfig(
    zarinpal.WithMerchantID("your-merchant-id"),
    zarinpal.WithAccessToken("your-access-token"),
    zarinpal.WithSandbox(true),
)
```

### **Using Environment Variables**:

```bash
export ZARINPAL_MERCHANT_ID="your-merchant-id"
export ZARINPAL_ACCESS_TOKEN="your-access-token"
export ZARINPAL_SANDBOX="true"
```

```go
config, err := zarinpal.NewConfig(
    zarinpal.WithEnvVars(),
)
```

### **Custom HTTP Client**:

```go
httpClient := &http.Client{
    Timeout: 60 * time.Second,
}

config, _ := zarinpal.NewConfig(
    zarinpal.WithMerchantID("your-merchant-id"),
    zarinpal.WithSandbox(true),
)

client := zarinpal.NewClientWithResources(config, 
    zarinpal.WithHTTPClient(httpClient),
)
```

## **Usage Examples**

### **1. Payment Creation and Verification**:

```go
// Create payment
paymentRequest := &resources.PaymentRequest{
    Amount:      50000,
    CallbackURL: "https://yourwebsite.com/callback",
    Description: "Product purchase",
    Mobile:      "09123456789",
    Email:       "customer@example.com",
}

response, err := client.Payments.Create(ctx, paymentRequest)
if err != nil {
    // Handle error
}

// After user returns from payment gateway, verify the payment
verifyRequest := &resources.VerificationRequest{
    Amount:    50000,
    Authority: response.Data.Authority,
}

verifyResponse, err := client.Verifications.Verify(ctx, verifyRequest)
```

### **2. Fee Calculation**:

```go
feeRequest := &resources.FeeRequest{
    Amount:   90000,
    Currency: "IRR",
}

feeResponse, err := client.Fee.Calculate(ctx, feeRequest)
if err != nil {
    // Handle error
}

fmt.Printf("Transaction fee: %d\n", feeResponse.Data.Fee)
fmt.Printf("Suggested amount: %d\n", feeResponse.Data.SuggestedAmount)
```

### **3. Transaction Listing**:

```go
transactionRequest := &resources.TransactionRequest{
    TerminalID: "your-terminal-id",
    Filter:     "PAID",
    Limit:      10,
    Offset:     0,
}

transactions, err := client.Transactions.List(ctx, transactionRequest)
if err != nil {
    // Handle error
}

for _, transaction := range transactions.Data.Session {
    fmt.Printf("Transaction ID: %s, Amount: %d\n", 
        transaction.ID, transaction.Amount)
}
```

### **4. Refund Processing**:

```go
refundRequest := &resources.RefundRequest{
    SessionID:   "session-id",
    Amount:      25000,
    Description: "Partial refund",
    Method:      "CARD",
    Reason:      "CUSTOMER_REQUEST",
}

refundResponse, err := client.Refunds.Create(ctx, refundRequest)
```

### **5. Transaction Reversal**:

```go
reversalRequest := &resources.ReversalRequest{
    Authority: "authority-code",
}

reversalResponse, err := client.Reversals.Reverse(ctx, reversalRequest)
```

## **Error Handling**

The SDK provides comprehensive error handling with custom error types:

```go
response, err := client.Payments.Create(ctx, paymentRequest)
if err != nil {
    if zarinpal.IsValidationError(err) {
        fmt.Printf("Validation error: %v\n", err)
    } else if zarinpal.IsHTTPError(err) {
        statusCode := zarinpal.GetHTTPStatusCode(err)
        fmt.Printf("HTTP error %d: %v\n", statusCode, err)
    } else {
        fmt.Printf("Other error: %v\n", err)
    }
}
```

## **Testing**

Run the test suite:

```bash
go test ./tests/... -v
```

Run tests with coverage:

```bash
go test ./tests/... -v -cover
```

## **Build Instructions**

### **Build the SDK**:

```bash
go build ./pkg/zarinpal/...
```

### **Build Examples**:

```bash
go build ./cmd/examples/basic
```

### **Run Example**:

```bash
# Set your credentials
export ZARINPAL_MERCHANT_ID="your-merchant-id"
export ZARINPAL_SANDBOX="true"

# Run the example
./basic
```

## **Major Design Changes from Python Version**

### **1. Type Safety**:
- **Go**: Strong typing with custom types for amounts, authorities, etc.
- **Python**: Dynamic typing with runtime validation

### **2. Error Handling**:
- **Go**: Explicit error returns with custom error types
- **Python**: Exception-based error handling

### **3. Context Support**:
- **Go**: Built-in `context.Context` support for all API calls
- **Python**: No native context support

### **4. Package Organization**:
- **Go**: Clear package separation with interfaces and implementations
- **Python**: Class-based organization with lazy loading

### **5. Configuration**:
- **Go**: Functional options pattern with environment variable support
- **Python**: Simple class-based configuration

### **6. Resource Management**:
- **Go**: Interface-based design for easy testing and mocking
- **Python**: Property-based lazy loading

### **7. HTTP Client**:
- **Go**: Built on standard `net/http` with customizable client options
- **Python**: Uses `requests` library

## **API Compatibility**

This Go SDK maintains 100% API compatibility with the Python version:

- ✅ All payment gateway endpoints
- ✅ All GraphQL operations
- ✅ Same validation rules
- ✅ Same request/response structures
- ✅ Same error codes and messages

## **Contributing**

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Ensure all tests pass
5. Submit a pull request

## **License**

This project is licensed under the MIT License - see the LICENSE file for details.

## **Support**

For issues and questions:
- GitHub Issues: [Create an issue](https://github.com/DrDesignX/zarinpal-go-sdk/issues)
- Documentation: Check the examples in `cmd/examples/`

---

By providing native Go implementation with modern practices, this SDK enables developers to integrate ZarinPal payments efficiently while maintaining type safety and performance.
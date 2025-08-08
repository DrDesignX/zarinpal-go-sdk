package resources

// PaymentRequest represents a payment creation request
type PaymentRequest struct {
	Amount      int64    `json:"amount"`
	CallbackURL string   `json:"callback_url"`
	Description string   `json:"description,omitempty"`
	Mobile      string   `json:"mobile,omitempty"`
	Email       string   `json:"email,omitempty"`
	CardPAN     []string `json:"card_pan,omitempty"`
	ReferrerID  string   `json:"referrer_id,omitempty"`
}

// PaymentResponse represents a payment creation response
type PaymentResponse struct {
	Data   PaymentData `json:"data"`
	Errors []string    `json:"errors,omitempty"`
}

// PaymentData represents the data portion of payment response
type PaymentData struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Authority string `json:"authority"`
	FeeType   string `json:"fee_type,omitempty"`
	Fee       int64  `json:"fee,omitempty"`
}

// VerificationRequest represents a payment verification request
type VerificationRequest struct {
	Amount    int64  `json:"amount"`
	Authority string `json:"authority"`
}

// VerificationResponse represents a payment verification response
type VerificationResponse struct {
	Data   VerificationData `json:"data"`
	Errors []string         `json:"errors,omitempty"`
}

// VerificationData represents the data portion of verification response
type VerificationData struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	CardHash string `json:"card_hash,omitempty"`
	CardPAN  string `json:"card_pan,omitempty"`
	RefID    int64  `json:"ref_id,omitempty"`
	FeeType  string `json:"fee_type,omitempty"`
	Fee      int64  `json:"fee,omitempty"`
}

// InquiryRequest represents a transaction inquiry request
type InquiryRequest struct {
	Authority string `json:"authority"`
}

// InquiryResponse represents a transaction inquiry response
type InquiryResponse struct {
	Data   InquiryData `json:"data"`
	Errors []string    `json:"errors,omitempty"`
}

// InquiryData represents the data portion of inquiry response
type InquiryData struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Authority string `json:"authority"`
	Amount    int64  `json:"amount"`
	RefID     int64  `json:"ref_id,omitempty"`
}

// RefundRequest represents a refund creation request
type RefundRequest struct {
	SessionID   string `json:"session_id"`
	Amount      int64  `json:"amount"`
	Description string `json:"description,omitempty"`
	Method      string `json:"method,omitempty"`
	Reason      string `json:"reason,omitempty"`
}

// RefundResponse represents a refund creation response
type RefundResponse struct {
	Data   RefundData `json:"data"`
	Errors []string   `json:"errors,omitempty"`
}

// RefundData represents the data portion of refund response
type RefundData struct {
	Resource RefundResource `json:"resource"`
}

// RefundResource represents the refund resource data
type RefundResource struct {
	TerminalID string         `json:"terminal_id"`
	ID         string         `json:"id"`
	Amount     int64          `json:"amount"`
	Timeline   RefundTimeline `json:"timeline"`
}

// RefundTimeline represents the refund timeline data
type RefundTimeline struct {
	RefundAmount int64  `json:"refund_amount"`
	RefundTime   string `json:"refund_time"`
	RefundStatus string `json:"refund_status"`
}

// ReversalRequest represents a transaction reversal request
type ReversalRequest struct {
	Authority string `json:"authority"`
}

// ReversalResponse represents a transaction reversal response
type ReversalResponse struct {
	Data   ReversalData `json:"data"`
	Errors []string     `json:"errors,omitempty"`
}

// ReversalData represents the data portion of reversal response
type ReversalData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// TransactionRequest represents a transaction listing request
type TransactionRequest struct {
	TerminalID string `json:"terminal_id"`
	Filter     string `json:"filter,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Offset     int    `json:"offset,omitempty"`
}

// TransactionResponse represents a transaction listing response
type TransactionResponse struct {
	Data   TransactionData `json:"data"`
	Errors []string        `json:"errors,omitempty"`
}

// TransactionData represents the data portion of transaction response
type TransactionData struct {
	Session []Transaction `json:"Session"`
}

// Transaction represents a single transaction
type Transaction struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

// UnverifiedResponse represents an unverified payments response
type UnverifiedResponse struct {
	Data   UnverifiedData `json:"data"`
	Errors []string       `json:"errors,omitempty"`
}

// UnverifiedData represents the data portion of unverified response
type UnverifiedData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	// Add other fields as needed based on actual API response
}

// FeeRequest represents a fee calculation request
type FeeRequest struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency,omitempty"`
}

// FeeResponse represents a fee calculation response
type FeeResponse struct {
	Data   FeeData  `json:"data"`
	Errors []string `json:"errors,omitempty"`
}

// FeeData represents the data portion of fee response
type FeeData struct {
	Amount          int64  `json:"amount"`
	Fee             int64  `json:"fee"`
	FeeType         string `json:"fee_type"`
	SuggestedAmount int64  `json:"suggested_amount"`
	Code            int    `json:"code"`
	Message         string `json:"message"`
}

// Wage represents a wage distribution entry
type Wage struct {
	IBAN        string `json:"iban"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
}
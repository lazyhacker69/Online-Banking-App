package request

type ApplyLoanRequest struct {
	AccountID uint    `json:"account_id"`
	LoanType  string  `json:"loan_type"`
	Amount    float64 `json:"amount"`
}
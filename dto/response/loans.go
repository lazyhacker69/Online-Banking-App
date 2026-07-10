package response

type ApplyLoanResponse struct {
	Message      string  `json:"message"`
	LoanID       uint    `json:"loan_id"`
	LoanType     string  `json:"loan_type"`
	LoanAmount   float64 `json:"loan_amount"`
	InterestRate float64 `json:"interest_rate"`
}
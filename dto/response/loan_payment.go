package response

type RepayLoanResponse struct {
	Message         string  `json:"message"`
	RemainingAmount float64 `json:"remaining_amount"`
	LoanStatus      string  `json:"loan_status"`
}
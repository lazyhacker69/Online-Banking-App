package response

type DepositRespone struct {
	Message string `json:"message"`
	AccountNumber string `json:"account_number"`
	NewBalance float64 `json:"new_balance"`
}
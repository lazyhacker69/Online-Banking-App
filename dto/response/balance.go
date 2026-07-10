package response

type BalanceResponse struct {
	Balance float64 `json:"amount"`
	AccountNumber string `json:"account_number"`
}
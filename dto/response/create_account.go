package response

type CreateAccountResponse struct {
	Message string `json:"message"`
	AccountNumber string `json:"account_number"`
	AccountType string `json:"account_type"`
	Balance float64 `json:"Balance"`
}
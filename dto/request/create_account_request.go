package request

type CreateAccountRequest struct {
	CustomerID uint `json:"customer_id"`
	AccountType string `json:"account_type"`
}

package models

import "time"

type Transaction struct {
	TransactionID 	uint      `json:"transaction_id"`
	AccountID 		uint 	  `json:"account_id"`
	TransactionType string    `json:"transaction_type"`
	Amount 			float64   `json:"amount"`
	CreatedAt 		time.Time `json:"Created_at"`
}
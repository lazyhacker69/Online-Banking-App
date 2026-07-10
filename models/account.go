package models

import "time"

type Account struct {
	AccountID     uint      `gorm:"primaryKey" json:"account_id"`
	AccountNumber string    `json:"account_number"`
	CustomerID    uint      `json:"customer_id"`
	AccountType   string    `json:"account_type"`
	Balance       float64   `json:"balance"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}
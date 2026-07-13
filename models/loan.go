package models

import "time"

type Loan struct {
	LoanID          uint      `gorm:"primaryKey" json:"loan_id"`
	AccountID       uint      `json:"account_id"`
	LoanType        string    `json:"loan_type"`
	PrincipalAmount float64   `json:"principal_amount"`
	RemainingAmount float64   `json:"remaining_amount"`
	InterestRate    float64   `json:"interest_rate"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
}


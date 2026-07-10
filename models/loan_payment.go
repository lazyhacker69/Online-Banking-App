package models

import "time"

type LoanPayment struct {
	PaymentID   uint      `gorm:"primaryKey" json:"payment_id"`
	LoanID      uint      `json:"loan_id"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
}
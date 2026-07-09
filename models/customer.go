package models

import "time"

type Customer struct{
	CustomerID uint `gorm:"primaryKey" json:"customer_id"`
	FirstName string `gorm:"not null" json:"first_name"`
	LastName string `gorm:"not null" json:"last_name"`

	Phone string `gorm:"unique;not null" json:"phone"`
	Email string `gorm:"unique;not null" json:"email"`

	Address string `gorm:"not null" json:"address"`

	BranchID uint `gorm:"not null" json:"branch_id"`

	CreatedAt time.Time `json:"created_at"`
}
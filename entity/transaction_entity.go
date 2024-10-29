package entity

import (
	"github.com/google/uuid"
	"time"
)

type transaction_type string

const (
    INCOME transaction_type = "income"
    EXPENSE transaction_type = "expense"
)

type Transactions struct {
    TransactionID   uuid.UUID   `json:"transaction_id"`
    UserID          uuid.UUID   `json:"user_id"`
    AccountID       uuid.UUID   `json:"account_id"`
    CategoryID      uuid.UUID   `json:"category_id"`
    Type            string      `sql:"type:transaction_type" json:"type"`
    Amount          float64     `json:"amount"`
    Description     string      `json:"description"`
    TransactionDate time.Time   `json:"transaction_date"`
    CreatedAt       time.Time   `json:"created_at"`
    UpdatedAt       time.Time   `json:"updated_at"`
}
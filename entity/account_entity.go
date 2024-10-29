package entity

import (
	"github.com/google/uuid"
	"time"
)

type account_type string

const (
    ASSET  account_type = "asset"
    LIABILITY account_type = "liability"
)

type Accounts struct {
    AccountID   uuid.UUID `json:"account_id"`
    UserID      uuid.UUID `json:"user_id"`
    AccountName string    `json:"name"`
    AccountType string    `sql:"type:account_type" json:"account_type"`
    Balance     float64   `json:"balance"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
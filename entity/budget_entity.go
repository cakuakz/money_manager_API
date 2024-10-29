package entity

import (
	"github.com/google/uuid"
	"time"
)

type Budgets struct {
    BudgetID   uuid.UUID    `json:"budget_id"`
    UserID     uuid.UUID	`json:"user_id"`
    CategoryID uuid.UUID   	`json:"category_id"`
    Amount     float64   	`json:"amount"`
    StartDate  time.Time 	`json:"start_date"`
    EndDate    time.Time 	`json:"end_date"`
    CreatedAt  time.Time 	`json:"created_at"`
    UpdatedAt  time.Time 	`json:"updated_at"`
}
package entity

import (
	"github.com/google/uuid"
	"time"
)

type Categories struct {
    CategoryID      uuid.UUID    `json:"category_id"`
    UserID          uuid.UUID    `json:"user_id"`
    CategoryName    string    	 `json:"name"`
    Type            string    	 `sql:"type:transaction_type" json:"type"`
    CreatedAt       time.Time 	 `json:"created_at"`
    UpdatedAt       time.Time 	 `json:"updated_at"`          
}
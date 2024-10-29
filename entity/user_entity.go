package entity

import (
	"github.com/google/uuid"
	"time"
)

type Users struct {
	UserID      uuid.UUID `json:"user_id"`
	Username    string    `json:"username"`
    Fullname    string    `json:"fullname"`
    Password  	string    `json:"password"`
    CreatedAt  	time.Time `json:"created_at"`
    UpdatedAt  	time.Time `json:"updated_at"`
}
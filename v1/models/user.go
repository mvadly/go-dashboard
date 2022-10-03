package models

import "time"

type Users struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

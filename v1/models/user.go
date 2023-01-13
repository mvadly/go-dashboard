package models

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID        uuid.UUID `bson:"_id" json:"_id"`
	Name      string    `bson:"name" json:"name"`
	Username  string    `bson:"username" json:"username"`
	Password  string    `bson:"password" json:"password"`
	CreatedBy string    `bson:"created_by" json:"created_by"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

func (u *Users) SetDefaultID() {
	u.ID = uuid.New()
}

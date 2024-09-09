package models

import (
	"time"
)

type User struct {
	ID        int64     `json:"id,omitempty"`
	Username  string    `json:"username,omitempty" validate:"required,min=3,max=20"`
	Password  string    `json:"password,omitempty" validate:"required,min=8"`
	Role      Role      `json:"role,omitempty" validate:"required,oneof=BASIC ADMIN"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

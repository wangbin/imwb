package models

import (
	"time"
)

type User struct {
	Id string `json:"id,omitempty"`
	UserName string `json:"username"`
	Password string `json:"password"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	IsStaff bool `json:"is_staff"`
	IsSuperUser bool `json:"is_superuser"`
	IsActive bool `json:"is_active"`
	DateJoined time.Time `json:"date_joined"`
	LastLogin time.Time `json:"last_login"`
	Groups []string `json:"groups"`
}

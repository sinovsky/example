package domain

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Status    Status    `json:"status"`
}

type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	BannedActive   Status = "banned"
)

// DTO
type UserResponse struct {
	User    *User
	Message string
}

package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID               uuid.UUID `db:"ID" json:"id"`
	Username         string    `db:"username" json:"username"`
	FirstName        string    `db:"first_name" json:"first_name"`
	LastName         string    `db:"last_name" json:"last_name"`
	Password         string    `db:"password_hash" json:"password"`
	Email            string    `db:"email" json:"email"`
	ProfilePicture   string    `db:"profile_picture" json:"profile_picture"`
	Income           float64   `db:"income" json:"income"`
	ExpenditureLimit float64   `db:"expenditure_limit" json:"expenditure_limit"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
}

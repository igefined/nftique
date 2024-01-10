package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UID         uuid.UUID `json:"uid" db:"id"`
	Web3Address string    `json:"web3Address" db:"web3_address"`
	Username    string    `json:"username" db:"username"`
	FirstName   string    `json:"first_name" db:"first_name"`
	LastName    string    `json:"last_name" db:"last_name"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

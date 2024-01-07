package domain

import "github.com/google/uuid"

type User struct {
	UID         uuid.UUID `json:"uid"`
	Web3Address string    `json:"web3Address"`
	Username    string    `json:"username"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
}

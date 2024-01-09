package handler

type User struct {
	FirstName   string `json:"first_name" validate:"omitempty,min=2"`
	LastName    string `json:"last_name" validate:"omitempty,min=2"`
	Username    string `json:"username" validate:"omitempty,min=2"`
	Web3Address string `json:"web3Address" validate:"required,eth_addr"`
}

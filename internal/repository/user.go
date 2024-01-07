package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/igefined/nftique/internal/domain"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetByWeb3Address(ctx context.Context, web3Address string) (*domain.User, error) {
	return &domain.User{
		UID:         uuid.New(),
		Web3Address: web3Address,
		Username:    "mock_username",
		FirstName:   "mock_first_name",
		LastName:    "mock_last_name",
	}, nil
}

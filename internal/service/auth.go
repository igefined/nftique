package service

import (
	"context"

	"github.com/igefined/nftique/internal/domain"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

//go:generate mockgen -source=auth.go -package=mocks -destination=./mocks/mock_user.go UserRepository

type UserRepository interface {
	GetByWeb3Address(ctx context.Context, web3Address string) (*domain.User, error)
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, uid uuid.UUID, user *domain.User) (*domain.User, error)
	DeactivateUser(ctx context.Context, uid uuid.UUID) error
}

type AuthService struct {
	logger *zap.Logger

	userRepository UserRepository
}

func NewAuthService(logger *zap.Logger, userRepository UserRepository) *AuthService {
	return &AuthService{logger: logger, userRepository: userRepository}
}

func (s *AuthService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	return nil, nil
}

package service

import (
	"context"

	"github.com/igefined/nftique/internal/domain"

	"go.uber.org/zap"
)

type UserRepository interface {
	GetByWeb3Address(ctx context.Context, web3Address string) (*domain.User, error)
}

type AuthService struct {
	logger *zap.Logger

	userRepository UserRepository
}

func NewAuthService(logger *zap.Logger, userRepository UserRepository) *AuthService {
	return &AuthService{logger: logger, userRepository: userRepository}
}

func (s *AuthService) Register(ctx context.Context, web3Address string) (*domain.User, error) {
	return s.userRepository.GetByWeb3Address(ctx, web3Address)
}

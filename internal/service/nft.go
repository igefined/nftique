package service

import (
	"context"

	"go.uber.org/zap"
)

type NFTService struct {
	logger *zap.Logger
}

func NewNFTService(logger *zap.Logger) *NFTService {
	return &NFTService{logger: logger}
}

func (s *NFTService) ListAllAvailable(ctx context.Context) ([]string, error) {
	return []string{"nft-1", "nft-2", "nft-3"}, nil
}

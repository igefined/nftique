package handler

import (
	"context"
	"net/http"

	"github.com/igefined/nftique/internal/domain"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type NFTService interface {
	ListAllAvailable(ctx context.Context) ([]*domain.NFT, error)
}

type NFTHandler struct {
	logger  *zap.Logger
	service NFTService
}

func NewNFTHandler(logger *zap.Logger, service NFTService) *NFTHandler {
	return &NFTHandler{logger: logger, service: service}
}

func (n NFTHandler) ListAllNFTs(ctx *fiber.Ctx) error {
	NFTs, err := n.service.ListAllAvailable(ctx.Context())
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(NFTs)
}

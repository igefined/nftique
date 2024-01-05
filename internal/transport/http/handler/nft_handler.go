package handler

import (
	"github.com/igefined/nftique/internal/domain"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type NFTService interface {
	ListAllAvailable(ctx *fiber.Ctx) ([]*domain.NFT, error)
}

type NFTHandler struct {
	logger  *zap.Logger
	service NFTService
}

func NewNFTHandler(logger *zap.Logger, service NFTService) *NFTHandler {
	return &NFTHandler{logger: logger, service: service}
}

func (n NFTHandler) ListAllNFTs(ctx *fiber.Ctx) error {
	nfts, err := n.service.ListAllAvailable(ctx)
	if err != nil {
		return err
	}

	return ctx.JSON(nfts)
}

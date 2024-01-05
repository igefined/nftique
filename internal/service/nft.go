package service

import (
	"math/rand"
	"time"

	"github.com/igefined/nftique/internal/domain"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type NFTService struct {
	logger *zap.Logger
}

func NewNFTService(logger *zap.Logger) *NFTService {
	return &NFTService{logger: logger}
}

func (s *NFTService) ListAllAvailable(ctx *fiber.Ctx) ([]*domain.NFT, error) {
	size := rand.Intn(10)
	mockNFTs := make([]*domain.NFT, size)

	for i := range mockNFTs {
		mockNFTs[i] = &domain.NFT{
			Creator:   "0x3E1D0cEd18A4454BA390b8F540682c718748b0e5",
			Owner:     "0x3E1D0cEd18A4454BA390b8F540682c718748b0e5",
			CreatedAt: time.Now().Add(-time.Minute * 2),
			Img:       "https://nftique.s3.amazonaws.com/63f902d79a33f77a446ce12c_qIsIXihKZeVDop6AZWt1j6gxOnYZ_oGfr09PzlJDL_H4YWasvDrNuTXK8Qrmh0oL6ppWI3RaGU5vMif2gNwO38UdWWei4eZCNhbfdrVlm5qHV3zVYIk6qtBuFvdkoo0HexhmSmvn.jpeg", //nolint:lll
		}
	}

	return mockNFTs, nil
}

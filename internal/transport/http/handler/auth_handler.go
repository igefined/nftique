package handler

import (
	"github.com/igefined/nftique/internal/domain"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type AuthHandler struct {
	logger *zap.Logger
}

func NewAuthHandler(logger *zap.Logger) *AuthHandler {
	return &AuthHandler{logger: logger}
}

func (n AuthHandler) SignIn(ctx *fiber.Ctx) error {
	return ctx.JSON(&domain.User{
		UID:         "user_uid",
		Web3Address: "0x3E1D0cEd18A4454BA390b8F540682c718748b0e5",
		Username:    "igefined",
		FirstName:   "Igor",
		LastName:    "Pomazkov",
	})
}

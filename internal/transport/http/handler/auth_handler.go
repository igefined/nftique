package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/igefined/nftique/internal/domain"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type AuthService interface {
	Register(ctx context.Context, web3Address string) (*domain.User, error)
}

type AuthHandler struct {
	logger *zap.Logger

	service AuthService
}

func NewAuthHandler(logger *zap.Logger, service AuthService) *AuthHandler {
	return &AuthHandler{logger: logger, service: service}
}

func (n AuthHandler) SignUp(ctx *fiber.Ctx) error {
	var req *GetWeb3AddressRequest
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		return err
	}

	user, err := n.service.Register(ctx.Context(), req.Web3Address)
	if err != nil {
		n.logger.Error("Sign up", zap.Error(err))
		return err
	}

	return ctx.Status(http.StatusOK).JSON(user)
}

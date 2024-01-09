package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/igefined/nftique/internal/domain"
	apiErr "github.com/igefined/nftique/internal/transport/http/error"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type AuthService interface {
	Register(ctx context.Context, web3Address string) (*domain.User, error)
}

type AuthHandler struct {
	logger    *zap.Logger
	validator *validator.Validate

	service AuthService
}

func NewAuthHandler(
	logger *zap.Logger,
	validator *validator.Validate,
	service AuthService,
) *AuthHandler {
	return &AuthHandler{logger: logger, validator: validator, service: service}
}

func (n AuthHandler) SignUp(ctx *fiber.Ctx) error {
	var req *User
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		return apiErr.RespondBadRequest(ctx, err)
	}

	if err := n.validator.StructCtx(ctx.Context(), req); err != nil {
		return apiErr.RespondBadRequest(ctx, err)
	}

	user, err := n.service.Register(ctx.Context(), req.Web3Address)
	if err != nil {
		n.logger.Error("Sign up", zap.Error(err))
		return err
	}

	return ctx.Status(http.StatusOK).JSON(user)
}

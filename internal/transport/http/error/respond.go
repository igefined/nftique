package error

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Respond(ctx *fiber.Ctx, err error) error {
	var apiErr *ApiError
	if errors.As(err, &apiErr) {
		ctx.Status(apiErr.StatusCode)
		return ctx.JSON(apiErr)
	}

	ctx.Status(http.StatusInternalServerError)
	return ctx.JSON(ErrInternal)
}

package http

import (
	"github.com/friendsofgo/errors"
	"github.com/gofiber/fiber/v2"
)

var errorHandler = func(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Send custom error page
	response := newResponse(nil, "error", e.Error())
	err = ctx.Status(code).JSON(response)
	if err != nil {
		// In case the SendFile fails
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	// Return from handler
	return nil
}

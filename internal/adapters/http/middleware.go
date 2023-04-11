package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func middlewareApply(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: fmt.Sprintf("%s,%s,%s,%s,%s,%s", fiber.MethodGet, fiber.MethodHead, fiber.MethodPut, fiber.MethodPatch, fiber.MethodPost, fiber.MethodDelete),
	}))

	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
	}))
}

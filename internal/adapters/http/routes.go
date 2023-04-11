package http

import "github.com/gofiber/fiber/v2"

var registerRoutes = func (router *fiber.App, controller TaskController)  {
	api := router.Group("api")
	v1 := api.Group("v1")
	v1.Get("/task/:id", controller.GetById)
	v1.Post("/task", controller.Create)
	// v1.Get("/task/:id", func(c *fiber.Ctx) error {return nil})
	// v1.Put("/task/:id", func(c *fiber.Ctx) error {return nil})
	// v1.Delete("/task/:id", func(c *fiber.Ctx) error {return nil})
}
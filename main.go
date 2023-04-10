package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saeedmdd/go-hexa/cmd"
)


func hello(c *fiber.Ctx) error {
	return c.JSON("helsloao")
}


func main(){
	app := fiber.New()


	app.Get("/", hello)
	
	app.Listen(":8080")
	cmd.Execute()
}
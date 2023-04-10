package main

//go:generate sqlboiler --wipe --add-soft-deletes psql -o internal/backend/adapters/models/sqlboiler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/saeedmdd/go-hexa/internal/adapters/models"
	"github.com/saeedmdd/go-hexa/utils"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func hello(c *fiber.Ctx) error {
	return c.JSON("helsloao")
}

func main() {
	// app := fiber.New()
	// app.Get("/", hello)
	task := &models.Task{Title: null.String{String: "asir shodim"}, Description: null.NewString("keili", true)}
	db, _ := utils.PostgresConnection("127.0.0.1", "user", "user", "5432", "hexagonal", "disable")
	err := task.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		panic(err)
	}
	// app.Listen(":8080")
	// cmd.Execute()
}

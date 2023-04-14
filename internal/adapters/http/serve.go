package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/saeedmdd/go-hexa/internal/adapters/repository"
	"github.com/saeedmdd/go-hexa/internal/core/usecase"
	"github.com/saeedmdd/go-hexa/pkg/config"
)

var BootTaskController TaskController

func Init() {
	conn, err := databaseConnection()
	if err != nil {
		panic(err)
	}

	taskRepository := repository.NewTaskRepository(conn)
	taskUseCase := usecase.NewTaskHandler(taskRepository)
	taskController := NewTaskHttpController(taskUseCase)
	app := fiber.New(fiber.Config{
		AppName:      config.AppConfig.App.AppName,
		ErrorHandler: errorHandler,
	})
	middlewareApply(app)
	registerRoutes(app, taskController)

	go log.Fatal(app.Listen(":" + config.AppConfig.App.AppPort))
}

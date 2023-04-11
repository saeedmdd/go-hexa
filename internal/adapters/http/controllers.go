package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saeedmdd/go-hexa/internal/core/models"
	"github.com/saeedmdd/go-hexa/internal/core/ports/outer"
)

type TaskController interface {
	// All(ctx fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	// Update(ctx fiber.Ctx) error
	// Delete(ctx fiber.Ctx) error
}

type taskHttpController struct {
	taskUseCase outer.TaskUseCase
}

func NewTaskHttpController(taskUseCase outer.TaskUseCase) TaskController {
	return &taskHttpController{taskUseCase: taskUseCase}
}

func (t *taskHttpController) GetById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "param is not valid")
	}
	task, err := t.taskUseCase.GetTask(ctx.Context(), id)
	if task == nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.JSON(task)
}

func (t *taskHttpController) Create(ctx *fiber.Ctx) error {
	task := new(models.TaskModel)
	err := ctx.BodyParser(task)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	id, err := t.taskUseCase.SaveTask(ctx.Context(), *task)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.JSON(id)
}

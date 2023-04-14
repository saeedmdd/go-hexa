package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saeedmdd/go-hexa/internal/core/models"
	"github.com/saeedmdd/go-hexa/internal/core/ports/outer"
)

type TaskController interface {
	All(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
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
	response := newResponse(task, "ok", "task with "+ctx.Params("id"))
	return ctx.JSON(response)
}

func (t *taskHttpController) Create(ctx *fiber.Ctx) error {
	task := new(models.TaskModel)
	err := ctx.BodyParser(task)

	errors := ValidateTask(*task)
	if errors != nil {
		return ctx.Status(fiber.StatusCreated).JSON(errors)
	}
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	task, err = t.taskUseCase.SaveTask(ctx.Context(), *task)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	response := newResponse(*task, "ok", "new task created")
	return ctx.JSON(response)
}

func (t *taskHttpController) All(ctx *fiber.Ctx) error {
	tasks, err := t.taskUseCase.AllTasks(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	response := newResponse(tasks, "ok", "all tasks")
	return ctx.JSON(response)
}

func (t *taskHttpController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		response := newResponse(err, "error", "opps")
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	task := new(models.TaskModel)
	ctx.BodyParser(task)
	errors := ValidateTask(*task)
	if errors != nil {
		response := newResponse(errors, "error", "opps")
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	if err != nil {
		response := newResponse(err, "error", "opps")
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	var taskCulomns map[string]interface{}
	ctx.BodyParser(&taskCulomns)
	task, err = t.taskUseCase.UpdateTask(ctx.Context(), id, taskCulomns)
	if err != nil {
		response := newResponse(err, "error", "opps")
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	response := newResponse(task, "ok", "updated")
	return ctx.JSON(response)
}

func (t *taskHttpController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		response := newResponse(err, "error", "opps")
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	result, err := t.taskUseCase.DeleteTask(ctx.Context(), id)

	if err != nil {
		response := newResponse(err, "error", "opps")
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	if result {
		response := newResponse(nil, "ok", "deleted")
		return ctx.JSON(response)
	}
	response := newResponse(nil, "error", "not found")
	return ctx.Status(fiber.StatusNotFound).JSON(response)
}

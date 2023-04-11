package usecase

import (
	"context"

	"github.com/saeedmdd/go-hexa/internal/core/models"
	"github.com/saeedmdd/go-hexa/internal/core/ports/dependency"
	"github.com/saeedmdd/go-hexa/internal/core/ports/outer"
)

type TaskHandler struct {
	TaskRepositoryDependency dependency.TaskRepositoryDependency
}

func NewTaskHandler(taskRepositoryDepenndency dependency.TaskRepositoryDependency) outer.TaskUseCase {
	return &TaskHandler{TaskRepositoryDependency: taskRepositoryDepenndency}
}

func (r *TaskHandler) SaveTask(ctx context.Context, task models.TaskModel) (int, error) {
	return r.TaskRepositoryDependency.Save(ctx, task)
}

func (r *TaskHandler) GetTask(ctx context.Context, id int) (*models.TaskModel, error) {
	return r.TaskRepositoryDependency.Get(ctx, id)
}

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

func (r *TaskHandler) SaveTask(ctx context.Context, task models.TaskModel) (*models.TaskModel, error) {
	return r.TaskRepositoryDependency.Save(ctx, task)
}

func (r *TaskHandler) GetTask(ctx context.Context, id int) (*models.TaskModel, error) {
	return r.TaskRepositoryDependency.Get(ctx, id)
}

func (r *TaskHandler) AllTasks(ctx context.Context) ([]*models.TaskModel, error) {
	return r.TaskRepositoryDependency.All(ctx)
}

func (r *TaskHandler) UpdateTask(ctx context.Context, id int, columns map[string]interface{}) (*models.TaskModel, error) {
	return r.TaskRepositoryDependency.Update(ctx, id, columns)
}

func (r *TaskHandler) DeleteTask(ctx context.Context, id int) (bool, error) {
	return r.TaskRepositoryDependency.Delete(ctx, id)
}

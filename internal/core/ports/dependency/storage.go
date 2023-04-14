package dependency

import (
	"context"

	"github.com/saeedmdd/go-hexa/internal/core/models"
)

type TaskRepositoryDependency interface {
	All(ctx context.Context) ([]*models.TaskModel, error)
	Save(ctx context.Context, task models.TaskModel) (*models.TaskModel, error)
	Get(ctx context.Context, id int) (*models.TaskModel, error)
	Update(ctx context.Context, id int, columns map[string]interface{}) (*models.TaskModel, error)
	Delete(ctx context.Context, id int) (bool, error)
}

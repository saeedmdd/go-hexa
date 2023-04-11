package dependency

import (
	"context"

	"github.com/saeedmdd/go-hexa/internal/core/models"
)

type TaskRepositoryDependency interface {
	Save(ctx context.Context, task models.TaskModel) (int, error)
	Get(ctx context.Context, id int) (*models.TaskModel, error)
}

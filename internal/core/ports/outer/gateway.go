package outer

import (
	"context"

	"github.com/saeedmdd/go-hexa/internal/core/models"
)

type TaskUseCase interface {
	SaveTask(ctx context.Context, task models.TaskModel) (int, error)
	GetTask(ctx context.Context, id int) (*models.TaskModel, error)
}

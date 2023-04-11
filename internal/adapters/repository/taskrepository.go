package repository

import (
	"context"
	"database/sql"
	"fmt"

	boilerModel "github.com/saeedmdd/go-hexa/internal/adapters/models"
	"github.com/saeedmdd/go-hexa/internal/core/models"
	"github.com/saeedmdd/go-hexa/internal/core/ports/dependency"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) dependency.TaskRepositoryDependency {
	return &TaskRepository{db: db}
}

func (t *TaskRepository) Save(ctx context.Context, task models.TaskModel) (int, error) {
	data := boilerModel.Task{
		Title:       null.NewString(task.Title, true),
		Description: null.NewString(task.Description, true),
		Color:       null.NewString(task.Color, true),
		StartsAt:    null.NewTime(task.StartsAt, true),
		DoneAt:      null.NewTime(task.DoneAt, true),
	}
	err := data.Insert(ctx, t.db, boil.Infer())

	return data.ID, err
}
func (t *TaskRepository) Get(ctx context.Context, id int) (*models.TaskModel, error) {
	task, err := boilerModel.Tasks(boilerModel.TaskWhere.ID.EQ(id)).One(ctx, t.db)
	if err != nil {
		fmt.Println(err)
	}
	if task == nil {
		return nil, err
	}
	return &models.TaskModel{
		ID:          task.ID,
		Title:       task.Title.String,
		Description: task.Description.String,
		Color:       task.Color.String,
		StartsAt:    task.StartsAt.Time,
		DoneAt:      task.DoneAt.Time,
		CreatedAt:   task.CreatedAt.Time,
		UpdatedAt:   task.UpdatedAt.Time,
	}, err
}

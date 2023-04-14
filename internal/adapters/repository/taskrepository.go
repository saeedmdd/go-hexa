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

func (t *TaskRepository) Save(ctx context.Context, task models.TaskModel) (*models.TaskModel, error) {
	data := boilerModel.Task{
		Title:       null.NewString(task.Title, true),
		Description: null.NewString(task.Description, true),
		Color:       null.NewString(task.Color, true),
		StartsAt:    null.NewTime(task.StartsAt, true),
		DoneAt:      null.NewTime(task.DoneAt, true),
	}
	err := data.Insert(ctx, t.db, boil.Infer())
	task.ID = data.ID
	task.Color = data.Color.String
	task.DoneAt = data.DoneAt.Time
	task.CreatedAt = data.CreatedAt.Time
	task.UpdatedAt = data.UpdatedAt.Time
	task.DeletedAt = data.DeletedAt.Time
	return &task, err
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

func (t *TaskRepository) All(ctx context.Context) (tasks []*models.TaskModel, err error) {
	allTasks, err := boilerModel.Tasks().All(ctx, t.db)
	for _, task := range allTasks {
		taskModel := new(models.TaskModel)
		taskModel.ID = task.ID
		taskModel.Title = task.Title.String
		taskModel.Description = task.Description.String
		taskModel.Color = task.Color.String
		taskModel.StartsAt = task.StartsAt.Time
		taskModel.DoneAt = task.DoneAt.Time
		taskModel.CreatedAt = task.CreatedAt.Time
		taskModel.UpdatedAt = task.UpdatedAt.Time
		taskModel.DeletedAt = task.DeletedAt.Time
		tasks = append(tasks, taskModel)
	}
	return tasks, err
}

func (t *TaskRepository) Update(ctx context.Context, id int, columns map[string]interface{}) (*models.TaskModel, error) {
	_, err := boilerModel.Tasks(boilerModel.TaskWhere.ID.EQ(id)).UpdateAll(ctx, t.db, columns)
	taskModel, err := boilerModel.Tasks(boilerModel.TaskWhere.ID.EQ(id)).One(ctx, t.db)
	if err != nil {
		return nil, err
	}
	return &models.TaskModel{
		ID:          taskModel.ID,
		Title:       taskModel.Title.String,
		Description: taskModel.Description.String,
		Color:       taskModel.Color.String,
		StartsAt:    taskModel.StartsAt.Time,
		DoneAt:      taskModel.DoneAt.Time,
		CreatedAt:   taskModel.CreatedAt.Time,
		UpdatedAt:   taskModel.UpdatedAt.Time,
		DeletedAt:   taskModel.DeletedAt.Time,
	}, nil
}

func (t *TaskRepository) Delete(ctx context.Context, id int) (bool, error) {
	rows, err := boilerModel.Tasks(boilerModel.TaskWhere.ID.EQ(id)).DeleteAll(ctx, t.db)

	if rows == 0 {
		return false, err
	}
	return true, err
}

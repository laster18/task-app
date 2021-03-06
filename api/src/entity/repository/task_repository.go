package repository

import "task-api/src/entity/model"

type TaskRepository interface {
	Create(tx Transaction, task *model.Task) (int64, error)
	FindByID(tx Transaction, id int) (*model.Task, error)
	Save(tx Transaction, task *model.Task) (int64, error)
	FetchByProjectID(tx Transaction, projectID int) ([]*model.Task, error)
	Delete(tx Transaction, taskID int) error
}

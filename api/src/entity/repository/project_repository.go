package repository

import "task-api/src/entity/model"

type ProjectRepository interface {
	Create(tx Transaction, project *model.Project) (int64, error)
	AddUser(tx Transaction, userID int64, projectID int64, role string) (int64, error)
	FindByUserID(tx Transaction, userID int64) ([]*model.ProjectResult, error)
	RoleByProjectID(tx Transaction, userID int64, projectID int) (string, error)
	Delete(tx Transaction, projectID int) error
}

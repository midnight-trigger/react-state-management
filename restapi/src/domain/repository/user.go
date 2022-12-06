package repository

import "restapi/src/infrastructure/entity"

type ITaskRepository interface {
	GetTasks() ([]*entity.Task, error)
	FindByID(id string) (*entity.Task, error)
	Create(task *entity.Task) (*entity.Task, error)
	// Update(user *entity.User) (*entity.User, error)
	// Delete(id string) error
}

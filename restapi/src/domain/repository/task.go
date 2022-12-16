package repository

import "restapi/src/infrastructure/entity"

type ITaskRepository interface {
	GetTasks() ([]*entity.Task, error)
	GetTaskById(string) (*entity.Task, error)
	CreateTask(*entity.Task) (*entity.Task, error)
	UpdateTask(*entity.Task) (*entity.Task, error)
	DeleteTask(string) error
}

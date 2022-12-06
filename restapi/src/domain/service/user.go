package service

import (
	"restapi/src/common/dto"
	"restapi/src/domain/repository"
	"restapi/src/infrastructure/entity"
)

type ITaskService interface {
	GetTasks() ([]*dto.Task, error)
	FindByID(id string) (*dto.Task, error)
	Create(Task *dto.Task) (*dto.Task, error)
	// Update(Task *dto.Task) (*dto.Task, error)
	// Delete(id string) error
}

type TaskService struct {
	repository.ITaskRepository
}

func NewTaskService(repo repository.ITaskRepository) ITaskService {
	return &TaskService{repo}
}

func (s *TaskService) GetTasks() ([]*dto.Task, error) {
	getTasks, err := s.ITaskRepository.GetTasks()
	if err != nil {
		return nil, err
	}

	var tasks []*dto.Task
	for _, v := range getTasks {
		tasks = append(tasks, s.convertTo(v))
	}
	return tasks, nil
}

func (s *TaskService) FindByID(id string) (*dto.Task, error) {
	task, err := s.ITaskRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return s.convertTo(task), nil
}

func (s *TaskService) Create(Task *dto.Task) (*dto.Task, error) {
	u := s.convertFrom(Task)
	createdTask, err := s.ITaskRepository.Create(u)
	if err != nil {
		return nil, err
	}
	return s.convertTo(createdTask), nil
}

// func (s *TaskService) Update(Task *dto.Task) (*dto.Task, error) {
// 	u := s.convertFrom(Task)
// 	updatedTask, err := s.ITaskRepository.Update(u)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return s.convertTo(updatedTask), nil
// }

// func (s *TaskService) Delete(id string) error {
// 	return s.ITaskRepository.Delete(id)
// }

func (s *TaskService) convertTo(Task *entity.Task) *dto.Task {
	return dto.NewTaskModel(Task.ID, Task.Title, Task.TagNumber, Task.Tag, Task.CreatedAt, Task.UpdatedAt)
}

func (s *TaskService) convertFrom(Task *dto.Task) *entity.Task {
	return entity.NewTask(Task.ID, Task.Title, Task.TagNumber, Task.Tag, Task.CreatedAt, Task.UpdatedAt)
}

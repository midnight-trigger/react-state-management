package service

import (
	"restapi/src/common/dto"
	"restapi/src/domain/repository"
	"restapi/src/infrastructure/entity"
)

type ITaskService interface {
	GetTasks() ([]*dto.Task, error)
	CreateTask(task *dto.Task) (*dto.Task, error)
	UpdateTask(task *dto.Task) (*dto.Task, error)
	DeleteTask(id string) error
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

func (s *TaskService) CreateTask(task *dto.Task) (*dto.Task, error) {
	t := s.convertFrom(task)
	createdTask, err := s.ITaskRepository.CreateTask(t)
	if err != nil {
		return nil, err
	}
	return s.convertTo(createdTask), nil
}

func (s *TaskService) UpdateTask(task *dto.Task) (*dto.Task, error) {
	t := s.convertFrom(task)
	UupdatedTask, err := s.ITaskRepository.UpdateTask(t)
	if err != nil {
		return nil, err
	}
	return s.convertTo(UupdatedTask), nil
}

func (s *TaskService) DeleteTask(id string) error {
	if _, err := s.ITaskRepository.GetTaskById(id); err != nil {
		return err
	}
	if err := s.ITaskRepository.DeleteTask(id); err != nil {
		return err
	}
	return nil
}

func (s *TaskService) convertTo(task *entity.Task) *dto.Task {
	return dto.NewTaskModel(task.ID, task.Title, task.Tag.Name, task.TagID, task.CreatedAt, task.UpdatedAt)
}

func (s *TaskService) convertFrom(task *dto.Task) *entity.Task {
	return entity.NewTask(task.ID, task.Title, task.TagID, task.CreatedAt, task.UpdatedAt)
}

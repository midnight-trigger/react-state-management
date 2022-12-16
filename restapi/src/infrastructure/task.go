package infrastructure

import (
	"restapi/src/domain/repository"
	"restapi/src/infrastructure/entity"

	"gorm.io/gorm"
)

type TaskRepository struct {
	*gorm.DB
}

func NewTaskRepository(db *gorm.DB) repository.ITaskRepository {
	return &TaskRepository{db}
}

func (r *TaskRepository) GetTasks() ([]*entity.Task, error) {
	var tasks []*entity.Task
	if err := r.Model(&entity.Task{}).Joins("Tag").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) GetTaskById(id string) (*entity.Task, error) {
	var task *entity.Task
	if err := r.Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepository) CreateTask(task *entity.Task) (*entity.Task, error) {
	if err := r.Create(&task).Error; err != nil {
		return nil, err
	}
	var tag entity.Tag
	r.Where("id = ?", task.TagID).First(&tag)
	task.Tag = tag
	return task, nil
}

func (r *TaskRepository) UpdateTask(task *entity.Task) (*entity.Task, error) {
	if err := r.Model(&task).Updates(map[string]interface{}{"Title": task.Title, "TagID": task.TagID}).Error; err != nil {
		return nil, err
	}
	if err := r.Model(&entity.Task{}).Joins("Tag").Find(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepository) DeleteTask(id string) error {
	if err := r.Delete(&entity.Task{ID: id}).Error; err != nil {
		return err
	}
	return nil
}

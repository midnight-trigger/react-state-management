package entity

import "time"

type Task struct {
	ID        string    `gorm:"column:id"`
	Title     string    `gorm:"column:title"`
	TagID     int       `gorm:"column:tag_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Tag       Tag
}

func NewTask(id, title string, tagId int, createdAt, updatedAt time.Time) *Task {
	return &Task{
		ID:        id,
		Title:     title,
		TagID:     tagId,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

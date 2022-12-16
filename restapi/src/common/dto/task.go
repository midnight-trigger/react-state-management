package dto

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        string    `json:"id" form:"id"`
	Title     string    `json:"title" form:"title" validate:"required"`
	TagID     int       `json:"tag_id" form:"tag_id" validate:"required"`
	TagName   string    `json:"tag_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTaskModel(id, title, tagName string, tagId int, createdAt, updatedAt time.Time) *Task {
	if id == "" {
		id = uuid.NewString()
	}
	return &Task{
		ID:        id,
		Title:     title,
		TagID:     tagId,
		TagName:   tagName,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

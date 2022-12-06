package entity

import "time"

type Task struct {
	ID, Title, TagNumber string
	Tag                  int
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func NewTask(id, title, tagNumber string, tag int, createdAt, updatedAt time.Time) *Task {
	return &Task{
		ID:        id,
		Title:     title,
		TagNumber: tagNumber,
		Tag:       tag,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

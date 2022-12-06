package dto

import "time"

// type UserID struct {
// 	ID string `query:"id" form:"id" validate:"required"`
// }

type Task struct {
	ID        string    `json:"id" form:"id" validate:"required"`
	Title     string    `json:"title" form:"title" validate:"required"`
	Tag       int       `json:"tag" form:"tag" validate:"required,max=32"`
	TagNumber string    `json:"tag_number" form:"tag_number" validate:"required,max=32"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTaskModel(id, title, tagNumber string, tag int, createdAt, updatedAt time.Time) *Task {
	return &Task{
		ID:        id,
		Title:     title,
		Tag:       tag,
		TagNumber: tagNumber,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

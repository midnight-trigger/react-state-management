package dto

import "time"

type Tag struct {
	ID        int       `json:"id" form:"id" validate:"required"`
	Name      string    `json:"name" form:"name" validate:"required,max=20"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTagModel(id int, name string, createdAt, updatedAt time.Time) *Tag {
	return &Tag{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

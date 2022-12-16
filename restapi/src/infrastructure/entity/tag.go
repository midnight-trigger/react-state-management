package entity

import "time"

type Tag struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTag(id int, name string) *Tag {
	return &Tag{
		ID:   id,
		Name: name,
	}
}

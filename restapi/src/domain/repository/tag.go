package repository

import "restapi/src/infrastructure/entity"

type ITagRepository interface {
	GetTags() ([]*entity.Tag, error)
	GetTagById(int) (*entity.Tag, error)
	CreateTag(*entity.Tag) (*entity.Tag, error)
	UpdateTag(*entity.Tag) (*entity.Tag, error)
	DeleteTag(int) error
}

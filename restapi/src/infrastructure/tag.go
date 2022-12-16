package infrastructure

import (
	"restapi/src/domain/repository"
	"restapi/src/infrastructure/entity"

	"gorm.io/gorm"
)

type TagRepository struct {
	*gorm.DB
}

func NewTagRepository(db *gorm.DB) repository.ITagRepository {
	return &TagRepository{db}
}

func (r *TagRepository) GetTags() ([]*entity.Tag, error) {
	var tags []*entity.Tag
	if err := r.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *TagRepository) GetTagById(id int) (*entity.Tag, error) {
	var tag *entity.Tag
	if err := r.Where("id = ?", id).First(&tag).Error; err != nil {
		return nil, err
	}
	return tag, nil
}

func (r *TagRepository) CreateTag(tag *entity.Tag) (*entity.Tag, error) {
	if err := r.Create(&tag).Error; err != nil {
		return nil, err
	}
	return tag, nil
}

func (r *TagRepository) UpdateTag(tag *entity.Tag) (*entity.Tag, error) {
	if err := r.Model(&tag).Update("name", tag.Name).Error; err != nil {
		return nil, err
	}
	return tag, nil
}

func (r *TagRepository) DeleteTag(id int) error {
	if err := r.Delete(&entity.Tag{ID: id}).Error; err != nil {
		return err
	}
	return nil
}

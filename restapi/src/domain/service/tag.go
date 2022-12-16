package service

import (
	"restapi/src/common/dto"
	"restapi/src/domain/repository"
	"restapi/src/infrastructure/entity"
)

type ITagService interface {
	GetTags() ([]*dto.Tag, error)
	CreateTag(*dto.Tag) (*dto.Tag, error)
	UpdateTag(*dto.Tag) (*dto.Tag, error)
	DeleteTag(int) error
}

type TagService struct {
	repository.ITagRepository
}

func NewTagService(repo repository.ITagRepository) ITagService {
	return &TagService{repo}
}

func (s *TagService) GetTags() ([]*dto.Tag, error) {
	getTags, err := s.ITagRepository.GetTags()
	if err != nil {
		return nil, err
	}

	var tags []*dto.Tag
	for _, v := range getTags {
		tags = append(tags, s.convertTo(v))
	}
	return tags, nil
}

func (s *TagService) CreateTag(tag *dto.Tag) (*dto.Tag, error) {
	t := s.convertFrom(tag)
	createdTag, err := s.ITagRepository.CreateTag(t)
	if err != nil {
		return nil, err
	}
	return s.convertTo(createdTag), nil
}

func (s *TagService) UpdateTag(tag *dto.Tag) (*dto.Tag, error) {
	t := s.convertFrom(tag)
	createdTag, err := s.ITagRepository.UpdateTag(t)
	if err != nil {
		return nil, err
	}
	return s.convertTo(createdTag), nil
}

func (s *TagService) DeleteTag(id int) error {
	if _, err := s.ITagRepository.GetTagById(id); err != nil {
		return err
	}
	if err := s.ITagRepository.DeleteTag(id); err != nil {
		return err
	}
	return nil
}

func (s *TagService) convertTo(tag *entity.Tag) *dto.Tag {
	return dto.NewTagModel(tag.ID, tag.Name, tag.CreatedAt, tag.UpdatedAt)
}

func (s *TagService) convertFrom(tag *dto.Tag) *entity.Tag {
	return entity.NewTag(tag.ID, tag.Name)
}

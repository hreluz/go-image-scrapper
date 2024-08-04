package servicetag

import (
	tagModel "github.com/hreluz/images-scrapper/pkg/html_processer/tag"
	repotag "github.com/hreluz/images-scrapper/pkg/repositories/tag"
)

type TagService struct {
	TagRepository repotag.TagRepository
}

func (s *TagService) GetAll() ([]*tagModel.Tag, error) {
	tags, err := s.TagRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (s *TagService) Save(t *tagModel.Tag) error {
	err := s.TagRepository.Save(t)

	if err != nil {
		return err
	}

	return nil
}

func (s *TagService) SaveAll(tags []*tagModel.Tag) error {
	err := s.TagRepository.SaveAll(tags)

	if err != nil {
		return err
	}

	return nil
}

func (s *TagService) GetById(id int) (*tagModel.Tag, error) {
	tag, err := s.TagRepository.GetById(id)

	if err != nil {
		return nil, err
	}

	return tag, nil
}

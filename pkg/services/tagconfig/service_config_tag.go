package servicetagconfig

import (
	tagModel "github.com/hreluz/images-scrapper/pkg/html_processer/tag"
	tR "github.com/hreluz/images-scrapper/pkg/repositories/tagconfig"
)

type TagConfigService struct {
	TagConfigRepository tR.TagConfigFileRepository
}

func (s *TagConfigService) GetAll() ([]*tagModel.TagConfig, error) {
	tags, err := s.TagConfigRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (s *TagConfigService) Save(t *tagModel.TagConfig) error {
	err := s.TagConfigRepository.Save(t)

	if err != nil {
		return err
	}

	return nil
}

func (s *TagConfigService) SaveAll(tagConfigs []*tagModel.TagConfig) error {
	err := s.TagConfigRepository.SaveAll(tagConfigs)

	if err != nil {
		return err
	}

	return nil
}

func (s *TagConfigService) GetById(id int) (*tagModel.TagConfig, error) {
	tc, err := s.TagConfigRepository.GetById(id)

	if err != nil {
		return nil, err
	}

	return tc, nil
}

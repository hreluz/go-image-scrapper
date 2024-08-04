package repotagconfig

import "github.com/hreluz/images-scrapper/pkg/html_processer/tag"

type TagRepository interface {
	Save(t *tag.TagConfig) error
	SaveAll(tagConfigs []*tag.TagConfig) error
	GetAll() ([]*tag.TagConfig, error)
}

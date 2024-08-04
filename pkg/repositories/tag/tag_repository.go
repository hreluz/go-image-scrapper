package repotag

import "github.com/hreluz/images-scrapper/pkg/html_processer/tag"

type TagRepository interface {
	Save(t *tag.Tag) error
	SaveAll(tags []*tag.Tag) error
	GetAll() ([]*tag.Tag, error)
	GetById(id int) (*tag.Tag, error)
}

package repotag

import (
	"github.com/hreluz/images-scrapper/pkg/fileutil"
	"github.com/hreluz/images-scrapper/pkg/html_processer/tag"
)

type TagFileRepository struct {
	FilePath string
}

func (r *TagFileRepository) Save(t *tag.Tag) error {

	err := r.SaveAll([]*tag.Tag{t})

	if err != nil {
		return err
	}

	return nil
}

func (r *TagFileRepository) SaveAll(tags []*tag.Tag) error {

	wrappers := make([]tag.TagWrapper, len(tags))

	for i, t := range tags {
		wrappers[i] = *t.GetWrapper()
	}

	err := fileutil.SaveToFile(r.FilePath, &wrappers)

	if err != nil {
		return err
	}

	return nil
}

func (r *TagFileRepository) GetAll() ([]*tag.Tag, error) {

	var wrappers []tag.TagWrapper

	err := fileutil.LoadFromFile(r.FilePath, &wrappers)

	if err != nil {
		return nil, err
	}

	tags := make([]*tag.Tag, len(wrappers))

	for i, wrapper := range wrappers {
		tags[i] = tag.LoadWrapper(&wrapper)
	}

	return tags, nil
}

func (r *TagFileRepository) GetById(id int) (*tag.Tag, error) {

	tags, err := r.GetAll()

	if err != nil {
		return nil, err
	}

	for _, t := range tags {
		if t.GetID() == id {
			return t, nil
		}
	}

	return nil, nil
}

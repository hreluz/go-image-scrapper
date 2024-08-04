package repotagconfig

import (
	"github.com/hreluz/images-scrapper/pkg/fileutil"
	"github.com/hreluz/images-scrapper/pkg/html_processer/tag"
	repotag "github.com/hreluz/images-scrapper/pkg/repositories/tag"
	servicetag "github.com/hreluz/images-scrapper/pkg/services/tag"
)

type TagConfigFileRepository struct {
	FilePath    string
	TagFilePath string
}

func (r *TagConfigFileRepository) Save(t *tag.TagConfig) error {

	err := r.SaveAll([]*tag.TagConfig{t})

	if err != nil {
		return err
	}

	return nil
}

func (r *TagConfigFileRepository) SaveAll(tagConfigs []*tag.TagConfig) error {

	wrappers := make([]tag.TagConfigWrapper, len(tagConfigs))

	var tags []*tag.Tag

	for i, tc := range tagConfigs {
		wrappers[i] = *tc.GetWrapper()
		tags = append(tags, tc.GetTags()...)
	}

	err := fileutil.SaveToFile(r.FilePath, &wrappers)

	if err != nil {
		return err
	}

	s := servicetag.TagService{TagRepository: &repotag.TagFileRepository{FilePath: r.TagFilePath}}

	err = s.SaveAll(tags)

	if err != nil {
		return err
	}

	return nil
}

func (r *TagConfigFileRepository) GetAll() ([]*tag.TagConfig, error) {

	var wrappers []tag.TagConfigWrapper

	err := fileutil.LoadFromFile(r.FilePath, &wrappers)

	if err != nil {
		return nil, err
	}

	tagConfigs := make([]*tag.TagConfig, len(wrappers))

	// Load Tags
	s := servicetag.TagService{TagRepository: &repotag.TagFileRepository{FilePath: r.TagFilePath}}

	for i, wrapper := range wrappers {
		var tags []*tag.Tag

		for _, tagId := range wrapper.TagIds {
			tagFound, _ := s.GetById(*tagId)
			tags = append(tags, tagFound)
		}

		tagConfigs[i] = tag.LoadTagConfigWrapper(&wrapper, tags)
	}

	return tagConfigs, nil
}

func (r *TagConfigFileRepository) GetById(id int) (*tag.TagConfig, error) {

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

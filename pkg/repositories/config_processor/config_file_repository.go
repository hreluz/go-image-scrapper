package repoconfigprocessor

import (
	"github.com/hreluz/images-scrapper/pkg/fileutil"
	"github.com/hreluz/images-scrapper/pkg/html_processer/tag"
	"github.com/hreluz/images-scrapper/pkg/models/config"
	repotagconfig "github.com/hreluz/images-scrapper/pkg/repositories/tagconfig"
	servicetagconfig "github.com/hreluz/images-scrapper/pkg/services/tagconfig"
)

type ConfigFileRepository struct {
	FilePath          string
	TagFilePath       string
	TagConfigFilePath string
}

func (r *ConfigFileRepository) Save(cp *config.ConfigProcessor) error {

	err := r.SaveAll([]*config.ConfigProcessor{cp})

	if err != nil {
		return err
	}

	return nil
}

func (r *ConfigFileRepository) SaveAll(cps []*config.ConfigProcessor) error {

	repoTC := &repotagconfig.TagConfigFileRepository{FilePath: r.TagConfigFilePath, TagFilePath: r.TagFilePath}

	wrappers := make([]config.ConfigProcessorWrapper, len(cps))

	for i, t := range cps {
		wrappers[i] = *t.GetWrapper()
	}

	err := fileutil.SaveToFile(r.FilePath, &wrappers)

	if err != nil {
		return err
	}

	for _, cp := range cps {
		image := cp.GetConfig(config.TAG_CONFIG_IMAGE)
		desc := cp.GetConfig(config.TAG_CONFIG_DESCRIPTION)
		title := cp.GetConfig(config.TAG_CONFIG_TITLE)

		repoTC.SaveAll([]*tag.TagConfig{image, desc, title})
	}

	return nil
}

func (r *ConfigFileRepository) GetAll() ([]*config.ConfigProcessor, error) {

	var wrappers []config.ConfigProcessorWrapper

	err := fileutil.LoadFromFile(r.FilePath, &wrappers)

	if err != nil {
		return nil, err
	}

	configs := make([]*config.ConfigProcessor, len(wrappers))

	// Load TagConfigs
	repoTC := &repotagconfig.TagConfigFileRepository{FilePath: r.TagConfigFilePath, TagFilePath: r.TagFilePath}
	s := servicetagconfig.TagConfigService{TagConfigRepository: *repoTC}

	for i, wrapper := range wrappers {
		imageConfig, _ := s.GetById(wrapper.ImageConfigID)
		descriptionConfig, _ := s.GetById(wrapper.DescriptionConfigID)
		titleConfig, _ := s.GetById(wrapper.TitleConfigID)

		tagConfigs := []*tag.TagConfig{imageConfig, descriptionConfig, titleConfig}

		configs[i] = config.LoadWrapper(&wrapper, tagConfigs)
	}

	return configs, nil
}

func (r *ConfigFileRepository) GetById(id int) (*config.ConfigProcessor, error) {

	configs, err := r.GetAll()

	if err != nil {
		return nil, err
	}

	for _, c := range configs {
		if c.GetID() == id {
			return c, nil
		}
	}

	return nil, nil
}

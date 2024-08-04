package repoconfigprocessor

import (
	"reflect"
	"testing"

	"github.com/hreluz/images-scrapper/pkg/fileutil"
	"github.com/hreluz/images-scrapper/pkg/html_processer/selector"
	"github.com/hreluz/images-scrapper/pkg/html_processer/tag"
	"github.com/hreluz/images-scrapper/pkg/models/config"
	repotagconfig "github.com/hreluz/images-scrapper/pkg/repositories/tagconfig"
)

func TestSaveOneConfigProcessorOnFile(t *testing.T) {

	repoTC := repotagconfig.TagConfigFileRepository{FilePath: "tagconfig.json", TagFilePath: "tag.json"}

	repo := &ConfigFileRepository{
		FilePath:          "config_processor.json",
		TagConfigFilePath: repoTC.FilePath,
		TagFilePath:       repoTC.TagFilePath,
	}

	defer fileutil.DeleteFile(repo.FilePath)
	defer fileutil.DeleteFile(repo.TagConfigFilePath)
	defer fileutil.DeleteFile(repo.TagFilePath)

	// imageConfig
	tag1 := tag.New(selector.Empty(), tag.P)
	tag2 := tag.New(selector.Empty(), tag.IMG)
	imageConfig := tag.NewConfig(5, []*tag.Tag{tag1, tag2})

	// paginationConfig
	// tag3 := tag.New(selector.Empty(), tag.P)
	// tag4 := tag.New(selector.Empty(), tag.IMG)
	// paginationConfig := pagination.New(tag.NewConfig(5, []*tag.Tag{tag3, tag4}), 4)

	// titleConfig
	tag5 := tag.New(selector.Empty(), tag.P)
	tag6 := tag.New(selector.Empty(), tag.IMG)
	titleConfig := tag.NewConfig(5, []*tag.Tag{tag5, tag6})

	// descriptionConfig
	tag7 := tag.New(selector.Empty(), tag.P)
	tag8 := tag.New(selector.Empty(), tag.IMG)
	descriptionConfig := tag.NewConfig(5, []*tag.Tag{tag7, tag8})

	// Initialize the image processor and downloader
	cp := config.NewProcessor("test-config", imageConfig, nil, titleConfig, descriptionConfig)

	repo.Save(cp)
	cpLoaded, _ := repo.GetAll()

	isPaginationEqual := !reflect.DeepEqual(cp.GetPagination(), cpLoaded[0].GetPagination())
	isImageEqual := !reflect.DeepEqual(cp.GetConfig(config.TAG_CONFIG_IMAGE), cpLoaded[0].GetConfig(config.TAG_CONFIG_IMAGE))
	isDescEqual := !reflect.DeepEqual(cp.GetConfig(config.TAG_CONFIG_DESCRIPTION), cpLoaded[0].GetConfig(config.TAG_CONFIG_DESCRIPTION))
	isTitleEqual := !reflect.DeepEqual(cp.GetConfig(config.TAG_CONFIG_TITLE), cpLoaded[0].GetConfig(config.TAG_CONFIG_TITLE))
	isCPAttrEqual := cp.GetID() == cpLoaded[0].GetID() && cp.GetName() == cpLoaded[0].GetName()

	if isPaginationEqual && isImageEqual && isDescEqual && isTitleEqual && isCPAttrEqual {
		t.Fatalf("The loaded config processor does not match with what was expected")
	}
}

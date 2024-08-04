package repotagconfig

import (
	"reflect"
	"testing"

	"github.com/hreluz/images-scrapper/pkg/fileutil"
	"github.com/hreluz/images-scrapper/pkg/html_processer/selector"
	"github.com/hreluz/images-scrapper/pkg/html_processer/tag"
)

func TestSaveOneTagConfigOnFile(t *testing.T) {

	repo := TagConfigFileRepository{FilePath: "tagconfig.json", TagFilePath: "tag.json"}

	defer fileutil.DeleteFile(repo.FilePath)
	defer fileutil.DeleteFile(repo.TagFilePath)

	tag1 := tag.New(selector.New(selector.ID, "first-container"), tag.DIV)
	tag2 := tag.New(selector.New(selector.CLASS, "second-container"), tag.DIV)
	tag3 := tag.New(selector.Empty(), tag.A)

	tagConfig := tag.NewConfig(3, []*tag.Tag{tag1, tag2, tag3})

	repo.Save(tagConfig)

	tagConfigLoaded, _ := repo.GetAll()

	if !reflect.DeepEqual(tagConfig, tagConfigLoaded[0]) {
		t.Fatalf("The loaded tag config does not match with what was expected")
	}
}

func TestSaveManyTagConfigsOnFile(t *testing.T) {

	repo := TagConfigFileRepository{FilePath: "tagconfig.json", TagFilePath: "tag.json"}
	defer fileutil.DeleteFile(repo.FilePath)
	defer fileutil.DeleteFile(repo.TagFilePath)

	tag1 := tag.New(selector.New(selector.ID, "first-container"), tag.DIV)
	tag2 := tag.New(selector.New(selector.CLASS, "second-container"), tag.DIV)

	tag3 := tag.New(selector.Empty(), tag.A)
	tag4 := tag.New(selector.New(selector.CLASS, "third-container"), tag.NAV)
	tag5 := tag.New(selector.New(selector.ID, "fourth-container"), tag.IMG)
	tag6 := tag.New(selector.Empty(), tag.A)

	tagConfig1 := tag.NewConfig(2, []*tag.Tag{tag1, tag2})

	tagConfig2 := tag.NewConfig(4, []*tag.Tag{tag3, tag4, tag5, tag6})

	tagConfigs := []*tag.TagConfig{tagConfig1, tagConfig2}

	repo.SaveAll(tagConfigs)

	tagsLoaded, _ := repo.GetAll()

	if !reflect.DeepEqual(tagConfigs, tagsLoaded) {
		t.Fatalf("The loaded tag configs do not match with the tags that was expected")
	}
}

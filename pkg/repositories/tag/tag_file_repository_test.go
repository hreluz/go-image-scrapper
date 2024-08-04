package repotag

import (
	"reflect"
	"testing"

	"github.com/hreluz/images-scrapper/pkg/fileutil"
	"github.com/hreluz/images-scrapper/pkg/html_processer/selector"
	"github.com/hreluz/images-scrapper/pkg/html_processer/tag"
)

func TestSaveOneTagOnFile(t *testing.T) {

	repo := TagFileRepository{FilePath: "tag.json"}

	defer fileutil.DeleteFile(repo.FilePath)

	tagSaved := tag.New(selector.New(selector.CLASS, "some-name"), tag.P)

	repo.Save(tagSaved)

	tagLoaded, _ := repo.GetAll()

	if !reflect.DeepEqual(tagSaved, tagLoaded[0]) {
		t.Fatalf("The loaded tag does not match with what was expected")
	}
}

func TestSaveManyTagsOnFile(t *testing.T) {

	repo := TagFileRepository{FilePath: "tags-many.json"}

	defer fileutil.DeleteFile(repo.FilePath)

	tagSaved1 := tag.New(selector.New(selector.CLASS, "some-name"), tag.P)

	tagSaved2 := tag.New(selector.New(selector.ID, "some-name-2"), tag.ARTICLE)

	tags := []*tag.Tag{tagSaved1, tagSaved2}

	repo.SaveAll(tags)

	tagsLoaded, _ := repo.GetAll()

	if !reflect.DeepEqual(tags, tagsLoaded) {
		t.Fatalf("The loaded tags does not match with the tags that was expected")
	}
}

func TestGetByIdOnFile(t *testing.T) {

	repo := TagFileRepository{FilePath: "tags-many.json"}

	defer fileutil.DeleteFile(repo.FilePath)

	tagSaved1 := tag.New(selector.New(selector.CLASS, "some-name"), tag.P)

	tagSaved2 := tag.New(selector.New(selector.ID, "some-name-2"), tag.ARTICLE)

	tags := []*tag.Tag{tagSaved1, tagSaved2}

	repo.SaveAll(tags)

	gotTag, _ := repo.GetById(tagSaved2.GetID())

	if !reflect.DeepEqual(tagSaved2, gotTag) {
		t.Fatalf("It was not found the tag in tags file, it was expected %v, it got %v", tagSaved2, gotTag)
	}
}

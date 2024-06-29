package tag

import (
	selector "github.com/hreluz/images-scrapper/pkg/html_processer/selector"
)

type TagName string
type TagNames []TagName

const (
	DIV     TagName = "div"
	ARTICLE TagName = "article"
	SPAN    TagName = "span"
)

var TAGS_OPTIONS = TagNames{DIV, ARTICLE, SPAN}

type Tag struct {
	selector *selector.Selector
	name     TagName
}

// New returns a new Tag
func New(selector *selector.Selector, name TagName) *Tag {
	return &Tag{
		selector,
		name,
	}
}

func (t *Tag) GetSelector() *selector.Selector {
	return t.selector
}

func (t *Tag) GetName() TagName {
	return t.name
}

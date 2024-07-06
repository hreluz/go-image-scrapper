package tag

import (
	"fmt"

	"github.com/anaskhan96/soup"
	selector "github.com/hreluz/images-scrapper/pkg/html_processer/selector"
)

type TagName string
type TagNames []TagName

const (
	DIV     TagName = "div"
	ARTICLE TagName = "article"
	SPAN    TagName = "span"
	A       TagName = "a"
	P       TagName = "p"
	IMG     TagName = "img"
	NAV     TagName = "nav"
	H1      TagName = "h1"
)

var TAGS_OPTIONS = TagNames{DIV, ARTICLE, SPAN, P, A, IMG, NAV, H1}

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

func (t *Tag) GetContentByTag(doc soup.Root) (soup.Root, error) {

	tagName := string(t.GetName())

	if t.GetSelector().GetType() == selector.NONE {
		selectorFound := doc.Find(tagName)

		if selectorFound.Error != nil {
			return soup.Root{}, fmt.Errorf("tag not found, error: %v", selectorFound.Error)
		}

		return selectorFound, nil
	}

	selectorName := string(t.GetSelector().GetName())
	selectorType := string(t.GetSelector().GetType())
	selectorFound := doc.Find(tagName, selectorType, selectorName)

	if selectorFound.Error != nil {
		return soup.Root{}, fmt.Errorf("%s provided (%s) does not exist, error: %v", selectorType, selectorName, selectorFound.Error)
	}

	return selectorFound, nil
}

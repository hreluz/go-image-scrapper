package tag

import (
	"fmt"

	"github.com/anaskhan96/soup"
	"github.com/hreluz/images-scrapper/pkg/html_processer/selector"
)

type TagConfig struct {
	levels int
	tags   []Tag
}

func (tc *TagConfig) GetLevels() int {
	return tc.levels
}

func (tc *TagConfig) GetTags() []Tag {
	return tc.tags
}

func (tc *TagConfig) AddTag(tag *Tag) {
	tc.tags = append(tc.tags, *tag)
}

func NewConfig(levels int, tags []Tag) *TagConfig {
	return &TagConfig{
		levels,
		tags,
	}
}

func (tc *TagConfig) GetLastTagContainer(container soup.Root) (soup.Root, error) {
	var err error

	for _, tag := range tc.GetTags() {
		container, err = tag.GetContentByTag(container)

		if err != nil {
			if tag.GetSelector().GetType() == selector.NONE {
				return soup.Root{},
					fmt.Errorf("there was an error trying to find the tag %v", tag.GetName())
			}

			return soup.Root{},
				fmt.Errorf("there was an error trying to find the selector %v with the name %v",
					tag.GetSelector().GetType(),
					tag.GetSelector().GetName(),
				)
		}
	}

	return container, nil
}

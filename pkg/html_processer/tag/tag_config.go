package tag

import (
	"fmt"
	"log"

	"github.com/anaskhan96/soup"
	"github.com/hreluz/images-scrapper/pkg/helpers"
	"github.com/hreluz/images-scrapper/pkg/html_processer/selector"
)

type TagConfig struct {
	id     int
	levels int
	tags   []*Tag
}

type TagConfigWrapper struct {
	ID     int    `json:"id"`
	Levels int    `json:"levels"`
	TagIds []*int `json:"tagIds"`
}

func (tc *TagConfig) GetWrapper() *TagConfigWrapper {
	var tagIds []*int

	for i := 0; i < len(tc.tags); i++ {
		tagIds = append(tagIds, &tc.tags[i].id)
	}

	return &TagConfigWrapper{
		tc.id,
		tc.levels,
		tagIds,
	}
}

func LoadTagConfigWrapper(tw *TagConfigWrapper, tags []*Tag) *TagConfig {
	return &TagConfig{
		tw.ID,
		tw.Levels,
		tags,
	}
}

func (tc *TagConfig) GetID() int {
	return tc.id
}

func (tc *TagConfig) GetLevels() int {
	return tc.levels
}

func (tc *TagConfig) GetTags() []*Tag {
	return tc.tags
}

func (tc *TagConfig) AddTag(tag *Tag) {
	tc.tags = append(tc.tags, tag)
}

func NewConfig(levels int, tags []*Tag) *TagConfig {
	return &TagConfig{
		id:     helpers.GetRandomNumber(),
		levels: levels,
		tags:   tags,
	}
}

func (tc *TagConfig) GetLastTagContainer(html soup.Root) (soup.Root, error) {
	var err error

	for _, tag := range tc.GetTags() {
		html, err = tag.GetContentByTag(html)

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

	return html, nil
}

func (tc *TagConfig) ProcessText(html soup.Root) string {
	text, err := tc.GetLastTagContainer(html)

	if err != nil {
		log.Fatalf("Error trying to get text tag for processing text, error: %v", err)
	}

	return text.Text()
}

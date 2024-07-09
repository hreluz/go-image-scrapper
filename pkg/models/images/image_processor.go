package images

import (
	"log"

	"github.com/anaskhan96/soup"
	"github.com/hreluz/images-scrapper/pkg/html_processer/pagination"
	"github.com/hreluz/images-scrapper/pkg/html_processer/tag"
)

type tagConfigType string

const (
	TAG_CONFIG_IMAGE       tagConfigType = "image"
	TAG_CONFIG_TITLE       tagConfigType = "title"
	TAG_CONFIG_DESCRIPTION tagConfigType = "description"
	TAG_CONFIG_PAGINATION  tagConfigType = "pagination"
)

type ImageProcessor struct {
	ic *tag.TagConfig
	pc *pagination.Pagination
	tc *tag.TagConfig
	dc *tag.TagConfig
}

func NewProcessor(ic *tag.TagConfig, pc *pagination.Pagination, tc *tag.TagConfig, dc *tag.TagConfig) *ImageProcessor {
	return &ImageProcessor{
		ic: ic,
		pc: pc,
		tc: tc,
		dc: dc,
	}
}

func (ip *ImageProcessor) GetConfig(c tagConfigType) *tag.TagConfig {
	switch c {
	case TAG_CONFIG_IMAGE:
		return ip.ic
	case TAG_CONFIG_TITLE:
		return ip.tc
	case TAG_CONFIG_DESCRIPTION:
		return ip.dc
	default:
		log.Fatalf("Error, this type of config type is not in the list %s", c)
	}

	return nil
}

func (ip *ImageProcessor) GetPagination() *pagination.Pagination {
	return ip.pc
}

func ProcessText(t *tag.TagConfig, html soup.Root) string {
	text, err := t.GetLastTagContainer(html)

	if err != nil {
		log.Fatalf("Error trying to get text tag for processing text, error: %v", err)
	}

	return text.Text()
}

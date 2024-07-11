package config

import (
	"log"
	"time"

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

type ConfigProcessor struct {
	id int64
	ic *tag.TagConfig
	pc *pagination.Pagination
	tc *tag.TagConfig
	dc *tag.TagConfig
}

func NewProcessor(ic *tag.TagConfig, pc *pagination.Pagination, tc *tag.TagConfig, dc *tag.TagConfig) *ConfigProcessor {
	return &ConfigProcessor{
		id: time.Now().Unix(),
		ic: ic,
		pc: pc,
		tc: tc,
		dc: dc,
	}
}

func (ip *ConfigProcessor) GetConfig(c tagConfigType) *tag.TagConfig {
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

func (cp *ConfigProcessor) GetPagination() *pagination.Pagination {
	return cp.pc
}

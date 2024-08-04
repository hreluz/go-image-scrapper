package config

import (
	"log"

	"github.com/hreluz/images-scrapper/pkg/helpers"
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
	id   int
	name string
	ic   *tag.TagConfig
	pc   *pagination.Pagination
	tc   *tag.TagConfig
	dc   *tag.TagConfig
}

type ConfigProcessorWrapper struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	ImageConfigID       int    `json:"image_config_id"`
	PaginationID        int    `json:"pagination_id"`
	TitleConfigID       int    `json:"title_config_id"`
	DescriptionConfigID int    `json:"description_config_id"`
}

func (cp *ConfigProcessor) GetWrapper() *ConfigProcessorWrapper {
	return &ConfigProcessorWrapper{
		ID:                  cp.id,
		Name:                cp.name,
		ImageConfigID:       cp.GetConfig(TAG_CONFIG_IMAGE).GetID(),
		PaginationID:        cp.pc.GetID(),
		TitleConfigID:       cp.GetConfig(TAG_CONFIG_TITLE).GetID(),
		DescriptionConfigID: cp.GetConfig(TAG_CONFIG_DESCRIPTION).GetID(),
	}
}

func (cp *ConfigProcessor) GetID() int {
	return cp.id
}

func (cp *ConfigProcessor) GetName() string {
	return cp.name
}

func LoadWrapper(cp *ConfigProcessorWrapper, tagConfigs []*tag.TagConfig) *ConfigProcessor {
	new := NewProcessor(cp.Name, tagConfigs[0], nil, tagConfigs[1], tagConfigs[2])
	new.id = cp.ID
	return new
}

func NewProcessor(name string, ic *tag.TagConfig, pc *pagination.Pagination, tc *tag.TagConfig, dc *tag.TagConfig) *ConfigProcessor {
	return &ConfigProcessor{
		id:   helpers.GetRandomNumber(),
		name: name,
		ic:   ic,
		pc:   pc,
		tc:   tc,
		dc:   dc,
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

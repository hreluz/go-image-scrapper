package images

import "github.com/hreluz/images-scrapper/pkg/html_processer/tag"

type ImageConfig struct {
	tc *tag.TagConfig
}

func NewImageConfig(tc *tag.TagConfig) *ImageConfig {
	return &ImageConfig{tc}
}

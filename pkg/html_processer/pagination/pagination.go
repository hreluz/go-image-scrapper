package pagination

import (
	tag "github.com/hreluz/images-scrapper/pkg/html_processer/tag"
)

type Pagination struct {
	tag    *tag.Tag
	number int
}

// New returns a new Selector
func New(tag *tag.Tag, number int) *Pagination {
	return &Pagination{
		tag,
		number,
	}
}

func (p *Pagination) GetNumber() int {
	return p.number
}

func (p *Pagination) GetTag() *tag.Tag {
	return p.tag
}

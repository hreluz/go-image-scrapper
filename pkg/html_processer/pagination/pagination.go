package pagination

import (
	"github.com/anaskhan96/soup"
	"github.com/hreluz/images-scrapper/pkg/helpers"
	tag "github.com/hreluz/images-scrapper/pkg/html_processer/tag"
)

type Pagination struct {
	id     int
	tc     *tag.TagConfig
	number int
}

type PaginationWrapper struct {
	ID                 int                   `json:"id"`
	PaginationConfigID *tag.TagConfigWrapper `json:"pagination_config_id"`
	Number             int                   `json:"number"`
}

func New(tc *tag.TagConfig, number int) *Pagination {
	return &Pagination{
		id:     helpers.GetRandomNumber(),
		tc:     tc,
		number: number,
	}
}

func (p *Pagination) GetID() int {

	if p == nil {
		return 0
	}

	return p.id
}

func (p *Pagination) GetNumber() int {
	return p.number
}

func (p *Pagination) GetTagConfig() *tag.TagConfig {
	return p.tc
}

func (p *Pagination) GetPaginationNextLink(container soup.Root) (string, error) {
	a, err := p.tc.GetLastTagContainer(container)

	if err != nil {
		return "", err
	}

	return a.Attrs()["href"], nil
}

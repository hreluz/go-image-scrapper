package pagination

import (
	"github.com/anaskhan96/soup"
	tag "github.com/hreluz/images-scrapper/pkg/html_processer/tag"
)

type Pagination struct {
	tc     *tag.TagConfig
	number int
}

// New returns a new Selector
func New(tc *tag.TagConfig, number int) *Pagination {
	return &Pagination{
		tc,
		number,
	}
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

package interaction

import (
	"fmt"

	"github.com/hreluz/images-scrapper/pkg/html_processer/selector"
	"github.com/hreluz/images-scrapper/pkg/html_processer/tag"
)

func ShowTagOptions(s string) {
	fmt.Printf("Choose a tag %s:\n", s)
	fmt.Println("------------------------------------")

	for i, option := range tag.TAGS_OPTIONS {
		fmt.Printf("(%d) %s\n", (i + 1), option)
	}
}

func ShowSelectorOptions(s string) {
	fmt.Printf("Choose a selector %s:\n", s)
	fmt.Println("------------------------------------")

	for i, option := range selector.SELECTOR_TYPE_OPTIONS {
		fmt.Printf("(%d) %s\n", (i + 1), option)
	}
}

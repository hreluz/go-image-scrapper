package interaction

import (
	"fmt"

	htmlprocesser "github.com/hreluz/images-scrapper/pkg/html_processer"
)

func ShowTagOptions(s string) {
	fmt.Printf("Choose a tag %s:\n", s)
	fmt.Println("------------------------------------")

	for i, option := range htmlprocesser.TAGS_OPTIONS {
		fmt.Printf("(%d) %s\n", (i + 1), option)
	}
}

func ShowSelectorOptions(s string) {
	fmt.Printf("Choose a selector %s:\n", s)
	fmt.Println("------------------------------------")

	for i, option := range htmlprocesser.SELECTOR_TYPE_OPTIONS {
		fmt.Printf("(%d) %s\n", (i + 1), option)
	}
}

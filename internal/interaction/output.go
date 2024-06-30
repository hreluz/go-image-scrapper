package interaction

import (
	"fmt"
	"reflect"

	"github.com/hreluz/images-scrapper/pkg/html_processer/selector"
	"github.com/hreluz/images-scrapper/pkg/html_processer/tag"
)

func ShowTagOptions(s string) {
	showOptions(tag.TAGS_OPTIONS, fmt.Sprintf("Choose a tag %s", s))
}

func ShowSelectorOptions(s string) {
	showOptions(selector.SELECTOR_TYPE_OPTIONS, fmt.Sprintf("Choose a selector %s", s))
}

func convertToStringSlice(v interface{}) []string {
	var options []string
	rv := reflect.ValueOf(v)
	for i := 0; i < rv.Len(); i++ {
		options = append(options, rv.Index(i).String())
	}
	return options
}

func showOptions(v interface{}, message string) {
	fmt.Printf("%v:\n ", message)
	fmt.Println("------------------------------------")

	var options []string
	switch v := v.(type) {
	case tag.TagNames:
		options = convertToStringSlice(v)
	case selector.SelectorTypes:
		options = convertToStringSlice(v)
	case []string:
		options = v
	default:
		fmt.Println("Unsupported type")
		return
	}

	for i, option := range options {
		fmt.Printf("(%d) %s\n", (i + 1), option)
	}
}

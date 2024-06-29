package htmlprocesser

import (
	"fmt"

	"github.com/anaskhan96/soup"
)

type TagName string
type TagNames []TagName

const (
	DIV     TagName = "div"
	ARTICLE TagName = "article"
	SPAN    TagName = "span"
)

var TAGS_OPTIONS = TagNames{DIV, ARTICLE, SPAN}

type SelectorName string
type SelectorType string
type SelectorTypes []SelectorType

const (
	ID    SelectorType = "id"
	CLASS SelectorType = "class"
)

var SELECTOR_TYPE_OPTIONS = SelectorTypes{ID, CLASS}

type ImageUrls []string

type Processer struct {
	url       string
	imageUrls ImageUrls
}

type Selector struct {
	Type SelectorType
	Name SelectorName
}

type Tag struct {
	Selector Selector
	Name     TagName
}

type Pagination struct {
	Tag    Tag
	Number int
}

func GetHTMLParsed(url string) (soup.Root, error) {
	resp, err := soup.Get(url)

	if err != nil {
		return soup.Root{}, fmt.Errorf("there was an error resolving the url %v", err)
	}

	doc := soup.HTMLParse(resp)
	return doc, nil
}

func GetDivBySelector(selector string, selectorName string, doc soup.Root) (soup.Root, error) {
	div_class := doc.Find("div", selector, selectorName)

	if div_class.Error != nil {
		return soup.Root{}, fmt.Errorf("%s provided (%s) does not exist, error: %v", selector, selectorName, div_class.Error)
	}

	return div_class, nil
}

func GetImageLinksFrom(container soup.Root) ImageUrls {
	var image_links ImageUrls

	images_tags := container.FindAll("img")

	for _, image_tag := range images_tags {
		image_url := image_tag.Attrs()["src"]
		image_links = append(image_links, image_url)
		fmt.Println("Image link added :", image_url)
	}

	return image_links
}

func GetPaginationNextLink(container soup.Root, className string) (string, error) {
	span := container.Find("span", "class", className)

	if span.Error != nil {
		return "", fmt.Errorf("span with class %s was not found in html", className)
	}

	a := span.Find("a")

	if a.Error != nil {
		return "", fmt.Errorf("element a was not found in span")
	}

	return a.Attrs()["href"], nil
}

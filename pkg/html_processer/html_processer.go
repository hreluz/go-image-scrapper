package htmlprocesser

import (
	"fmt"

	"github.com/anaskhan96/soup"
	"github.com/hreluz/images-scrapper/pkg/html_processer/pagination"
	"github.com/hreluz/images-scrapper/pkg/html_processer/tag"
)

type ImageUrls []string

type Processer struct {
	url       string
	imageUrls ImageUrls
}

func GetHTMLParsed(url string) (soup.Root, error) {
	resp, err := soup.Get(url)

	if err != nil {
		return soup.Root{}, fmt.Errorf("there was an error resolving the url %v", err)
	}

	doc := soup.HTMLParse(resp)
	return doc, nil
}

func GetBySelector(t *tag.Tag, doc soup.Root) (soup.Root, error) {
	selectorName := string(t.GetSelector().GetName())
	selectorType := string(t.GetSelector().GetType())
	tagName := string(t.GetName())

	selectorFound := doc.Find(tagName, selectorType, selectorName)

	if selectorFound.Error != nil {
		return soup.Root{}, fmt.Errorf("%s provided (%s) does not exist, error: %v", selectorType, selectorName, selectorFound.Error)
	}

	return selectorFound, nil
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

func GetPaginationNextLink(container soup.Root, p *pagination.Pagination) (string, error) {

	selectorName := string(p.GetTag().GetSelector().GetName())
	selectorType := string(p.GetTag().GetSelector().GetType())
	tagName := string(p.GetTag().GetName())

	span := container.Find(tagName, selectorType, selectorName)

	if span.Error != nil {
		return "", fmt.Errorf("span with %s %s was not found in html", selectorType, selectorName)
	}

	a := span.Find("a")

	if a.Error != nil {
		return "", fmt.Errorf("element a was not found in %s", selectorName)
	}

	return a.Attrs()["href"], nil
}

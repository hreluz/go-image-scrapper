package htmlprocesser

import (
	"fmt"
	"log"

	"github.com/anaskhan96/soup"
)

type ImageUrls []string

func GetHTMLParsed(url string) (soup.Root, error) {
	resp, err := soup.Get(url)

	if err != nil {
		return soup.Root{}, fmt.Errorf("there was an error resolving the url %v", err)
	}

	doc := soup.HTMLParse(resp)
	return doc, nil
}

func GetDivByClass(class_name string, doc soup.Root) (soup.Root, error) {
	div_class := doc.Find("div", "class", class_name)

	if div_class.Error != nil {
		return soup.Root{}, fmt.Errorf("class provided (%s) does not exist, error: %v", class_name, div_class.Error)
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

func GetPaginationNextLink(container soup.Root, class_name string) string {
	return container.Find("span", "class", class_name).Find("a").Attrs()["href"]
}

func ExecByClass(url, class_name string) ImageUrls {
	html_content, err := GetHTMLParsed(url)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	div_class, err := GetDivByClass(class_name, html_content)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return GetImageLinksFrom(div_class)
}

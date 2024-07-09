package images

import (
	"fmt"

	"github.com/anaskhan96/soup"
)

// htm,  soup.Root

type Image struct {
	html        string
	url         string
	webUrl      string
	nextUrl     string
	title       string
	description string
}

type ImageUrls []string
type Images []Image

func New(html, url, webUrl, nextUrl, title, description string) *Image {
	return &Image{
		html,
		url,
		webUrl,
		nextUrl,
		title,
		description,
	}
}

func (i *Image) GetHtmlSoup() soup.Root {
	return soup.HTMLParse(i.html)
}

func (i *Image) GetUrl() string {
	return i.url
}

func (i *Image) GetNextUrl() string {
	return i.nextUrl
}

func (i *Image) GetWebUrl() string {
	return i.webUrl
}

func (i *Image) String() string {
	multiLine := "------------------------------------\n"
	webUrl := fmt.Sprintf("Web Url: %s\n", i.webUrl)
	imageUrl := fmt.Sprintf("Image Url: %s\n", i.url)
	title := fmt.Sprintf("Title: %s\n", i.title)
	description := fmt.Sprintf("Description: %s\n", i.description)

	return fmt.Sprintf("\n%s%s%s%s%s%s", multiLine, webUrl, imageUrl, title, description, multiLine)
}

package image

import (
	"fmt"
	"log"

	"github.com/anaskhan96/soup"
	"github.com/hreluz/images-scrapper/pkg/html_processer/loader"
	"github.com/hreluz/images-scrapper/pkg/html_processer/pagination"
	"github.com/hreluz/images-scrapper/pkg/html_processer/tag"
)

type ImageUrls []string
type Images []Image

type ImageProcessor struct {
	ic *tag.TagConfig
	pc *pagination.Pagination
}

type Image struct {
	HTMLParsed soup.Root
	url        string
	webUrl     string
	processed  bool
	nextUrl    string
}

func NewProcessor(ic *tag.TagConfig, pc *pagination.Pagination) *ImageProcessor {
	return &ImageProcessor{
		ic: ic,
		pc: pc,
	}
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

func Process(ip *ImageProcessor, webUrl string) *Image {

	var i Image

	i.webUrl = webUrl

	htmlParsed, err := loader.GetHTMLParsed(webUrl)

	if err != nil {
		log.Printf("Error parsing HTML: %s", err)
	}

	imageTag, err := ip.ic.GetLastTagContainer(htmlParsed)

	if err != nil {
		log.Fatalf("Error trying to get image tag on fill image func, error: %v", err)
	}

	i.url = imageTag.Attrs()["src"]

	fmt.Println("Image link added :", i.url)

	if ip.pc != nil && ip.pc.GetNumber() > 1 {
		i.nextUrl, err = ip.pc.GetPaginationNextLink(htmlParsed)

		fmt.Printf("Pagination link added %s", i.nextUrl)

		if err != nil {
			log.Fatalf("Error trying to get the pagination for the next link, error: %v", err)
		}
	}

	return &i
}

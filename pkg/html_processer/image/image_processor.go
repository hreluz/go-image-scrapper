package image

import (
	"fmt"
	"log"

	"github.com/anaskhan96/soup"
	"github.com/hreluz/images-scrapper/pkg/html_processer/loader"
	"github.com/hreluz/images-scrapper/pkg/html_processer/pagination"
	"github.com/hreluz/images-scrapper/pkg/html_processer/tag"
)

type ImageProcessor struct {
	ic *tag.TagConfig
	pc *pagination.Pagination
	tc *tag.TagConfig
	dc *tag.TagConfig
}

func NewProcessor(ic *tag.TagConfig, pc *pagination.Pagination, tc *tag.TagConfig, dc *tag.TagConfig) *ImageProcessor {
	return &ImageProcessor{
		ic: ic,
		pc: pc,
		tc: tc,
		dc: dc,
	}
}

func processText(t *tag.TagConfig, html soup.Root) string {
	text, err := t.GetLastTagContainer(html)

	if err != nil {
		log.Fatalf("Error trying to get text tag, error: %v", err)
	}
	return text.Text()
}

func NewImage(ip *ImageProcessor, webUrl string) *Image {

	if ip == nil {
		log.Fatalf("Image processor cannot be nil")
	}

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

	if ip.tc != nil {
		i.title = processText(ip.tc, htmlParsed)
	}

	if ip.dc != nil {
		i.description = processText(ip.dc, htmlParsed)
	}

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

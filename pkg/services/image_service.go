package services

import (
	"fmt"
	"log"

	"github.com/hreluz/images-scrapper/pkg/html_processer/loader"
	"github.com/hreluz/images-scrapper/pkg/models/images"
)

type ImageService struct {
	ip     *images.ImageProcessor
	webUrl string
}

func NewImageService(ip *images.ImageProcessor) func(s string) *ImageService {
	return func(webUrl string) *ImageService {
		return &ImageService{ip, webUrl}
	}
}

func ProcessImage(is *ImageService) *images.Image {

	if is.ip == nil {
		log.Fatalf("Image processor cannot be nil")
	}

	htmlParsed, err := loader.GetHTMLParsed(is.webUrl)

	if err != nil {
		log.Printf("Error parsing HTML: %s", err)
	}

	imageTag, err := is.ip.GetConfig(images.TAG_CONFIG_IMAGE).GetLastTagContainer(htmlParsed)

	if err != nil {
		log.Fatalf("Error trying to get image tag on fill image func, error: %v", err)
	}

	imageUrl := imageTag.Attrs()["src"]

	title := ""
	description := ""

	if is.ip.GetConfig(images.TAG_CONFIG_TITLE) != nil {
		title = images.ProcessText(is.ip.GetConfig(images.TAG_CONFIG_TITLE), htmlParsed)
	}

	if is.ip.GetConfig(images.TAG_CONFIG_DESCRIPTION) != nil {
		description = images.ProcessText(is.ip.GetConfig(images.TAG_CONFIG_DESCRIPTION), htmlParsed)
	}

	fmt.Println("Image link added :", imageUrl)

	nextUrl := ""

	pag := is.ip.GetPagination()

	if pag != nil && pag.GetNumber() > 1 {
		nextUrl, err = pag.GetPaginationNextLink(htmlParsed)

		fmt.Printf("Pagination link added %s", nextUrl)

		if err != nil {
			log.Fatalf("Error trying to get the pagination for the next link, error: %v", err)
		}
	}

	return images.New(
		htmlParsed.HTML(),
		imageUrl,
		is.webUrl,
		nextUrl,
		title,
		description,
	)
}

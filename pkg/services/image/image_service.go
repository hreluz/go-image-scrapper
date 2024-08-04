package serviceimage

import (
	"fmt"
	"log"

	"github.com/hreluz/images-scrapper/pkg/html_processer/loader"
	"github.com/hreluz/images-scrapper/pkg/models/config"
	"github.com/hreluz/images-scrapper/pkg/models/images"
)

type ImageService struct {
	ip     *config.ConfigProcessor
	webUrl string
}

func NewImageService(cp *config.ConfigProcessor) func(s string) *ImageService {
	return func(webUrl string) *ImageService {
		return &ImageService{cp, webUrl}
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

	imageTag, err := is.ip.GetConfig(config.TAG_CONFIG_IMAGE).GetLastTagContainer(htmlParsed)

	if err != nil {
		log.Fatalf("Error trying to get image tag on fill image func, error: %v", err)
	}

	imageUrl := imageTag.Attrs()["src"]

	title := ""
	description := ""

	if is.ip.GetConfig(config.TAG_CONFIG_TITLE) != nil {
		tagConfigTitle := is.ip.GetConfig(config.TAG_CONFIG_TITLE)
		title = tagConfigTitle.ProcessText(htmlParsed)
	}

	if is.ip.GetConfig(config.TAG_CONFIG_DESCRIPTION) != nil {
		tagConfigDescription := is.ip.GetConfig(config.TAG_CONFIG_DESCRIPTION)
		description = tagConfigDescription.ProcessText(htmlParsed)
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

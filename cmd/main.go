package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hreluz/images-scrapper/internal/interaction"
	htmlprocesser "github.com/hreluz/images-scrapper/pkg/html_processer"
	pagination "github.com/hreluz/images-scrapper/pkg/html_processer/pagination"
	selector "github.com/hreluz/images-scrapper/pkg/html_processer/selector"
	tag "github.com/hreluz/images-scrapper/pkg/html_processer/tag"
	imagedownloader "github.com/hreluz/images-scrapper/pkg/image_downloader"
)

func getPagination() *pagination.Pagination {
	resp := interaction.GetUserInputWithErrorHandling("Does this URL have pagination (Y/N)?")

	if resp == "Y" {
		v := interaction.GetUserInputWithErrorHandling("How many URLs would you like to check?")

		number, err := strconv.Atoi(v)

		if err != nil {
			log.Fatalf("Invalid number for pagination, error: %s", err)
		}

		t := *getTagData(
			"where to extract the pagination (it should be the parent of pagination)",
			"Insert selector name where to get the pagination link",
		)

		return pagination.New(
			&t,
			number,
		)
	}

	return nil
}

func processImages(url string, t *tag.Tag, p *pagination.Pagination) htmlprocesser.ImageUrls {
	var imageUrls htmlprocesser.ImageUrls

	for i := 0; i < p.GetNumber(); i++ {

		htmlParsed, err := htmlprocesser.GetHTMLParsed(url)

		if err != nil {
			log.Printf("Error parsing HTML: %s", err)
			continue
		}

		divClass, err := htmlprocesser.GetBySelector(t, htmlParsed)

		if err != nil {
			log.Printf("Error parsing div: %s", err)
			continue
		}

		imageUrls = append(imageUrls, htmlprocesser.GetImageLinksFrom(divClass)...)

		if p.GetNumber() > 1 {
			url, err = htmlprocesser.GetPaginationNextLink(htmlParsed, p)
			if err != nil {
				log.Fatalf("there was en error when trying to get the next link in the pagination, error was %s", err.Error())
			}
		}
	}

	return imageUrls
}

func getTagData(tagOptionText string, selectorTypeText string) *tag.Tag {

	interaction.ShowTagOptions(tagOptionText)

	tagName := interaction.GetTagChoice()

	interaction.ShowSelectorOptions(fmt.Sprintf("for the %s", tagName))

	return tag.New(
		selector.New(
			interaction.GetSelectorTypeChoice(),
			interaction.GetUserInputWithErrorHandling(selectorTypeText),
		),
		tagName,
	)
}

func main() {

	imgChannel := make(chan bool)

	imagesToProcess := 0

	tagToExtractImage := getTagData(
		"where to extract the img (it should be the parent of <img>)",
		"Insert selector name where to pull all the images",
	)

	url := interaction.GetUserInputWithErrorHandling("Insert URL")

	pagination := getPagination()

	imageUrls := processImages(url, tagToExtractImage, pagination)

	id := &imagedownloader.ImageDownloader{
		Download_folder_path: "../downloaded_images",
		Img_channel:          imgChannel,
		Prefix_image:         "image_",
	}

	for _, image_url := range imageUrls {
		imagesToProcess++
		go imagedownloader.Download(id, image_url)
	}

	for i := 0; i < imagesToProcess; i++ {
		<-imgChannel
	}
}

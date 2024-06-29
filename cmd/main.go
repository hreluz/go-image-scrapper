package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hreluz/images-scrapper/internal/interaction"
	htmlprocesser "github.com/hreluz/images-scrapper/pkg/html_processer"
	imagedownloader "github.com/hreluz/images-scrapper/pkg/image_downloader"
)

func getPagination() *htmlprocesser.Pagination {
	pagination := &htmlprocesser.Pagination{
		Tag: htmlprocesser.Tag{
			Selector: htmlprocesser.Selector{
				Type: "",
				Name: "",
			},
			Name: "",
		},
		Number: 1,
	}

	resp := interaction.GetUserInputWithErrorHandling("Does this URL have pagination (Y/N)?")

	if resp != "Y" {
		return pagination
	}

	v := interaction.GetUserInputWithErrorHandling("How many URLs would you like to check?")

	number, err := strconv.Atoi(v)

	if err != nil {
		log.Fatalf("Invalid number for pagination, error: %s", err)
	}

	pagination.Tag = *getTagData(
		"where to extract the pagination (it should be the parent of pagination)",
		"Insert selector name where to get the pagination link",
	)

	pagination.Number = number

	return pagination
}

func processImages(url string, t *htmlprocesser.Tag, pagination *htmlprocesser.Pagination) htmlprocesser.ImageUrls {
	var imageUrls htmlprocesser.ImageUrls

	for i := 0; i < pagination.Number; i++ {

		htmlParsed, err := htmlprocesser.GetHTMLParsed(url)

		if err != nil {
			log.Printf("Error parsing HTML: %s", err)
			continue
		}

		divClass, err := htmlprocesser.GetDivBySelector(string(t.Selector.Type), string(t.Selector.Name), htmlParsed)

		if err != nil {
			log.Printf("Error parsing div: %s", err)
			continue
		}

		imageUrls = append(imageUrls, htmlprocesser.GetImageLinksFrom(divClass)...)

		if pagination.Number > 1 {
			url, err = htmlprocesser.GetPaginationNextLink(htmlParsed, string(pagination.Tag.Selector.Name))
			if err != nil {
				log.Fatalf("there was en error when trying to get the next link in the pagination, error was %s", err.Error())
			}
		}
	}

	return imageUrls
}

func getTagData(tagOptionText string, selectorTypeText string) *htmlprocesser.Tag {
	tag := &htmlprocesser.Tag{
		Selector: htmlprocesser.Selector{
			Type: "",
			Name: "",
		},
		Name: "",
	}

	interaction.ShowTagOptions(tagOptionText)

	tag.Name = interaction.GetTagChoice()

	interaction.ShowSelectorOptions(fmt.Sprintf("for the %s", string(tag.Name)))

	tag.Selector.Type = interaction.GetSelectorTypeChoice()

	selectorName := interaction.GetUserInputWithErrorHandling(selectorTypeText)

	tag.Selector.Name = htmlprocesser.SelectorName(selectorName)

	return tag
}

func main() {

	tagToExtractImage := getTagData(
		"where to extract the img (it should be the parent of <img>)",
		"Insert selector name where to pull all the images",
	)

	imgChannel := make(chan bool)

	imagesToProcess := 0

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

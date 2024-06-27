package main

import (
	"log"
	"strconv"

	"github.com/hreluz/images-scrapper/internal/interaction"
	htmlprocesser "github.com/hreluz/images-scrapper/pkg/html_processer"
	imagedownloader "github.com/hreluz/images-scrapper/pkg/image_downloader"
)

func getPagination() (int, string) {
	resp := interaction.GetUserInputWithErrorHandling("Does this URL have pagination (Y/N)?")

	if resp != "Y" {
		return 0, ""
	}

	v := interaction.GetUserInputWithErrorHandling("How many URLs would you like to check?")

	number, err := strconv.Atoi(v)

	if err != nil {
		log.Fatalf("Invalid number for pagination, error: %s", err)
	}

	className := interaction.GetUserInputWithErrorHandling("Enter the classname for the pagination")

	return number, className
}

func processImages(url string, className string, pNumber int, pClassName string) htmlprocesser.ImageUrls {
	var imageUrls htmlprocesser.ImageUrls

	for i := 0; i < pNumber; i++ {

		htmlParsed, err := htmlprocesser.GetHTMLParsed(url)

		if err != nil {
			log.Printf("Error parsing HTML: %s", err)
			continue
		}

		divClass, err := htmlprocesser.GetDivByClass(className, htmlParsed)

		if err != nil {
			log.Printf("Error parsing div: %s", err)
			continue
		}

		imageUrls = append(imageUrls, htmlprocesser.GetImageLinksFrom(divClass)...)

		url = htmlprocesser.GetPaginationNextLink(htmlParsed, pClassName)
	}

	return imageUrls
}

func main() {

	imgChannel := make(chan bool)

	imagesToProcess := 0

	className := interaction.GetUserInputWithErrorHandling("Insert class name where to pull all the images")

	url := interaction.GetUserInputWithErrorHandling("Insert URL")

	pNumber, pClassName := getPagination()

	imageUrls := processImages(url, className, pNumber, pClassName)

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

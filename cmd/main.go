package main

import (
	"github.com/hreluz/images-scrapper/internal/interaction"
	"github.com/hreluz/images-scrapper/pkg/html_processer/image"
	imagedownloader "github.com/hreluz/images-scrapper/pkg/image_downloader"
)

func main() {

	imageUrlsChannel := make(chan string)
	webUrlsChannel := make(chan string)
	imgChannel := make(chan bool)

	// Get user input and configuration
	webUrl := interaction.GetUserInputWithErrorHandling("Insert URL")
	paginationConfig := interaction.GetPagination()
	imageConfig := interaction.GetTagConfig("Insert how many levels the img will have")

	// Initialize the image processor and downloader
	iprocessor := image.NewProcessor(imageConfig, paginationConfig)
	id := &imagedownloader.ImageDownloader{
		Download_folder_path: "../downloaded_images",
		Img_channel:          imgChannel,
		Prefix_image:         "image_",
	}

	// Launch goroutine for URL processing
	go func() {
		webUrlsChannel <- webUrl
	}()

	for i := 0; i < paginationConfig.GetNumber(); i++ {

		go func() {
			im := image.Process(iprocessor, <-webUrlsChannel)

			imageUrlsChannel <- im.GetUrl()

			if paginationConfig.GetNumber() > 1 {
				webUrlsChannel <- im.GetNextUrl()
			}
		}()
		go imagedownloader.Download(id, <-imageUrlsChannel)
	}

	for i := 0; i < paginationConfig.GetNumber(); i++ {
		<-imgChannel
	}
}

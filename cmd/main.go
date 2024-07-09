package main

import (
	"fmt"

	"github.com/hreluz/images-scrapper/internal/interaction"
	imagedownloader "github.com/hreluz/images-scrapper/pkg/image_downloader"
	"github.com/hreluz/images-scrapper/pkg/models/images"
	"github.com/hreluz/images-scrapper/pkg/services"
)

func main() {

	imageUrlsChannel := make(chan string)
	webUrlsChannel := make(chan string)
	imgChannel := make(chan bool)
	var imagesCollection images.Images

	// Get user input and configuration
	webUrl := interaction.GetUserInputWithErrorHandling("Insert URL")
	paginationConfig := interaction.GetPagination()
	imageConfig := interaction.GetTagConfig("Insert how many levels the img will have")
	titleConfig := interaction.GetTitle()
	descriptionConfig := interaction.GetDescription()

	// Initialize the image processor and downloader
	iprocessor := images.NewProcessor(imageConfig, paginationConfig, titleConfig, descriptionConfig)
	iservice := services.NewImageService(iprocessor)

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
			im := services.ProcessImage(iservice(<-webUrlsChannel))

			imagesCollection = append(imagesCollection, *im)

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

	for _, i := range imagesCollection {
		fmt.Println(i.String())
	}
}

package main

import (
	"github.com/hreluz/images-scrapper/internal/interaction"
	"github.com/hreluz/images-scrapper/pkg/html_processer/image"
	imagedownloader "github.com/hreluz/images-scrapper/pkg/image_downloader"
)

func main() {

	var imageUrls image.ImageUrls

	url := interaction.GetUserInputWithErrorHandling("Insert URL")

	paginationConfig := interaction.GetPagination()

	imageConfig := interaction.GetTagConfig("Insert how many levels the img will have: ")

	iprocessor := image.NewProcessor(imageConfig, paginationConfig)

	imgChannel := make(chan bool)

	id := &imagedownloader.ImageDownloader{
		Download_folder_path: "../downloaded_images",
		Img_channel:          imgChannel,
		Prefix_image:         "image_",
	}

	for i := 0; i < paginationConfig.GetNumber(); i++ {
		im := image.Process(iprocessor, url)

		imageUrls = append(imageUrls, im.GetUrl())

		if paginationConfig.GetNumber() > 1 {
			url = im.GetNextUrl()
		}
	}

	for _, iu := range imageUrls {
		go imagedownloader.Download(id, iu)
	}

	for i := 0; i < paginationConfig.GetNumber(); i++ {
		<-imgChannel
	}
}

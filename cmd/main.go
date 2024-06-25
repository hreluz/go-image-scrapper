package main

import (
	"log"

	"github.com/hreluz/images-scrapper/internal/interaction"
	htmlprocesser "github.com/hreluz/images-scrapper/pkg/html_processer"
	imagedownloader "github.com/hreluz/images-scrapper/pkg/image_downloader"
)

func userInput() (string, string) {
	// Get user's url
	url, err := interaction.GetUserInput("Insert url: ")

	if err != nil {
		log.Fatalf("There was an error when saving the url, error: %s", err)
	}

	// Get user's classname
	class_name, err := interaction.GetUserInput("Insert class name where to pull all the images: ")

	if err != nil {
		log.Fatalf("There was an error when saving the classnames, error: %s", err)
	}

	return url, class_name
}

func main() {
	url, class_name := userInput()
	img_channel := make(chan bool)
	to_process := 0
	image_urls := htmlprocesser.ExecByClass(url, class_name)

	id := &imagedownloader.ImageDownloader{
		Download_folder_path: "../downloaded_images",
		Img_channel:          img_channel,
		Prefix_image:         "image_",
	}

	for _, image_url := range image_urls {
		to_process++
		go imagedownloader.Download(id, image_url)
	}

	for i := 0; i < to_process; i++ {
		<-img_channel
	}
}

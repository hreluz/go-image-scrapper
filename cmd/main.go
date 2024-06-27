package main

import (
	"log"
	"strconv"

	"github.com/hreluz/images-scrapper/internal/interaction"
	htmlprocesser "github.com/hreluz/images-scrapper/pkg/html_processer"
	imagedownloader "github.com/hreluz/images-scrapper/pkg/image_downloader"
)

func getUrl() string {
	url, err := interaction.GetUserInput("Insert url")

	if err != nil {
		log.Fatalf("There was an error when getting the url, error: %s", err)
	}

	return url
}

func getClassNames() string {
	class_name, err := interaction.GetUserInput("Insert class name where to pull all the images")

	if err != nil {
		log.Fatalf("There was an error when getting the classnames, error: %s", err)
	}

	return class_name
}

func getPagination() (number int, class_name string) {
	resp, err := interaction.GetUserInput("Does this url have pagination (Y/N)?")

	if err != nil {
		log.Fatalf("There was an error when getting the pagination, error: %s", err)
	}

	if resp == "Y" {
		v, _ := interaction.GetUserInput("How many urls would you like to check?")

		if _, err := strconv.Atoi(v); err == nil {
			number, _ = strconv.Atoi(v)
		}

		class_name, err = interaction.GetUserInput("Enter the classname for the pagination")

		if err != nil {
			log.Fatalf("There was an error when getting the clasname for the pagination, error: %s", err)
		}
	}

	return number, class_name
}

func main() {

	var image_urls htmlprocesser.ImageUrls

	img_channel := make(chan bool)

	images_to_process := 0

	class_name := getClassNames()

	url := getUrl()

	p_number, p_class_name := getPagination()

	id := &imagedownloader.ImageDownloader{
		Download_folder_path: "../downloaded_images",
		Img_channel:          img_channel,
		Prefix_image:         "image_",
	}

	for i := 0; i < p_number; i++ {

		html_parsed, _ := htmlprocesser.GetHTMLParsed(url)

		div_class, _ := htmlprocesser.GetDivByClass(class_name, html_parsed)

		image_urls = append(image_urls, htmlprocesser.GetImageLinksFrom(div_class)...)

		url = htmlprocesser.GetPaginationNextLink(html_parsed, p_class_name)
	}

	for _, image_url := range image_urls {
		images_to_process++
		go imagedownloader.Download(id, image_url)
	}

	for i := 0; i < images_to_process; i++ {
		<-img_channel
	}
}

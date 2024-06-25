package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/hreluz/images-scrapper/interaction"
)

func getFilename(image_url string) string {
	v := strings.Split(image_url, "/")
	return v[len(v)-1]
}

func download_image(image_url string, img_channel chan bool) {
	response, e := http.Get(image_url)
	if e != nil {
		log.Fatal(e)
	}

	defer response.Body.Close()

	file, err := os.Create("downloaded_images/" + getFilename(image_url))

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Success!")
	img_channel <- true
}

func main() {
	// Get user's url
	url, err := interaction.GetUserInput("Insert url: ")

	if err != nil {
		log.Printf("There was an error when saving the url, error: %s", err)
	}

	// Get user's classname
	class_name, err := interaction.GetUserInput("Insert class name where to pull all the images: ")

	if err != nil {
		log.Printf("There was an error when saving the classnames, error: %s", err)
	}

	img_channel := make(chan bool)
	to_process := 0

	// Get url html content
	resp, err := soup.Get(url)

	if err != nil {
		fmt.Printf("There was an error resolving the url %v", err)
		os.Exit(1)
	}

	doc := soup.HTMLParse(resp)

	// Find the classname div
	div_class := doc.Find("div", "class", class_name)

	if div_class.Error != nil {
		fmt.Printf("Class provided (%s) does not exist, error: %v", class_name, div_class.Error)
		os.Exit(1)
	}

	// Find all the images in that div
	images_tags := div_class.FindAll("img")

	for _, image_tag := range images_tags {
		image_url := image_tag.Attrs()["src"]
		go download_image(image_url, img_channel)
		fmt.Println("Image link :", image_url)
		to_process++
	}

	for i := 0; i < 1; i++ {
		<-img_channel
	}
}

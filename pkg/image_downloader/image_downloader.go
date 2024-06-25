package imagedownloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type ImageDownloader struct {
	Download_folder_path string
	Img_channel          chan bool
	Prefix_image         string
}

func getFilename(image_url, prefix string) string {
	v := strings.Split(image_url, "/")
	return prefix + v[len(v)-1]
}

func Download(id *ImageDownloader, image_url string) {
	response, e := http.Get(image_url)
	if e != nil {
		log.Fatalf("%v", e)
	}

	defer response.Body.Close()

	filename := getFilename(image_url, id.Prefix_image)
	file, err := os.Create(id.Download_folder_path + "/" + filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The image %v was downloaded successfully!", filename)
	id.Img_channel <- true
}

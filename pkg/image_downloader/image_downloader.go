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

func getFilename(image_url, prefix string) (string, error) {
	v := strings.Split(image_url, "/")
	s_length := len(v)

	if s_length <= 1 {
		return "", fmt.Errorf("filename was not found on image url %s", image_url)
	}

	return prefix + v[s_length-1], nil
}

func Download(id *ImageDownloader, image_url string) {
	response, e := http.Get(image_url)

	if e != nil {
		log.Fatalf("%v", e)
	}

	defer response.Body.Close()

	filename, err := getFilename(image_url, id.Prefix_image)

	if err != nil {
		log.Fatalf("Error on getting filename, error: %v", err)
	}

	file, err := os.Create(id.Download_folder_path + "/" + filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The image %v was downloaded successfully!\n", filename)
	id.Img_channel <- true
}

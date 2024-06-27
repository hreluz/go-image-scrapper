package imagedownloader

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type StorageMockError struct{}

func TestGetFilename_is_correct(t *testing.T) {
	got, err := getFilename("http://something.com/images/some-image.jpg", "thumbnail_")
	expected := "thumbnail_some-image.jpg"

	if err != nil || got != expected {
		t.Errorf("Error: It was expected %s got %s", expected, got)
	}
}

func TestGetFilename_is_incorrect(t *testing.T) {
	got, err := getFilename("", "thumbnail_")

	if err == nil {
		t.Errorf("Error: It was expected an error, but got %s", got)
	}
}

func deleteFileIfExists(filepath string) {
	if _, err := os.Stat(filepath); err == nil {
		os.Remove(filepath)
	}
}

func TestDownload(t *testing.T) {
	channel := make(chan bool)
	filename := "gopher_scrapper.jpg"

	id := &ImageDownloader{
		Download_folder_path: "../../downloaded_images",
		Img_channel:          channel,
		Prefix_image:         "image_",
	}

	filepath := id.Download_folder_path + "/" + id.Prefix_image + filename
	defer deleteFileIfExists(filepath)

	imageHandler := func(w http.ResponseWriter, r *http.Request) {
		imageData, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Fatalf("Failed to read image file: %v", err)
		}

		w.Header().Set("Content-Type", "image/jpeg")
		w.Write(imageData)
	}

	mockServer := httptest.NewServer(http.HandlerFunc(imageHandler))

	defer mockServer.Close()

	url := mockServer.URL + "/" + filename

	go Download(id, url)

	<-channel

	if _, err := os.Stat(filepath); err != nil {
		t.Errorf("Error: The image was not found in the downloaded_images folder")
	}
}

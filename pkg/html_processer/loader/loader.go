package loader

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/anaskhan96/soup"
)

func GetHTMlContent(htmlName string) []byte {
	htmlFilePath := fmt.Sprintf("testdata/%s.html", htmlName)

	htmlContent, err := ioutil.ReadFile(htmlFilePath)

	if err != nil {
		log.Fatalf("Failed to read HTML file: %v", err)
	}

	return htmlContent
}

func GetHTMLParsed(url string) (soup.Root, error) {
	resp, err := soup.Get(url)

	if err != nil {
		return soup.Root{}, fmt.Errorf("there was an error resolving the url: %s, %v", url, err)
	}

	doc := soup.HTMLParse(resp)

	return doc, nil
}

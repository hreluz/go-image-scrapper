package htmlprocesser

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/anaskhan96/soup"
)

func getHTMlContent(htmlName string) []byte {
	htmlFilePath := fmt.Sprintf("testdata/%s.html", htmlName)

	htmlContent, err := ioutil.ReadFile(htmlFilePath)

	if err != nil {
		log.Fatalf("Failed to read HTML file: %v", err)
	}

	return htmlContent
}

func TestGetHTML(t *testing.T) {

	htmlContent := getHTMlContent("no_content")

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(htmlContent)
	}))

	defer mockServer.Close()

	url := mockServer.URL
	html, err := GetHTMLParsed(url)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := len(soup.HTMLParse(string(htmlContent)).HTML())
	got := len(html.HTML())

	if expected != got {
		t.Fatalf("html content length does not match, it was expected %d, it got %d", expected, got)
	}
}

func TestGetHTML_cannot_load(t *testing.T) {
	_, err := GetHTMLParsed("http://localhost")

	if err == nil {
		t.Fatalf("an error happened, it resolved url when it shouldn't")
	}
}

func TestGetDivBySelector(t *testing.T) {
	htmlContent := getHTMlContent("id_container")
	htmlParsed := soup.HTMLParse(string(htmlContent))
	div, _ := GetDivBySelector("id", "container-image", htmlParsed)
	got := string(div.HTML())

	expected := `<div id="container-image">
        <p>hello</p>
    </div>`

	if got != expected {
		t.Fatalf("div did not match, it was expected %s, and got %s", expected, got)
	}
}

func TestGetDivBySelector_was_not_found(t *testing.T) {
	htmlContent := getHTMlContent("no_content")
	htmlParsed := soup.HTMLParse(string(htmlContent))
	_, err := GetDivBySelector("id", "container-image", htmlParsed)

	if err == nil {
		t.Fatalf("div was found when it shouldn't")
	}
}

func TestGetImageLinksFrom(t *testing.T) {
	htmlContent := getHTMlContent("container_with_urls")
	htmlParsed := soup.HTMLParse(string(htmlContent))
	expected := ImageUrls{"http://somewhere.com/image1.jpeg", "http://somewhere.com/image15.jpeg", "http://somewhere.com/image2.jpeg"}
	got := GetImageLinksFrom(htmlParsed)

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("image links did not match, it was expected %v, it got %v", expected, got)
	}
}

func TestGetImageLinksFrom_when_there_are_no_links(t *testing.T) {
	htmlContent := getHTMlContent("no_content")
	htmlParsed := soup.HTMLParse(string(htmlContent))
	expected := ImageUrls{}
	got := GetImageLinksFrom(htmlParsed)

	if reflect.DeepEqual(got, expected) {
		t.Fatalf("image links was not empty, whhen it should, it was expected %v, it got %v", expected, got)
	}
}

func TestGetPaginationNextLink(t *testing.T) {
	htmlContent := getHTMlContent("pagination")
	htmlParsed := soup.HTMLParse(string(htmlContent))
	expected := "http://something-else.com/10"
	got, err := GetPaginationNextLink(htmlParsed, "previous")

	if err != nil {
		log.Fatalf("there was an error on the pagination, error: %s", err)
	}

	if expected != got {
		t.Fatalf("pagination link did not match, it was expected %v, it got %v", expected, got)
	}
}

func TestGetPaginationNextLink_was_not_found(t *testing.T) {
	htmlContent := getHTMlContent("no_content")
	htmlParsed := soup.HTMLParse(string(htmlContent))
	_, err := GetPaginationNextLink(htmlParsed, "previous")
	expected := "span with class previous was not found in html"

	if err.Error() != expected {
		log.Fatalf("the error '%s' was expected, but got %s", expected, err)
	}
}

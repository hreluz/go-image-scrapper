package image

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/anaskhan96/soup"
	"github.com/hreluz/images-scrapper/pkg/html_processer/selector"
	"github.com/hreluz/images-scrapper/pkg/html_processer/tag"
)

func TestGetImageLink(t *testing.T) {

	content := `<p><img src="http://link.com/image.jpeg" alt=""></p>`

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(content))
	}))

	defer mockServer.Close()
	url := mockServer.URL + "/hello"

	tag1 := tag.New(selector.Empty(), tag.P)
	tag2 := tag.New(selector.Empty(), tag.IMG)
	tagConfig := tag.NewConfig(5, []tag.Tag{*tag1, *tag2})

	ip := NewProcessor(tagConfig, nil, nil, nil)
	i := NewImage(ip, url)

	expected := `http://link.com/image.jpeg`
	got := i.GetUrl()

	if got != expected {
		t.Fatalf("the image src for processing on image did not match, it was expected \n%s, and got \n%s", expected, got)
	}
}

func TestGetImageString(t *testing.T) {

	i := Image{
		HTMLParsed:  soup.Root{},
		url:         "http://web.com/image.jpeg",
		webUrl:      "http://web.com",
		processed:   true,
		nextUrl:     "",
		title:       "A title",
		description: "A description",
	}

	expectedString := `
------------------------------------
Web Url: http://web.com
Image Url: http://web.com/image.jpeg
Title: A title
Description: A description
------------------------------------`
	got := i.String()

	if normalizeString(expectedString) != normalizeString(got) {
		t.Fatalf("Image strings does not match, it was expected %s, it got %s", expectedString, got)
	}
}

func normalizeString(s string) string {
	// Trim leading and trailing whitespace
	s = strings.TrimSpace(s)
	// Convert all line endings to \n
	s = strings.ReplaceAll(s, "\r\n", "\n")
	return s
}

package loader

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anaskhan96/soup"
)

func TestGetHTML(t *testing.T) {

	htmlContent := GetHTMlContent("no_content")

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

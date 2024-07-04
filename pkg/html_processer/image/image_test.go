package image

import (
	"net/http"
	"net/http/httptest"
	"testing"

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

	ip := NewProcessor(tagConfig, nil)
	i := Process(ip, url)

	expected := `http://link.com/image.jpeg`
	got := i.GetUrl()

	if got != expected {
		t.Fatalf("the image src for processing on image did not match, it was expected \n%s, and got \n%s", expected, got)
	}
}

package tag

import (
	"testing"

	"github.com/anaskhan96/soup"
	"github.com/hreluz/images-scrapper/pkg/html_processer/loader"
	"github.com/hreluz/images-scrapper/pkg/html_processer/selector"
)

func TestGetContentByTag(t *testing.T) {
	htmlContent := loader.GetHTMlContent("id_container")
	htmlParsed := soup.HTMLParse(string(htmlContent))
	s := selector.New(
		selector.ID,
		"container-image",
	)

	tag := New(s, DIV)

	selector, _ := tag.GetContentByTag(htmlParsed)
	got := string(selector.HTML())
	expected := `<div id="container-image">
        <p>hello</p>
    </div>`

	if got != expected {
		t.Fatalf("div did not match, it was expected\n %s, and got\n %s", expected, got)
	}
}

func TestGetContentByTag_was_not_found(t *testing.T) {
	htmlContent := loader.GetHTMlContent("id_container")
	htmlParsed := soup.HTMLParse(string(htmlContent))
	s := selector.New(
		selector.CLASS,
		"some-class-that-does-not-exist",
	)

	tag := New(s, DIV)

	_, err := tag.GetContentByTag(htmlParsed)

	if err == nil {
		t.Fatalf("div was found when it shouldn't")
	}
}

func TestGetContentByTag_when_selector_is_empty(t *testing.T) {
	htmlContent := loader.GetHTMlContent("id_container")
	htmlParsed := soup.HTMLParse(string(htmlContent))
	tag := New(selector.Empty(), DIV)
	selector, err := tag.GetContentByTag(htmlParsed)

	if err != nil {
		t.Fatalf("there was an error %v", err)
	}

	got := string(selector.HTML())
	expected := `<div id="container-image">
        <p>hello</p>
    </div>`

	if got != expected {
		t.Fatalf("div did not match, it was expected\n %s, and got\n %s", expected, got)
	}
}

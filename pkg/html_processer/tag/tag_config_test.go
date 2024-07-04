package tag

import (
	"testing"

	"github.com/anaskhan96/soup"
	"github.com/hreluz/images-scrapper/pkg/html_processer/loader"
	"github.com/hreluz/images-scrapper/pkg/html_processer/selector"
)

func TestGetLastTagContainer(t *testing.T) {
	htmlContent := loader.GetHTMlContent("nested")
	htmlParsed := soup.HTMLParse(string(htmlContent))

	tag1 := New(selector.New(selector.ID, "first-container"), DIV)
	tag2 := New(selector.New(selector.CLASS, "second-container"), DIV)
	tag3 := New(selector.New(selector.ID, "third"), ARTICLE)
	tag4 := New(selector.New(selector.CLASS, "fourth-div"), DIV)
	tag5 := New(selector.New(selector.CLASS, "last-tag"), DIV)
	tagConfig := NewConfig(5, []Tag{*tag1, *tag2, *tag3, *tag4, *tag5})

	selector, _ := tagConfig.GetLastTagContainer(htmlParsed)
	got := string(selector.HTML())
	expected := `<div class="last-tag"><p>hello</p></div>`

	if got != expected {
		t.Fatalf("the last container didn't match the expected, it was expected \n%s, and got \n%s", expected, got)
	}
}

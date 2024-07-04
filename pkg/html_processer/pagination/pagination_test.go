package pagination

import (
	"testing"

	"github.com/anaskhan96/soup"
	"github.com/hreluz/images-scrapper/pkg/html_processer/loader"
	"github.com/hreluz/images-scrapper/pkg/html_processer/selector"
	"github.com/hreluz/images-scrapper/pkg/html_processer/tag"
)

func TestGetPaginationNextLink(t *testing.T) {
	htmlContent := loader.GetHTMlContent("nested_pagination")
	htmlParsed := soup.HTMLParse(string(htmlContent))

	tag1 := tag.New(selector.New(selector.ID, "first-container"), tag.DIV)
	tag2 := tag.New(selector.New(selector.CLASS, "second-container"), tag.DIV)
	tag3 := tag.New(selector.New(selector.ID, "third"), tag.ARTICLE)
	tag4 := tag.New(selector.New(selector.CLASS, "fourth-div"), tag.DIV)
	tag5 := tag.New(selector.New(selector.CLASS, "last-tag"), tag.DIV)
	tag6 := tag.New(selector.Empty(), tag.SPAN)
	tag7 := tag.New(selector.Empty(), tag.A)

	tagConfig := tag.NewConfig(7, []tag.Tag{*tag1, *tag2, *tag3, *tag4, *tag5, *tag6, *tag7})

	pagination := New(tagConfig, 1)

	gotLink, error := pagination.GetPaginationNextLink(htmlParsed)

	if error != nil {
		t.Fatalf("there was an error when retrieving link on pagination, error: %v", error)
	}

	expectedLink := "http://something-else.com/10"

	if gotLink != expectedLink {
		t.Fatalf("pagination link did not match, it was expected \n%s, and got \n%s", expectedLink, gotLink)
	}
}

func TestGetPaginationNextLink_was_not_found(t *testing.T) {
	htmlContent := loader.GetHTMlContent("nested_pagination")
	htmlParsed := soup.HTMLParse(string(htmlContent))

	tag1 := tag.New(selector.New(selector.ID, "invented-id"), tag.DIV)
	tag2 := tag.New(selector.Empty(), tag.A)

	tagConfig := tag.NewConfig(2, []tag.Tag{*tag1, *tag2})

	pagination := New(tagConfig, 1)

	_, error := pagination.GetPaginationNextLink(htmlParsed)

	expectedError := "there was an error trying to find the selector id with the name invented-id"

	if error.Error() != expectedError {
		t.Fatalf("it was expected the error %v, but got %v", expectedError, error)
	}
}

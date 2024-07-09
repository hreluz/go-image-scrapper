package images

import (
	"strings"
	"testing"
)

func TestGetImageString(t *testing.T) {

	i := Image{
		html:   "",
		url:    "http://web.com/image.jpeg",
		webUrl: "http://web.com",
		// processed:   true,
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

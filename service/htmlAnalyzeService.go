package service

import (
	"github.com/pandulaDW/home24-page-analyzer/models"
	"github.com/pandulaDW/home24-page-analyzer/parsers"
	"golang.org/x/net/html"
	"io"
)

// HtmlPageDetails returns the required details extracted from the given html
// page tokenizer.
func HtmlPageDetails(r io.Reader) *models.HTMLPageDetails {
	tokenizer := html.NewTokenizer(r)
	var title string
	var headingCount models.HeadingCount
	var linkCount models.LinkCount

	for tokenType := tokenizer.Next(); tokenType != html.ErrorToken; {
		currentTagInBytes, _ := tokenizer.TagName()
		currentTag := string(currentTagInBytes)

		// extracting the title
		if currentTag == "title" && tokenType == html.StartTagToken {
			tokenType = tokenizer.Next()
			title = tokenizer.Token().String()
		}

		// extracting the heading and link information
		if tokenType == html.StartTagToken {
			parsers.CountHeadings(currentTag, &headingCount)
			parsers.GetLinkInformation(tokenizer, currentTag, &linkCount)
		}

		// iterating to the next node
		tokenType = tokenizer.Next()
	}

	return &models.HTMLPageDetails{
		Title:        title,
		HeadingCount: headingCount,
		LinkCount:    linkCount,
	}
}

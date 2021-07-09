package service

import (
	"github.com/pandulaDW/home24-page-analyzer/models"
	"github.com/pandulaDW/home24-page-analyzer/parsers"
	"golang.org/x/net/html"
)

// HtmlPageDetails returns the required details extracted from the given html
// page tokenizer.
func HtmlPageDetails(doc *html.Tokenizer) *models.HTMLPageDetails {
	var title string
	var headingCount models.HeadingCount

	for tokenType := doc.Next(); tokenType != html.ErrorToken; {
		currentTagBytes, _ := doc.TagName()
		currentTag := string(currentTagBytes)

		// extracting the title
		if string(currentTag) == "title" && tokenType == html.StartTagToken {
			tokenType = doc.Next()
			title = doc.Token().String()
		}

		// extracting the heading tags
		if tokenType == html.StartTagToken {
			parsers.CountHeadings(currentTag, &headingCount)
		}

		// iterating to the next node
		tokenType = doc.Next()
	}

	return &models.HTMLPageDetails{
		Title:        title,
		HeadingCount: headingCount,
	}
}

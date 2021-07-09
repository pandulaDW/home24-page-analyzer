package service

import (
	"github.com/pandulaDW/home24-page-analyzer/models"
	"github.com/pandulaDW/home24-page-analyzer/parsers"
	"golang.org/x/net/html"
	"io"
)

// HtmlPageDetails returns the required details extracted from the given html page.
func HtmlPageDetails(r io.Reader) *models.HTMLPageDetails {
	tokenizer := html.NewTokenizer(r)
	externalLinks := make([]string, 0)

	var headingCount models.HeadingCount
	var linkCount models.LinkCount

	outputModel := new(models.HTMLPageDetails)

	for tokenType := tokenizer.Next(); tokenType != html.ErrorToken; {
		currentTagInBytes, _ := tokenizer.TagName()
		currentTag := string(currentTagInBytes)

		// extracting the title
		if currentTag == "title" && tokenType == html.StartTagToken {
			tokenType = tokenizer.Next()
			outputModel.Title = tokenizer.Token().String()
		}

		// extracting the heading and link information
		if tokenType == html.StartTagToken {
			parsers.CountHeadings(currentTag, &headingCount)

			link := parsers.GetLinkInformation(tokenizer, currentTag, &linkCount)
			if link != "" {
				externalLinks = append(externalLinks, link)
			}
		}

		// checking if a login element exists inside a form
		if currentTag == "form" && tokenType == html.StartTagToken {
			for {
				tokenType = tokenizer.Next()
				if tag, _ := tokenizer.TagName(); string(tag) == "form" && tokenType == html.EndTagToken {
					break
				}
				if tokenType == html.TextToken && parsers.IsLogin(tokenizer.Token().String()) {
					outputModel.IsLoginForm = true
					break
				}
				if tokenType == html.ErrorToken {
					break
				}
			}
		}

		// iterating to the next node
		tokenType = tokenizer.Next()
	}

	// get inaccessible link count
	linkCount.InaccessibleLinkCount = checkAccessibility(externalLinks)

	outputModel.HeadingCount = headingCount
	outputModel.LinkCount = linkCount
	return outputModel
}

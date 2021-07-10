package services

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
	outputModel := new(models.HTMLPageDetails)

	for tokenType := tokenizer.Next(); tokenType != html.ErrorToken; {
		// extracting the html version
		if tokenType == html.DoctypeToken {
			outputModel.HTMLVersion = parsers.GetHTMLVersion(tokenizer.Token().Data)
		}

		currentTagInBytes, _ := tokenizer.TagName()
		currentTag := string(currentTagInBytes)

		// extracting the title
		if currentTag == "title" && tokenType == html.StartTagToken {
			tokenType = tokenizer.Next()
			outputModel.Title = tokenizer.Token().String()
		}

		// extracting the heading and link information
		if tokenType == html.StartTagToken {
			parsers.CountHeadings(currentTag, &outputModel.HeadingCount)

			link := parsers.GetLinkInformation(tokenizer, currentTag, &outputModel.LinkCount)
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
	outputModel.LinkCount.InaccessibleLinkCount = checkAccessibility(externalLinks)

	return outputModel
}

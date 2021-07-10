package parsers

import (
	"github.com/pandulaDW/home24-page-analyzer/models"
	"golang.org/x/net/html"
	"regexp"
)

var (
	// internalLinkRegex matches the internal links
	internalLinkRegex = regexp.MustCompile(`^/[\w.=\-/]+`)

	// externalLinkRegex matches the external links
	externalLinkRegex = regexp.MustCompile(
		`https?://(www\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_+.~#?&=]*)`)
)

// GetLinkInformation set the information regarding the currently scanned tag if it's an anchor tag.
//
// The function also returns the link if the link is an external link for checking accessibility later.
func GetLinkInformation(tokenizer *html.Tokenizer, currentTag string, count *models.LinkCount) string {
	var externalLink string

	if currentTag == "a" {
		link := getLinkUrl(tokenizer)
		if isInternalLink(link) {
			count.InternalLinkCount++
		} else if isExternalLink(link) {
			count.ExternalLinkCount++
			externalLink = link
		}
	}

	return externalLink
}

// getLinkUrl examines the current analyzed token and returns the url of the
// anchor attribute
func getLinkUrl(tokenizer *html.Tokenizer) string {
	var url string

	for {
		attr, val, moreAttr := tokenizer.TagAttr()
		if string(attr) == "href" {
			url = string(val)
			break
		}
		if !moreAttr {
			break
		}
	}

	return url
}

// isInternalLink returns true if the given link url is an internal link, false otherwise
func isInternalLink(url string) bool {
	return internalLinkRegex.MatchString(url)
}

// isExternalLink returns true if the given link url is an external link, false otherwise
func isExternalLink(url string) bool {
	return externalLinkRegex.MatchString(url)
}

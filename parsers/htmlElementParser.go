package parsers

import (
	"github.com/pandulaDW/home24-page-analyzer/models"
	"regexp"
)

var (
	isLoginRegex = regexp.MustCompile(`(?i)^(login|sign)(\sin)?$`)
	isHtml5      = regexp.MustCompile(`(?i)^html$`)
	isHtml4      = regexp.MustCompile(`(?i)^.*/html4/.*$`)
	isXhtml      = regexp.MustCompile(`(?i)^.*/(xhtml1|xhtml11)/.*$`)
)

// GetHTMLVersion will check the doctype text attribute and assess whether the correct
// html version.
//
// If a match is not found, it will return models.UNKNOWN
func GetHTMLVersion(text string) models.HTMLVersion {
	if isHtml5.MatchString(text) {
		return models.HTML5
	}
	if isHtml4.MatchString(text) {
		return models.HTML4
	}
	if isXhtml.MatchString(text) {
		return models.XHTML
	}
	return models.UNKNOWN
}

// CountHeadings will increment the heading count by examining the given node
func CountHeadings(currentTag string, count *models.HeadingCount) {
	switch currentTag {
	case "h1":
		count.H1Count++
	case "h2":
		count.H2Count++
	case "h3":
		count.H3Count++
	case "h4":
		count.H4Count++
	case "h5":
		count.H5Count++
	case "h6":
		count.H6Count++
	}
}

// IsLogin returns whether a given text corresponds to a sign in text
func IsLogin(text string) bool {
	return isLoginRegex.MatchString(text)
}

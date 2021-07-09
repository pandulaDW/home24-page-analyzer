package parsers

import (
	"github.com/pandulaDW/home24-page-analyzer/models"
	"regexp"
)

var (
	isLoginRegex = regexp.MustCompile(`(?i)^(login|sign)(\sin)?$`)
)

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

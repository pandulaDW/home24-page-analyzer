package parsers

import (
	"github.com/pandulaDW/home24-page-analyzer/models"
	"golang.org/x/net/html"
)

// CountHeadings will increment the heading count by examining the given node
func CountHeadings(node *html.Node, count *models.HeadingCount) {
	if node.Type == html.ElementNode {
		switch node.Data {
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
}

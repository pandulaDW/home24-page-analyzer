package parsers

import (
	"fmt"
	"github.com/pandulaDW/home24-page-analyzer/models"
	"golang.org/x/net/html"
)

// GetLinkInformation returns information regarding the currently scanned tag if it's an anchor tag
func GetLinkInformation(tokenizer *html.Tokenizer, currentTag string, count *models.LinkCount) {
	if currentTag == "a" {
		link := getLinkUrl(tokenizer)
		fmt.Println(link)
		count.InternalLinkCount++
	}
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

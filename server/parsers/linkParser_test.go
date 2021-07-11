package parsers

import (
	"strings"
	"testing"

	"github.com/pandulaDW/home24-page-analyzer/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestGetLinkInformation(t *testing.T) {
	// assert that internal link count is incremented correctly
	tokenizer := html.NewTokenizer(strings.NewReader(`<a class="shop-main" href="/shop/233">Shop</a>`))
	tokenizer.Next()
	currentTag, _ := tokenizer.TagName()
	linkCountObj := new(models.LinkCount)
	link := GetLinkInformation(tokenizer, string(currentTag), linkCountObj)
	assert.Equal(t, &models.LinkCount{InternalLinkCount: 1}, linkCountObj)

	// assert that blank link is returned for internal links
	assert.Empty(t, link)

	// assert that external link count is incremented correctly
	tokenizer = html.NewTokenizer(strings.NewReader(`<a class="shop-main" href="https://www.amazon.com">Shop</a>`))
	tokenizer.Next()
	currentTag, _ = tokenizer.TagName()
	linkCountObj = new(models.LinkCount)
	link = GetLinkInformation(tokenizer, string(currentTag), linkCountObj)
	assert.Equal(t, &models.LinkCount{ExternalLinkCount: 1}, linkCountObj)

	// assert that external link is returned correctly
	assert.Equal(t, "https://www.amazon.com", link)
}

func TestGetLinkURL(t *testing.T) {
	// assert that link is correctly identified
	tokenizer := html.NewTokenizer(strings.NewReader(`<a class="shop-main" href="/shop/233">Shop</a>`))
	tokenizer.Next()
	assert.Equal(t, "/shop/233", getLinkUrl(tokenizer))

	// assert that an empty string returns if href attribute is not present
	tokenizer = html.NewTokenizer(strings.NewReader(`<a class="shop-main" id="shop-id">Shop</a>`))
	tokenizer.Next()
	assert.Empty(t, getLinkUrl(tokenizer))
}

func TestIsInternalLink(t *testing.T) {
	// assert that correct internal links are identified
	assert.True(t, isInternalLink("/about"))
	assert.True(t, isInternalLink("/about/shop/123"))

	// assert that incorrect internal links are identified
	assert.False(t, isInternalLink("test/about"))
	assert.False(t, isInternalLink("https://google.com"))
}

func TestIsExternalLink(t *testing.T) {
	// assert that correct external links are identified
	assert.True(t, IsExternalLink("https://google.com"))
	assert.True(t, IsExternalLink("http://shoplify.com"))
	assert.True(t, IsExternalLink("https://www.amazon.com"))

	// assert that incorrect external links are identified
	assert.False(t, IsExternalLink("/about/shop"))
	assert.False(t, IsExternalLink("test-site"))
}

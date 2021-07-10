package parsers

import (
	"testing"

	"github.com/pandulaDW/home24-page-analyzer/models"
	"github.com/stretchr/testify/assert"
)

func TestGetHTMLVersion(t *testing.T) {
	// assert that html5 version is identified correctly
	assert.Equal(t, models.HTMLVersion(models.HTML5), GetHTMLVersion("html"))

	// assert that html4 version is identified correctly
	assert.Equal(t, models.HTMLVersion(models.HTML4),
		GetHTMLVersion(`HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd"`))
	assert.Equal(t, models.HTMLVersion(models.HTML4),
		GetHTMLVersion(`HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd"`))
	assert.Equal(t, models.HTMLVersion(models.HTML4),
		GetHTMLVersion(`HTML PUBLIC "-//W3C//DTD HTML 4.01 Frameset//EN" "http://www.w3.org/TR/html4/frameset.dtd"`))

	// assert that xhtml version is identified correctly
	assert.Equal(t, models.HTMLVersion(models.XHTML),
		GetHTMLVersion(`html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd"`))
	assert.Equal(t, models.HTMLVersion(models.XHTML),
		GetHTMLVersion(`html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd"`))
	assert.Equal(t, models.HTMLVersion(models.XHTML),
		GetHTMLVersion(`html PUBLIC "-//W3C//DTD XHTML 1.0 Frameset//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd"`))
	assert.Equal(t, models.HTMLVersion(models.XHTML),
		GetHTMLVersion(`html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd"`))

	// assert that any other version is matches as unknown
	assert.Equal(t, models.HTMLVersion(models.UNKNOWN), GetHTMLVersion("html testing"))
}

func TestCountHeadings(t *testing.T) {
	// assert that h1 element will be incremented correctly
	headingCountObj := new(models.HeadingCount)
	CountHeadings("h1", headingCountObj)
	assert.Equal(t, &models.HeadingCount{H1Count: 1}, headingCountObj)

	// assert that h2 element will be incremented correctly
	headingCountObj = new(models.HeadingCount)
	CountHeadings("h2", headingCountObj)
	assert.Equal(t, &models.HeadingCount{H2Count: 1}, headingCountObj)

	// assert that h3 element will be incremented correctly
	headingCountObj = new(models.HeadingCount)
	CountHeadings("h3", headingCountObj)
	assert.Equal(t, &models.HeadingCount{H3Count: 1}, headingCountObj)

	// assert that h4 element will be incremented correctly
	headingCountObj = new(models.HeadingCount)
	CountHeadings("h4", headingCountObj)
	assert.Equal(t, &models.HeadingCount{H4Count: 1}, headingCountObj)

	// assert that h5 element will be incremented correctly
	headingCountObj = new(models.HeadingCount)
	CountHeadings("h5", headingCountObj)
	assert.Equal(t, &models.HeadingCount{H5Count: 1}, headingCountObj)

	// assert that h6 element will be incremented correctly
	headingCountObj = new(models.HeadingCount)
	CountHeadings("h6", headingCountObj)
	assert.Equal(t, &models.HeadingCount{H6Count: 1}, headingCountObj)
}

func TestIsLogin(t *testing.T) {
	// assert that correct login formats are identified
	assert.True(t, IsLogin("Login"))
	assert.True(t, IsLogin("LOGIN"))
	assert.True(t, IsLogin("Login in"))
	assert.True(t, IsLogin("Sign IN"))
	assert.True(t, IsLogin("SIGN IN"))
	assert.True(t, IsLogin("log in"))

	// assert that incorrect formats are identified
	assert.False(t, IsLogin("test sign"))
	assert.False(t, IsLogin("sign up"))
}

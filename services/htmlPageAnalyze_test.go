package services

import (
	"github.com/pandulaDW/home24-page-analyzer/models"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestHtmlPageDetails(t *testing.T) {
	page := `<!DOCTYPE html>
			<html lang="en" translate="no">
			<head>
				<meta charset="UTF-8">
			</head>
			<title>My title</title>
			<body>
				<h1>Heading 1</h1>
				<section>New section</section>
				<h2>Heading 2</h2>
				<h2>Heading 2 second</h2>
			<ul>
				<li><a href="/about"></a></li>
			   <li><a href="https://www.google.com"></a></li>
			</ul>
			<form>
				<input type="text" name="username"/>
				<button>Sign in</button>
			</form>
			</body>
		   </html>
	`
	expected := HtmlPageDetails(strings.NewReader(page))
	actual := &models.HTMLPageDetails{
		HTMLVersion:  models.HTML5,
		Title:        "My title",
		HeadingCount: models.HeadingCount{H1Count: 1, H2Count: 2},
		LinkCount:    models.LinkCount{InternalLinkCount: 1, ExternalLinkCount: 1},
		IsLoginForm:  true,
	}

	assert.Equal(t, expected, actual)
}

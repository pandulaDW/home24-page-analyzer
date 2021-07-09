package main

import (
	"bytes"
	"fmt"
	"github.com/pandulaDW/home24-page-analyzer/models"
	"github.com/pandulaDW/home24-page-analyzer/parsers"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	start := time.Now()
	content, _ := ioutil.ReadFile("data/w3schools.html")
	doc, err := html.Parse(bytes.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}

	headingCount := models.HeadingCount{}

	var f func(*html.Node)
	f = func(n *html.Node) {
		// count the headings
		parsers.CountHeadings(n, &headingCount)

		// return the next html token recursively
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	fmt.Println(headingCount)
	fmt.Println(time.Since(start))
}

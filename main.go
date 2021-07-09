package main

import (
	"bytes"
	"fmt"
	"github.com/pandulaDW/home24-page-analyzer/service"
	"golang.org/x/net/html"
	"io/ioutil"
	"time"
)

func main() {
	start := time.Now()
	content, _ := ioutil.ReadFile("data/w3schools.html")
	doc := html.NewTokenizer(bytes.NewReader(content))

	pageDetails := service.HtmlPageDetails(doc)
	fmt.Println(pageDetails)
	fmt.Println(time.Since(start))
}

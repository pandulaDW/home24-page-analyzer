package main

import (
	"bytes"
	"fmt"
	"github.com/pandulaDW/home24-page-analyzer/service"
	"io/ioutil"
	"time"
)

func main() {
	start := time.Now()
	content, _ := ioutil.ReadFile("data/amazon.html")
	doc := bytes.NewReader(content)

	pageDetails := service.HtmlPageDetails(doc)
	fmt.Println(pageDetails)
	fmt.Println(time.Since(start))
}

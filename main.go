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
	content, _ := ioutil.ReadFile("data/w3schools.html")
	doc := bytes.NewReader(content)

	content = []byte(`<ul><li><a class="moka" href="/about">Link 1<a/></li><a href="/vision">Link 2<a/><li></li></ul>`)

	pageDetails := service.HtmlPageDetails(doc)
	fmt.Println(pageDetails)
	fmt.Println(time.Since(start))
}

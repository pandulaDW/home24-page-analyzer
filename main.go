package main

import (
	"bytes"
	"fmt"
	"github.com/pandulaDW/home24-page-analyzer/service"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	content, _ := ioutil.ReadFile("data/w3schools.html")
	doc := bytes.NewReader(content)

	pageDetails := service.HtmlPageDetails(doc)
	fmt.Println(pageDetails)
	fmt.Println(time.Since(start))

	response, err := http.Get("https://images-na.ssl-images-amazon.com/images/G/01/Recommendations/MissionExperience/BIA/bia-atc-confirm-icon._CB485946458_.png")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
	}
}

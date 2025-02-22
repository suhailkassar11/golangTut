package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome to post request")

	performPostRequest()
}

func performPostRequest() {
	const myurl = "http://youtube.com"

	requestBody := strings.NewReader(
		`{
		"courseName":"java",
		"price":0,
		"platform":"learnCodeOnline.in"

	}`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

	defer response.Body.Close()
}

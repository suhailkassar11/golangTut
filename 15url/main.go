package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://www.youtube.com:3000/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {
	fmt.Println("lets do url")

	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	qparams := result.Query()
	fmt.Println(qparams["v"])

	for _, value := range qparams {
		fmt.Println("params is", value)
	}

	//it is used for creating a url
	anotherUrl := &url.URL{
		Scheme: "https",
		Host:   "www.google.com",
		Path:   "/search",
	}

	myurl := anotherUrl.String()

	fmt.Println(myurl)
}

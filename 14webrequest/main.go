package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("lets do web request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("type of response is %T\n", response)

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(databytes))

	response.Body.Close()

}

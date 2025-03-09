package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// go greeter("hello")
	// greeter("world")
	websiteList := []string{
		"https://www.google.com",
		"https://www.bing.com",
		"https://www.duckduckgo.com",
	}
	for _, website := range websiteList {
		go getStatusCode(website)
		wg.Add(1)
	}
	wg.Wait()
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}

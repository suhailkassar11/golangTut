package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitEmpty"`
}

func main() {
	fmt.Println("welcome to json learn")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs Bootcamp", 299, "LearnCodeOnline.in", "suhail123", []string{"web-dev", "js"}},
		{"vueJs Bootcamp", 278, "LearnCodeOnline.in", "suhail234", nil},
		{"Nexts Bootcamp", 169, "LearnCodeOnline.in", "arhan123", []string{"web-dev", "js"}},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

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
	// EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	jsonDataFromWeb := []byte(`{
		"courseName": "Nexts Bootcamp",
		"Price": 169,
		"Platform": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb) //gives only boolean value

	if checkValid {
		fmt.Println("json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("json was not invalid")
	}

	//some cases where you just want to add data to key value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("%v : %v type is %T\n", k, v, v)
	}
}

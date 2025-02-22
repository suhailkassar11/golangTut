package main

import (
	"fmt"
)

func main() {
	var userName string
	fmt.Println(userName)
	fmt.Printf("type of username is %T \n", userName)

	var isBool bool = false
	fmt.Printf("type of variable is %T \n", isBool)

	var smallVal uint8 = 255
	fmt.Printf("type of variable is %T \n", smallVal)
	var smallfloat float32 = 255.4546464664488777777777777777777
	fmt.Println(smallfloat)
	fmt.Printf("type of variable is %T \n", smallfloat)
	var smallfloat1 float64 = 255.4546464664488777777777777777777
	fmt.Println(smallfloat1)
	fmt.Printf("type of variable is %T \n", smallfloat1)
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("type of variable is %T \n", anotherVariable)
	var website = "leetcode.com"
	fmt.Println(website)
	fmt.Printf("type of website is %T \n", website)

	numebrofUser := 215555
	fmt.Println(numebrofUser)
	fmt.Printf("type of numberofUser is %T \n", numebrofUser)
}

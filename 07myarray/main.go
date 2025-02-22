package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golang")
	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[2] = "peach"
	fruitList[3] = "mango"
	fmt.Println("the array is", fruitList)
}

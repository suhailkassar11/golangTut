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
	var number [5]int
	
	number[0]=1
	number[1]=2
	number[2]=8
	number[3]=6
	number[4]=5

	fmt.Println("the array is", number)

	for i:=0;i<5;i++{
		fmt.Print(number[i]," ")
	}
}

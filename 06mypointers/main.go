package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	number := 23
	var ptr = &number
	fmt.Println(ptr)
	fmt.Println(*ptr)
	*ptr = *ptr + 7
	fmt.Println(number)

}

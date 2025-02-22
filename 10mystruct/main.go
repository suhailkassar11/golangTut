package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	suhail := User{"suhail", "sk@gmail.com", true, 25}
	fmt.Println(suhail)
	fmt.Printf("the detail of suhail is: %+v \n", suhail)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

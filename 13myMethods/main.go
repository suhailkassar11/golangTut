package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	suhail := User{"suhail", "sk@gmail.com", true, 25}
	fmt.Println(suhail)
	fmt.Printf("the detail of suhail is: %+v \n", suhail)
	suhail.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

//this is the implementation of method
func (u User) GetStatus() {
	fmt.Println("status is: ", u.Status)
}

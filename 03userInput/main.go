package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

func main() {
	welcome := "welcome to user Input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello, ", name)
	age,_:=reader.ReadString('\n')
	fmt.Println("age is      ", age)
	count:=0
	fmt.Println("please enter count")
	_,err:=fmt.Scanf("%d",&count)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("counting is",count)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user Input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello, ", name)
}

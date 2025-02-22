package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("lets do conversion ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter rating between 1 to 10")
	rating, _ := reader.ReadString('\n')
	fmt.Println("the rating is ", rating)
	input, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 2 in rating", input+2)
	}
}

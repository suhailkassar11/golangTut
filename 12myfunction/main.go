package main

import "fmt"

func main() {
	country()
	fmt.Println("welcome to the function")
	var ans int = add(1, 3)
	fmt.Println(ans)
	proRes := proAdder(1, 2, 3, 4, 5, 5)
	fmt.Println(proRes)

}

func country() {
	fmt.Println("welcome to india")
}

func add(value1 int, value2 int) int {
	return value1 + value2
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

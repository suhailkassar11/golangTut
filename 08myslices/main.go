package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices")

	var fruitList = []string{"banana", "apple", "mango"}
	fmt.Println(fruitList)
	fruitList = append(fruitList, "apple")
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 1
	highScores[1] = 12
	highScores[2] = 13
	highScores[3] = 14

	fmt.Println(highScores)

	highScores = append(highScores, 43, 65, 1111, 56, 78)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//remove from slice
	var index int = 2
	highScores = append(highScores[:index], highScores[index+1:]...)
	fmt.Println(highScores)
}

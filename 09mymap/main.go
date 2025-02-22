package main

import "fmt"

func main() {
	fmt.Println("welcome to map in go")

	cities := make(map[string]string)
	cities["delhi"] = "india"
	cities["mumbai"] = "maharashtra"
	cities["kolkata"] = "kerela"
	fmt.Println(cities)
	fmt.Println(cities["delhi"])
	delete(cities, "delhi")
	fmt.Println(cities)

}

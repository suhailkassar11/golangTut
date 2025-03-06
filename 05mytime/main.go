package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("chlo time ke bare me pdhte h")
	abhikatime := time.Now()
	fmt.Println("bhi ka time h: ", abhikatime)
	fmt.Println(abhikatime.Format("02-jan-2006 15:04:05 Monday"))

}

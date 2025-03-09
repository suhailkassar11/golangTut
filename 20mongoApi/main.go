package main

import (
	"fmt"
	"log"
	"myMongoApi/router"
	"net/http"
)

func main() {
	fmt.Println("welcome to mongodb")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening port at 4000... ")
}

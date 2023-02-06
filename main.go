package main

import (
	"fmt"
	"geo-location/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Backend for RestorePhotos")

	r := router.Router()

	err := http.ListenAndServe(":4068", r)
	if err != nil {
		log.Fatal(err)
	}
}

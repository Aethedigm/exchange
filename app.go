package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting...")
	r := routes()

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

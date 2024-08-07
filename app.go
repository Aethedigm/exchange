package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting...")
	r := routes()

	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r); err != nil {
		log.Fatal(err)
	}
}

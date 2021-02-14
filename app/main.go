package main

import (
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
)

// godotenv.Load()
// os.Load()

func main() {
	// os.Load()

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Hello, 世界")
	fmt.Println(os.Getenv("TOKEN"))
}

package main

import (
	"fmt"
	"log"
	"os"

	k "github.com/jakecallery/iiria/server/keymaps"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fmt.Printf("Key: %v\n", os.Getenv(k.EnvKeyMap[k.APIkey]))

}

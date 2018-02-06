package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Empty dotenv")
	}

	c := clientcredentials.Config{}
}

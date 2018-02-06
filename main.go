package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Empty dotenv")
	}

	client, err := getClient()
	if err != nil {
		log.Fatal(err)
	}

	_, err = getMeasurements(client)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println()
}

func getClient() (*http.Client, error) {
	c := clientcredentials.Config{}
	c.ClientID = os.Getenv("CLIENT_ID")
	if c.ClientID == "" {
		return nil, errors.New("Client id not found")
	}
	c.ClientSecret = os.Getenv("CLIENT_SECRET")
	if c.ClientSecret == "" {
		return nil, errors.New("Client secret not found")
	}
	c.TokenURL = "https://api.netatmo.com/oauth2/token"
	client := c.Client(context.Background())
	return client, nil
}

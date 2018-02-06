package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
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
	c := oauth2.Config{}
	c.ClientID = os.Getenv("CLIENT_ID")
	if c.ClientID == "" {
		return nil, errors.New("Client id not found")
	}
	c.ClientSecret = os.Getenv("CLIENT_SECRET")
	if c.ClientSecret == "" {
		return nil, errors.New("Client secret not found")
	}
	c.Endpoint = oauth2.Endpoint{
		AuthURL:  "https://api.netatmo.com/oauth2/token",
		TokenURL: "https://api.netatmo.com/oauth2/token",
	}
	username := os.Getenv("USER_NAME")
	if username == "" {
		return nil, errors.New("Username secret not found")
	}
	password := os.Getenv("USER_PASSWORD")
	if password == "" {
		return nil, errors.New("Password secret not found")
	}
	ctx := context.Background()
	token, err := c.PasswordCredentialsToken(ctx, username, password)
	if err != nil {
		return nil, err
	}

	fmt.Println(token.Valid())

	return nil, nil
}

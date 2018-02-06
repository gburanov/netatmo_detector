package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type measurement struct {
	a string
}

func getMeasurements(client *http.Client) (*measurement, error) {
	response, err := client.Get("/api/getstationsdata")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	return nil, nil
}

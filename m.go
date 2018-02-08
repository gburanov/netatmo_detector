package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Jeffail/gabs"
)

type measurements map[string]moduleMeasurement

type moduleMeasurement struct {
	moduleName  string
	temperature float32
	co2         float32
	timestamp   time.Time
}

func getMeasurements(client *http.Client) (*measurements, error) {
	response, err := client.Get(apiURL + "/api/getstationsdata")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		return nil, err
	}

	devices, err := jsonParsed.S("body", "devices").ChildrenMap()
	if err != nil {
		return nil, err
	}

	for key, child := range devices {
		fmt.Printf("key: %v, value: %v\n", key, child.Data().(string))
	}

	return nil, nil
}

package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Jeffail/gabs"
)

type measurements map[string]moduleMeasurement

type moduleMeasurement struct {
	moduleName  string
	temperature float64
	co2         float64
	timestamp   time.Time
}

func (m measurements) add(measurement moduleMeasurement) {
	measurement.timestamp = time.Now()
	m[measurement.moduleName] = measurement
	return
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

	m := measurements{}

	devices, err := jsonParsed.Search("body", "devices").Children()
	if err != nil {
		return nil, err
	}

	for _, device := range devices {
		singleMeasurement, err := getMeasurement(device)
		if err != nil {
			return nil, err
		}
		m.add(singleMeasurement)

		modules, err := device.Search("modules").Children()
		if err != nil {
			return nil, err
		}

		for _, module := range modules {
			singleMeasurement, err := getMeasurement(module)
			if err != nil {
				return nil, err
			}
			m.add(singleMeasurement)
		}
	}

	return &m, nil
}

func getMeasurement(c *gabs.Container) (moduleMeasurement, error) {
	singleMeasurement := moduleMeasurement{}
	singleMeasurement.moduleName = c.Search("module_name").Data().(string)
	singleMeasurement.temperature = c.Search("dashboard_data", "Temperature").Data().(float64)
	co2 := c.Search("dashboard_data", "CO2")
	if co2.Data() != nil {
		singleMeasurement.co2 = c.Search("dashboard_data", "CO2").Data().(float64)
	}
	return singleMeasurement, nil
}

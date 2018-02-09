package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/boltdb/bolt"
)

func collectMeasurementsPeriodically(client *http.Client, db *bolt.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(2 * time.Minute)
	for _ = range ticker.C {
		err := collectMeasurements(client, db)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func collectMeasurements(client *http.Client, db *bolt.DB) error {
	m, err := getMeasurements(client)
	if err != nil {
		return err
	}
	fmt.Println("Got measurements from Netatmo")
	return store(m, db)
}

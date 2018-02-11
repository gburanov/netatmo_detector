package main

import (
	"github.com/boltdb/bolt"
)

func store(m *measurements, db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		for _, singleMeasurement := range *m {
			err := storeSingle(singleMeasurement, db)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func storeSingle(m *moduleMeasurement, db *bolt.DB) {

}

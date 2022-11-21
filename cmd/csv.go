package main

import (
	"jatis/pkg/load_data"
	"log"
)

func runCsv() {
	err := load_data.InsertDB("./csv")
	if err != nil {
		log.Fatal(err)
	}
}

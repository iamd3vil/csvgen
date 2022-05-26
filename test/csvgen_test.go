package main

import (
	"log"
	"os"
	"testing"
)

func TestParseCsv(t *testing.T) {
	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatalln(err)
	}

	recs, err := ParsemcxBhavCSV(f)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(recs)
}

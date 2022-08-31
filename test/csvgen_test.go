package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"testing"
)

func TestParseCsv(t *testing.T) {
	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatalln(err)
	}

	rdr := csv.NewReader(f)
	for {
		rec, err := rdr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatalf("error reading csv: %v", err)
		}
		bhav := mcxBhav{}
		if err := bhav.ParseCSV(rec); err != nil {
			t.Fatalf("error parsing csv record: %v", err)
		}
		log.Println(bhav)
	}
}

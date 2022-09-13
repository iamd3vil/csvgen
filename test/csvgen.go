package main

import (
	"fmt"
	"strconv"
)

func (s *testCsv) ParseCSV(rec []string) error {

	s.TestStr = rec[0]

	testint32, err := strconv.ParseInt(rec[1], 0, 32)
	if err != nil {
		return fmt.Errorf("error while parsing TestInt32 at index: %d", 1)
	}
	s.TestInt32 = int32(testint32)

	testint64, err := strconv.ParseInt(rec[2], 0, 64)
	if err != nil {
		return fmt.Errorf("error while parsing TestInt64 at index: %d", 2)
	}
	s.TestInt64 = testint64

	testfloat32, err := strconv.ParseFloat(rec[3], 32)
	if err != nil {
		return fmt.Errorf("error while parsing TestFloat32 at index: %d", 3)
	}
	s.TestFloat32 = float32(testfloat32)

	testfloat64, err := strconv.ParseFloat(rec[4], 64)
	if err != nil {
		return fmt.Errorf("error while parsing TestFloat64 at index: %d", 4)
	}
	s.TestFloat64 = testfloat64

	return nil
}

func (s *testCsvWithTags) ParseCSV(rec []string) error {

	s.TestStr = rec[0]

	testint64, err := strconv.ParseInt(rec[2], 0, 32)
	if err != nil {
		return fmt.Errorf("error while parsing TestInt64 at index: %d", 2)
	}
	s.TestInt64 = int32(testint64)

	testfloat32, err := strconv.ParseFloat(rec[3], 32)
	if err != nil {
		return fmt.Errorf("error while parsing TestFloat32 at index: %d", 3)
	}
	s.TestFloat32 = float32(testfloat32)

	return nil
}

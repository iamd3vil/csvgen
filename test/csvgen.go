package main

import (
	"fmt"
	"strconv"
)

func (s *testCsv) ParseCSV(rec []string) error {

	s.TestStr = rec[0]

	testint32, err := parse2EUjLQTR2xDWilCryXGkH6ttwApInt32(rec[1])
	if err != nil {
		return fmt.Errorf("error while parsing TestInt32 at index: %d", 1)
	}
	s.TestInt32 = testint32

	testint64, err := parse2EUjLQTR2xDWilCryXGkH6ttwApInt64(rec[2])
	if err != nil {
		return fmt.Errorf("error while parsing TestInt64 at index: %d", 2)
	}
	s.TestInt64 = testint64

	testfloat32, err := parse2EUjLQTR2xDWilCryXGkH6ttwApFloat32(rec[3])
	if err != nil {
		return fmt.Errorf("error while parsing TestFloat32 at index: %d", 3)
	}
	s.TestFloat32 = testfloat32

	testfloat64, err := parse2EUjLQTR2xDWilCryXGkH6ttwApFloat64(rec[4])
	if err != nil {
		return fmt.Errorf("error while parsing TestFloat64 at index: %d", 4)
	}
	s.TestFloat64 = testfloat64

	return nil
}

func (s *testCsvWithTags) ParseCSV(rec []string) error {

	s.TestStr = rec[0]

	testint64, err := parse2EUjLQTR2xDWilCryXGkH6ttwApInt32(rec[2])
	if err != nil {
		return fmt.Errorf("error while parsing TestInt64 at index: %d", 2)
	}
	s.TestInt64 = testint64

	testfloat32, err := parse2EUjLQTR2xDWilCryXGkH6ttwApFloat32(rec[3])
	if err != nil {
		return fmt.Errorf("error while parsing TestFloat32 at index: %d", 3)
	}
	s.TestFloat32 = testfloat32

	return nil
}

func parse2EUjLQTR2xDWilCryXGkH6ttwApInt32(s string) (int32, error) {
	if s == "" {
		return 0, nil
	}
	i, err := strconv.ParseInt(s, 0, 32)
	return int32(i), err
}

func parse2EUjLQTR2xDWilCryXGkH6ttwApInt64(s string) (int64, error) {
	if s == "" {
		return 0, nil
	}
	return strconv.ParseInt(s, 0, 64)
}

func parse2EUjLQTR2xDWilCryXGkH6ttwApFloat32(s string) (float32, error) {
	if s == "" {
		return 0, nil
	}
	i, err := strconv.ParseFloat(s, 32)
	return float32(i), err
}

func parse2EUjLQTR2xDWilCryXGkH6ttwApFloat64(s string) (float64, error) {
	if s == "" {
		return 0, nil
	}
	return strconv.ParseFloat(s, 64)
}

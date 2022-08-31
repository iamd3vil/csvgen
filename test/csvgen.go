package main

import (
	"fmt"
	"strconv"
)

func (s *testCsv) ParseCSV(rec []string) error {
	index := 0

	s.TestStr = rec[index]

	index += 1

	testint32, err := parse2E8Bp5Vfel8J7JhvR6dTVcgc5UxInt32(rec[index])
	if err != nil {
		return fmt.Errorf("error while parsing TestInt32 at index: %d", index)
	}
	s.TestInt32 = testint32

	index += 1

	testint64, err := parse2E8Bp5Vfel8J7JhvR6dTVcgc5UxInt64(rec[index])
	if err != nil {
		return fmt.Errorf("error while parsing TestInt64 at index: %d", index)
	}
	s.TestInt64 = testint64

	index += 1

	testfloat32, err := parse2E8Bp5Vfel8J7JhvR6dTVcgc5UxFloat32(rec[index])
	if err != nil {
		return fmt.Errorf("error while parsing TestFloat32 at index: %d", index)
	}
	s.TestFloat32 = testfloat32

	index += 1

	testfloat64, err := parse2E8Bp5Vfel8J7JhvR6dTVcgc5UxFloat64(rec[index])
	if err != nil {
		return fmt.Errorf("error while parsing TestFloat64 at index: %d", index)
	}
	s.TestFloat64 = testfloat64

	return nil
}

func parse2E8Bp5Vfel8J7JhvR6dTVcgc5UxInt32(s string) (int32, error) {
	if s == "" {
		return 0, nil
	}
	i, err := strconv.ParseInt(s, 0, 32)
	return int32(i), err
}

func parse2E8Bp5Vfel8J7JhvR6dTVcgc5UxInt64(s string) (int64, error) {
	if s == "" {
		return 0, nil
	}
	return strconv.ParseInt(s, 0, 64)
}

func parse2E8Bp5Vfel8J7JhvR6dTVcgc5UxFloat32(s string) (float32, error) {
	if s == "" {
		return 0, nil
	}
	i, err := strconv.ParseFloat(s, 32)
	return float32(i), err
}

func parse2E8Bp5Vfel8J7JhvR6dTVcgc5UxFloat64(s string) (float64, error) {
	if s == "" {
		return 0, nil
	}
	return strconv.ParseFloat(s, 64)
}

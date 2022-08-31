package main

import (
	"fmt"
	"strconv"
)

func (s *mcxBhav) ParseCSV(rec []string) error {
	index := 0

	s.Date = rec[index]

	index += 1

	sessionid, err := parse2E81XK93q8o8kdqPuPf48b13RftInt32(rec[index])
	if err != nil {
		return fmt.Errorf("error while parsing Sessionid at index: %d", index)
	}
	s.Sessionid = sessionid

	index += 1

	s.Markettype = rec[index]

	index += 1

	instrumentid, err := parse2E81XK93q8o8kdqPuPf48b13RftInt32(rec[index])
	if err != nil {
		return fmt.Errorf("error while parsing Instrumentid at index: %d", index)
	}
	s.Instrumentid = instrumentid

	index += 1

	s.Instrumentname = rec[index]

	index += 1

	s.Symbol = rec[index]

	index += 1

	s.Expirydate = rec[index]

	index += 1

	s.Reserved1 = rec[index]

	index += 1

	index += 1

	s.Opttype = rec[index]

	index += 1

	index += 1

	index += 1

	index += 1

	index += 1

	index += 1

	volume, err := parse2E81XK93q8o8kdqPuPf48b13RftInt64(rec[index])
	if err != nil {
		return fmt.Errorf("error while parsing Volume at index: %d", index)
	}
	s.Volume = volume

	index += 1

	index += 1

	index += 1

	index += 1

	s.Unit = rec[index]

	index += 1

	index += 1

	index += 1

	index += 1

	index += 1

	s.Reserved2 = rec[index]

	index += 1

	s.Currencycode = rec[index]

	return nil
}

func parse2E81XK93q8o8kdqPuPf48b13RftInt32(s string) (int32, error) {
	if s == "" {
		return 0, nil
	}
	i, err := strconv.ParseInt(s, 0, 32)
	return int32(i), err
}

func parse2E81XK93q8o8kdqPuPf48b13RftInt64(s string) (int64, error) {
	if s == "" {
		return 0, nil
	}
	return strconv.ParseInt(s, 0, 64)
}

func parse2E81XK93q8o8kdqPuPf48b13RftFloat32(s string) (float32, error) {
	if s == "" {
		return 0, nil
	}
	i, err := strconv.ParseFloat(s, 32)
	return float32(i), err
}

func parse2E81XK93q8o8kdqPuPf48b13RftFloat64(s string) (float64, error) {
	if s == "" {
		return 0, nil
	}
	return strconv.ParseFloat(s, 64)
}

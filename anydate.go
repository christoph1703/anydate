package anydate

import (
	"errors"
	"strings"
	"time"
)

var (
	dateFormats = []string{
		"2.1.2006",
		"2006-2-1",
		"Jan 02, 2006",
	}
)

// Parse scans a string and returns a list of all dates contained in the string.
func Parse(s string) []time.Time {
	dates := []time.Time{}

	elements := strings.Fields(s)

	for i, e := range elements {
		yearStart := strings.Index(e, "20")
		if yearStart < 0 {
			// not a year
			continue
		}
		if len(e)-yearStart < 4 {
			// no space for year
			continue
		}

		switch {
		case len(e) == 4:
			e = strings.Join(elements[i-2:i+1], " ")
		}

		d, err := parseDate(e)
		if err != nil {
			continue
		}
		dates = append(dates, d)
	}

	return dates
}

func parseDate(s string) (time.Time, error) {
	for _, f := range dateFormats {
		t, err := time.Parse(f, s)
		if err == nil {
			return t, nil
		}
	}

	return time.Now(), errors.New("could not read date")
}

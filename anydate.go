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
		"Jan 2, 2006",
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

		for j := 0; j < i && j < 3; j++ {
			e = strings.Join(elements[i-j:i+1], " ")
			d, err := parseDate(e)
			if err != nil {
				continue
			}
			dates = append(dates, d)
			break
		}

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

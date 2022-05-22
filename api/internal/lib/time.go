package lib

import "time"

func ParseDateTime(s string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Time{}, err
	}
	return t.UTC(), nil
}

func FormatDateTime(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}

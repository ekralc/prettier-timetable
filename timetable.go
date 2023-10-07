package main

import (
	"net/http"

	ical "github.com/arran4/golang-ical"
)

func fetchCalendarFromURL(url string) (*ical.Calendar, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ical.ParseCalendar(resp.Body)
}

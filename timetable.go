package main

import (
	"fmt"
	"net/http"
	"strings"

	ical "github.com/arran4/golang-ical"
	ics "github.com/arran4/golang-ical"
)

func fetchCalendarFromURL(url string) (*ical.Calendar, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ical.ParseCalendar(resp.Body)
}

func getModuleName(event *ical.VEvent) (string, error) {
	description := event.GetProperty(ics.ComponentPropertyDescription).Value

	// fmt.Println(description)

	for i, line := range strings.Split(description, "\r\n") {
		fmt.Println(i, line)
	}
	// lines := strings.Split(description, "\n")

	// for _, line := range lines {
	// 	fmt.Println(line)
	// 	if name, found := strings.CutPrefix(line, "Module Name: "); found {
	// 		return name, nil
	// 	}
	// }

	return "", nil
}

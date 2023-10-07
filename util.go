package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	ics "github.com/arran4/golang-ical"
)

func TransformCalendar(calendar *ics.Calendar) {
	for _, event := range calendar.Events() {
		description := GetCleanEventDescription(event)

		name := GetModuleNameFromString(description)
		activity := GetActivityTypeFromString(description)

		if name == "" {
			name = GetModuleCodeFromString(description)
		}

		cleanTitle := fmt.Sprintf("(%v) %v", activity, name)
		event.SetSummary(cleanTitle)
	}
}

func FetchCalendarFromURL(url string) (*ics.Calendar, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ics.ParseCalendar(resp.Body)
}

func GetCleanEventDescription(event *ics.VEvent) string {
	description := event.GetProperty(ics.ComponentPropertyDescription).Value
	// Replace newline characters with actual newlines (golang-ical is weird)
	description = strings.ReplaceAll(description, "\\n", "\n")

	return description
}

func GetFieldFromEventDescription(description string, fieldName string) string {
	r := regexp.MustCompile(fmt.Sprintf("%v: (.+)", fieldName))
	matches := r.FindStringSubmatch(description)

	if len(matches) < 2 {
		return ""
	}

	return matches[1]
}

func GetModuleNameFromString(s string) string {
	return GetFieldFromEventDescription(s, "Module Name")
}

func GetModuleCodeFromString(s string) string {
	return GetFieldFromEventDescription(s, "Module code")
}

func GetActivityTypeFromString(s string) string {
	return GetFieldFromEventDescription(s, "Activity/Session Type")
}

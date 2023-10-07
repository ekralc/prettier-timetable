package main

import (
	"fmt"
	"net/http"
	"net/url"

	ics "github.com/arran4/golang-ical"
	"github.com/gin-gonic/gin"
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

func main() {
	r := gin.Default()

	r.UseRawPath = true
	r.UnescapePathValues = false

	r.GET("/:timetable", func(c *gin.Context) {
		timetable_url, _ := url.QueryUnescape(c.Params.ByName("timetable"))

		_, err := url.Parse(timetable_url)
		if err != nil {
			c.String(http.StatusBadRequest, "invalid timetable url")
			return
		}

		calendar, err := fetchCalendarFromURL(timetable_url)
		if err != nil {
			c.String(http.StatusInternalServerError, "couldn't parse timetable")
			return
		}

		TransformCalendar(calendar)
		c.Writer.Header().Set("Content-Type", "text/calendar")
		c.Writer.Header().Set("Content-Disposition", "attachment; filename=\"calendar.ics\"")
		calendar.SerializeTo(c.Writer)
		c.Status(200)
	})

	r.Run()
}

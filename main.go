package main

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func RequestHandler(c *gin.Context) {
	// We expect a URL encoded timetable URL
	timetable_url, err := url.QueryUnescape(c.Params.ByName("timetable"))
	if err != nil || !UrlAllowed(timetable_url) {
		c.String(http.StatusBadRequest, "Invalid URL. Please request with a URL-encoded MyTimetable link.")
	}

	exclude_list := []string{}
	exclude_param := c.Query("exclude")
	if exclude_param != "" {
		exclude_param_list := strings.Split(exclude_param, ",")
		exclude_list = append(exclude_list, exclude_param_list...)
	}

	raw_calendar, err := FetchCalendarFromURL(timetable_url)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing timetable")
		return
	}

	transformed_calendar := TransformCalendar(raw_calendar, exclude_list)

	c.Writer.Header().Set("Content-Type", "text/calendar")
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=\"calendar.ics\"")
	transformed_calendar.SerializeTo(c.Writer)

	c.Status(200)
}

func main() {
	r := gin.Default()

	// These settings are necessary to process URL encoded parameters
	r.UseRawPath = true
	r.UnescapePathValues = false

	r.GET("/:timetable", RequestHandler)
	r.Run()
}

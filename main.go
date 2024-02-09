package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func RequestHandler(c *gin.Context) {
	if c.Query("eu") == "" || c.Query("h") == "" {
		c.String(http.StatusBadRequest, "Invalid MyTimetable parameters")
		return
	}

	timetableUrl := fmt.Sprintf("https://mytimetable.leeds.ac.uk%v", c.Request.RequestURI)
	calendar, err := FetchCalendarFromURL(timetableUrl)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing timetable")
		return
	}

	TransformCalendar(calendar)

	c.Writer.Header().Set("Content-Type", "text/calendar")
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=\"calendar.ics\"")
	calendar.SerializeTo(c.Writer)

	c.Status(200)
}

func BackwardsCompatibleRequestHandler(c *gin.Context) {
	// We expect a URL encoded timetable URL
	timetableUrl, err := url.QueryUnescape(c.Params.ByName("timetable"))
	if err != nil || !UrlAllowed(timetableUrl) {
		c.String(http.StatusBadRequest, "Invalid URL. Please request with a URL-encoded MyTimetable link.")
		return
	}

	calendar, err := FetchCalendarFromURL(timetableUrl)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing timetable")
		return
	}

	TransformCalendar(calendar)

	c.Writer.Header().Set("Content-Type", "text/calendar")
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=\"calendar.ics\"")
	calendar.SerializeTo(c.Writer)

	c.Status(200)
}

func main() {
	r := gin.Default()

	// These settings are necessary to process URL encoded parameters
	r.UseRawPath = true
	r.UnescapePathValues = false

	r.GET("/ical", RequestHandler)
	// Deprecated
	r.GET("/:timetable", BackwardsCompatibleRequestHandler)

	r.Run()
}

package main

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func RequestHandler(c *gin.Context) {
	// We expect a URL encoded timetable URL
	timetable_url, _ := url.QueryUnescape(c.Params.ByName("timetable"))
	if timetable_url == "" {
		c.String(http.StatusBadRequest, "no timetable url")
		return
	}

	calendar, err := FetchCalendarFromURL(timetable_url)
	if err != nil {
		c.String(http.StatusInternalServerError, "couldn't parse timetable")
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

	r.GET("/:timetable", RequestHandler)
	r.Run()
}

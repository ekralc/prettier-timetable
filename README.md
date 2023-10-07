# Prettier Timetable

Simplifies my university schedule with readable event titles.

The service has a single endpoint that takes the original timetable link as a URL encoded parameter. The service fetches the calendar from the original URL, rewrites the titles into a human-readable format, and returns the modified timetable to the user.

For example, `COMP381101 - COMP381101/LEC 1/01` becomes `(Lecture) Computer Graphics`


# Prettier Timetable
Rewrites the MyTimetable iCal feed to be more human-readable. For example, the title `COMP381101 - COMP381101/LEC 1/01` becomes `(Lecture) Computer Graphics`.

The service acts as proxy for MyTimetable by fetching the calendar from the provided URL, rewriting event titles and returning it as usual. 

To use the service, take your MyTimetable iCal URL and replace the domain with `timetable.jclarke.tech`, e.g. `https://mytimetable.leeds.ac.uk/eu=foo&h=bar` becomes `https://timetable.jclarke.tech/eu=foo&h=bar`


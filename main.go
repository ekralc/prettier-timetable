package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var TIMETABLE_URL = os.Getenv("TIMETABLE_URL")

func main() {
	fmt.Println("Hello, World!")

	calendar, err := fetchCalendarFromURL(TIMETABLE_URL)
	if err != nil {
		fmt.Printf("error loading calendar: %v", err)
	}

	fmt.Println(len(calendar.Events()))

	for _, event := range calendar.Events() {
		name, _ := getModuleName(event)
		fmt.Printf("%v\n", name)
	}
}

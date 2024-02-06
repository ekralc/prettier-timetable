package main

import "testing"

func TestGetModuleNameBasic(t *testing.T) {
	ans := GetModuleNameFromString("Module Name: Distributed Systems")
	if ans != "Distributed Systems" {
		t.Errorf("got %v, want %v", ans, "Distributed Systems")
	}
}

func TestGetModuleNameWithOtherText(t *testing.T) {
	ans := GetModuleNameFromString("Module activity: Lab\n\nModule Name: Distributed Systems\n\nOther details: stuff")
	if ans != "Distributed Systems" {
		t.Errorf("got %v, want %v", ans, "Distributed Systems")
	}
}

func TestGetActivityTypeBasic(t *testing.T) {
	input := "Activity/Session Type: Lecture"
	ans := GetActivityTypeFromString(input)
	if ans != "Lecture" {
		t.Errorf("got %v, want %v", ans, "Lecture")
	}
}

func TestUrlAllowedAccepts(t *testing.T) {
	url := "https://mytimetable.leeds.ac.uk/foobar"
	allowed := UrlAllowed(url)
	if !allowed {
		t.Errorf("got %v, want %v", allowed, true)
	}
}

func TestUrlAllowedRejects(t *testing.T) {
	urls := []string{"https://google.com", "www.example.com", "http://malicioussite.com", "ftp://somewhere", "random_protocol://hello"}
	for _, url := range urls {
		if UrlAllowed(url) {
			t.Errorf("%v should not be allowed by the filter", url)
		}
	}
}

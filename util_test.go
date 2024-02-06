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

// Package solutions provides ...
package solutions

import (
	"testing"
)

func TestBetterLinearSearch(t *testing.T) {
	var arr = []string{
		"NZ4b",
		"O0tL",
		"0K34",
		"kNjv",
		"MSyb",
		"XQ7f",
		"ALVk",
		"tURG",
		"tDAg",
		"Fctk",
	}
	res, i := BetterLinearSearch(arr, len(arr), "tURG")
	if res[0] == "not found" {
		t.Error("no value found")
	}
	if i > 7 {
		t.Error("for loop did not break")
	}
}

// Package solutions provides ...
package solutions

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
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
	var res []string
	res = LinearSearch(arr, len(arr), "")
	if res[0] == "not found" {
		t.Error("no value found")
	}
}

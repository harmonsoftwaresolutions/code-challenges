package solutions

import (
	"fmt"
	"testing"
)

func TestSentinelLinearSearch(t *testing.T) {
	var a = []string{
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
	n := len(a) - 1
	s := SentinelLinearSearch(a, n, a[7])
	fmt.Printf("s is %v", s)
	if s == "not found" {
		t.Error("no value found")
		return
	}

	// if s != "7" {
	// t.Error("incorrect index %s", s)
	// }

}

package practice

import "testing"

func TestFirstReverse(t *testing.T) {
	var s string
	s = FirstReverse("Hello World and Coders")

	if s != "sredoC dna dlroW olleH" {
		t.Error("Expected string to be reversed, got ", s)
	}
}

package stringutils

import "testing"

func TestReverseString(t *testing.T) {
	var s string
	s = ReverseString("This is a test")

	if s != "tset a si sihT" {
		t.Error("Expected string to be reversed, got ", s)
	}
}

func TestLongestWord(t *testing.T) {
	var s string
	s = LongestWord("This is a sentence with an enormously long word")

	if s != "enormously" {
		t.Error("Expected longest word to be enormously, got ", s)
	}
}

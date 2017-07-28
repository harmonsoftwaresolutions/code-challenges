package stringutils

import "testing"

func TestLongestWord(t *testing.T) {
	var s string
	s = LongestWord("This is a sentence with an enormously long word")

	if s != "enormously" {
		t.Error("Expected longest word to be enormously, got ", s)
	}
}

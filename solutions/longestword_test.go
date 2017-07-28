package solutions

import "testing"

func TestLongestWord(t *testing.T) {
	var s1 string
	s1 = LongestWord("a beautiful sentence^&!")
	ls1 := "beautiful"
	if s1 != ls1 {
		t.Error("Expected longest word to be beautiful, got ", s1)
	}

	s2 := LongestWord("letter after letter")
	ls2 := "letter"
	if s2 != ls2 {
		t.Error("Expected longest word to be letter, got ", s2)
	}

	s3 := LongestWord("hello world")
	ls3 := "hello"
	if s3 != ls3 {
		t.Error("Expected longest word to be hello, got ", s3)
	}

	s4 := LongestWord("123456789 98765432")
	ls4 := "123456789"
	if s4 != ls4 {
		t.Error("Expected longest word to be hello, got ", s3)
	}
}

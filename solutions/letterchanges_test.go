package solutions

import "testing"

func TestLetterChanges(t *testing.T) {
	ls1 := "hello*3"
	exp1 := "Ifmmp*3"
	var s string
	s = LetterChanges(ls1)

	if s != exp1 {
		t.Error("expected ", exp1, "got ", s)
	}
}

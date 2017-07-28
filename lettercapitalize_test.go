package stringutils

import "testing"

func TestLetterCapitalize(t *testing.T) {
	var s string
	s = LetterCapitalize("this is all lower case letters")

	if s != "This Is All Lower Case Letters" {
		t.Error("expected every letter to be capitalized, got ", s)
	}
}

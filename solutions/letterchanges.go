// Package solutions provides
package solutions

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Have the function LetterChanges(str) take the str parameter being passed
// and modify it using the following algorithm. Replace every letter in the
// string with the letter following it in the alphabet
// (ie. c becomes d, z becomes a). Then capitalize every vowel in this new
// string (a, e, i, o, u) and finally return this modified string.

func LetterChanges(s string) string {
	// break up to array of unicode utf-8
	vowels := "aeiou"
	arr := strings.Split(s, "")

	// make new array for transposing
	narr := make([]string, len(arr))
	for i := range arr {
		// check if letter and only tranpose +1 for letters
		r, _ := utf8.DecodeRune([]byte(arr[i]))
		if l := unicode.IsLetter(r); l == true {
			r++
		}
		// check for vowel and capitalize
		if v := strings.Contains(vowels, string(r)); v == true {
			r = unicode.ToUpper(r)
		}
		narr[i] = string(r)

	}

	var str string
	str = strings.Join(narr, "")
	return str
}

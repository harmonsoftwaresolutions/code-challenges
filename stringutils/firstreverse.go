// Package stringutils challenges
package stringutils

import (
	"fmt"
)

// Have the function FirstReverse(str) take the str parameter being passed and
// return the string in reversed order. For example: if the input string is
// "Hello World and Coders" then your program should return the string
// "sredoC dna dlroW olleH".

// FirstReverse reverse a string using runes
func FirstReverse(s string) string {
	// convert to array of runes
	r := []rune(s)
	fmt.Printf("convert to array of runes:\n%v\n", r)
	// i to iterate, n to track last string
	// flips letters by moving from outside to meeting in the middle
	for i, n := 0, len(s)-1; i < len(s)/2; i, n = i+1, n-1 {
		fmt.Printf("before: %v\t", string(r))
		fmt.Printf("i: %v\t", i)
		fmt.Printf("n: %v\t", n)
		r[i], r[n] = r[n], r[i]
		fmt.Printf("after: %v\n", string(r))
	}
	// convert back to utf8
	str := string(r)
	return str
}

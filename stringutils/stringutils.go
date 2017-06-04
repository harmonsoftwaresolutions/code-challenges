// Package stringutils challenges
package stringutils

import (
	"fmt"
	"strings"
	"unicode"
)

// ReverseString reverse a string using runes
func ReverseString(s string) string {
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

// LongestWord in a string
func LongestWord(s string) string {
	// track word index and node length
	var idx int
	bnode := 1
	fmt.Printf("string: %s\n", s)
	arr := strings.Fields(s)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("i: %d\t", i)
		w := arr[i]
		fmt.Printf("word: %v\t", w)
		fmt.Printf("length: %v\t", len(w))
		if len(w) > bnode {
			bnode = len(w)
			idx = i
		}
		fmt.Printf("idx: %v\n", idx)
	}

	return arr[idx]
}

// LetterCapitalize capitalize first letter
func LetterCapitalize(s string) string {
	// split by space
	arr := strings.Split(s, " ")
	fmt.Println("array len ", len(arr))
	// capitalize first letters
	for i := range arr {
		fmt.Println("i ", i)
		fmt.Println("arr item", arr[i])
		// fmt.Println("index check ", len(arr[i]))
		// fmt.Println("array check ", arr[i])
		r := []rune(arr[i])
		fmt.Println("rune ", r)
		fmt.Println("rune ", r[0])
		r[0] = unicode.ToUpper(r[0])
		arr[i] = string(r)
	}
	fmt.Println("arr ", arr)
	// add spaces back in to new string
	newStr := ""
	for i := 0; i < len(arr); i++ {
		if i != 0 {
			newStr = newStr + " " + arr[i]
		} else {
			newStr = arr[i]
		}
		fmt.Println("newStr ", newStr)
	}
	return newStr
}

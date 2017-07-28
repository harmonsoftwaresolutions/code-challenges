// Package stringutils challenges
package stringutils

import (
	"fmt"
	"strings"
	"unicode"
)

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

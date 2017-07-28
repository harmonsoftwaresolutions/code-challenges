// Package stringutils challenges
package stringutils

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// Have the function LongestWord(sen) take the sen parameter being passed and
// return the largest word in the string. If there are two or more words that
// are the same length, return the first word from the string with that length.
// Ignore punctuation and assume sen will not be empty.

// LongestWord in a string
func LongestWord(sen string) string {
	// track word index and node length
	var idx int
	bnode := 1
	// regex to remove non alpha chars
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]")
	if err != nil {
		log.Fatal(err)
	}
	s := reg.ReplaceAllString(sen, "")
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

package main

import (
	"fmt"
	_ "github.com/harmonsoftwaresolutions/coderbyte-challenges/factorials"
	"github.com/harmonsoftwaresolutions/coderbyte-challenges/stringutils"
)

func main() {
	// firstFact := factorials.FirstFactorial(8)
	// fmt.Println(firstFact)
	// revStr := stringutils.ReverseString("this is a test")
	// fmt.Println(revStr)
	longestWord := stringutils.LongestWord("this is a testing center")
	fmt.Println(longestWord, len(longestWord))
}

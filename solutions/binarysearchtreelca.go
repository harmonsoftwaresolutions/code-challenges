// Package solutions provides ...
package solutions

import (
	"fmt"
	"log"
	"regexp"
)

func BinarySearchTreeLCA(strArr string) int {
	ans := 0

	// remove strings
	reg, err := regexp.Compile("[\"]")
	if err != nil {
		log.Fatal(err)
	}
	s := reg.ReplaceAllString(strArr, "")
	fmt.Printf("s is %v\n", s)

	// turn first item in argument array into array of integers
	if err != nil {
		log.Fatal(err)
	}
	// s2 := regexp.MustCompile("[\\[\\]]").ReplaceAllString(s, "")
	s2 := regexp.MustCompile(`[^\\[](.*)[^\\]]`).FindString(s)
	// s2 := regexp.MustCompile("[\\[\\]]").Split(s2, 6)
	fmt.Printf("s2 is %v\n", s2)
	return ans
}

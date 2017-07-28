// Package stringutils challenges
package stringutils

import (
	"fmt"
	"strings"
	"unicode"
)

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

// Package factorials challenges
package factorials

import (
	"fmt"
)

// FirstFactorial find the first factorial
func FirstFactorial(num int) int {
	fmt.Printf("Num input: %d\n", num)
	factorial := 1
	for i := 1; i <= num; i++ {
		fmt.Printf("i: %v\t", i)
		fmt.Printf("factorial: %v\t", factorial)
		factorial = factorial * i
		fmt.Printf("factorial: %v\n", factorial)
	}
	return factorial
}

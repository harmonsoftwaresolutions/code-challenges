// Package factorials challenges
package factorials

import ()

// FirstFactorial find the first factorial
func FirstFactorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * FirstFactorial(num-1)
}

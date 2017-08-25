// Package solutions provides ...
package solutions

import (
	"strconv"
)

func SentinelLinearSearch(a []string, n int, x string) string {
	ans := "not found"
	last := a[n]
	a[n] = x

	i := 0
	for i < n {
		if a[i] == x {
			break
		}
		i++
	}

	a[n] = last
	if i < n || a[n] == x {
		ans = strconv.Itoa(int(i))
	}

	return ans
}

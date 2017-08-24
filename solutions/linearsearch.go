// Package solutions provides ...
package solutions

import (
	"fmt"
)

func LinearSearch(a []string, n int, x string) []string {
	ans := []string{"not found"}
	for i := 0; i < n; i += 1 {
		if s := a[i]; s == x {
			ans[0] = s
		}
	}
	fmt.Println(ans)
	return ans
}

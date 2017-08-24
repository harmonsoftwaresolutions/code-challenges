// Package solutions provides ...
package solutions

import (
	"fmt"
)

// Track i and return to confirm the break
func BetterLinearSearch(a []string, n int, x string) ([]string, int) {
	ans := []string{"not found"}
	i := 0
	for i = 0; i < n; i += 1 {
		if s := a[i]; s == x {
			ans[0] = s
			break
		}
	}
	fmt.Println(ans)
	return ans, i
}

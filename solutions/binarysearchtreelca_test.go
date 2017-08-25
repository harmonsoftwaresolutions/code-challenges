// Package solutions provides ...
package solutions

import (
	"testing"
)

func TestBinarySearchTreeLCA(t *testing.T) {
	// Input:"[10, 5, 1, 7, 40, 50]", "5", "10"

	input := "\"[10,5,1,7,40,50]\", \"5\", \"10\""
	res := BinarySearchTreeLCA(input)
	if res == 0 {
		t.Error("should not be 0")
	}
	return
}

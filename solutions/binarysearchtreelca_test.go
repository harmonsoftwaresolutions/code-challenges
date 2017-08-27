// Package solutions provides ...
package solutions

import (
	"testing"
)

func TestBinarySearchTreeLCA(t *testing.T) {

	t.Run("A=5", func(t *testing.T) {
		input := "\"[10,5,1,7,40,50]\", \"1\", \"7\""
		e := "5"
		res := BinarySearchTreeLCA(input)
		if res != e {
			t.Error("Result should be", e)
		}
	})

	t.Run("A=10", func(t *testing.T) {
		input := "\"[10,5,1,7,40,50]\", \"5\", \"10\""
		e := "10"
		res := BinarySearchTreeLCA(input)
		if res != e {
			t.Error("Result should be", e)
		}
	})

	t.Run("A=12", func(t *testing.T) {
		input := "\"[3,2,1,12,4,5,13]\", \"5\", \"13\""
		e := "12"
		res := BinarySearchTreeLCA(input)
		if res != e {
			t.Error("Result should be", e)
		}
	})
}

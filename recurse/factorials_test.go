package recurse

import "testing"

func TestFirstFactorial(t *testing.T) {
	var num int
	num = FirstFactorial(8)

	if num != 40320 {
		t.Error("Expected 40320, got ", num)
	}
}

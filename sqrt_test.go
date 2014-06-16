package newmath

import "testing"

func checkSqrt(in, out float64, t *testing.T) {
	if x := Sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	}
}

func TestSqrt(t *testing.T) {
	const in, out = 4, 2
	checkSqrt(4, 2, t)
	checkSqrt(4, 3, t)
}

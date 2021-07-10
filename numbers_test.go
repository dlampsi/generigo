package generigo

import (
	"testing"
)

func Test_IsOdd(t *testing.T) {
	f := func(num int, want bool) {
		t.Helper()
		odd := IsOdd(num)
		if odd != want {
			t.Fatalf("result mismatch; get: %v , want: %v", odd, want)
		}
	}

	f(4, false)
	f(7, true)
	f(0, false)
}

func Test_AverageFloat64(t *testing.T) {
	f := func(floats []float64, want float64) {
		t.Helper()
		av := AverageFloat64(floats)
		if av != want {
			t.Fatalf("unexpected average; want: %v, get: %v", want, av)
		}
	}

	f([]float64{1.1, 1.1}, 1.1)
	f([]float64{2.0, 1.0}, 1.5)
}

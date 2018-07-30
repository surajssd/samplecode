package math

import "testing"

func TestAdd(t *testing.T) {
	tcs := []struct {
		a      float64
		b      float64
		result float64
	}{
		{1, 2, 3},
		{5, -7, -2},
	}

	for _, tc := range tcs {
		if Add(tc.a, tc.b) != tc.result {
			t.Errorf("the output does not match for %f+%f=%f", tc.a, tc.b, tc.result)
		}
	}
}

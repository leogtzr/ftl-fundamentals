package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	const want = 6.0
	got := calculator.Multiply(2, 3)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b        float64
		want        float64
		errExpected bool
	}

	tests := []testCase{
		{
			a:           6,
			b:           0,
			want:        0,
			errExpected: true,
		},
		{
			a:           3,
			b:           2,
			want:        1.5,
			errExpected: false,
		},
	}

	for _, tc := range tests {
		got, err := calculator.Divide(tc.a, tc.b)
		if tc.errExpected {
			if err == nil {
				t.Error("expecting an error")
			}
		}
		if got != tc.want {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}

}

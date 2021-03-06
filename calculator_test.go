package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
	"time"
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
		if errReceived := err != nil; tc.errExpected != errReceived {
			t.Fatalf("Divide(%f, %f): unexpected error status: %v", tc.a, tc.b, errReceived)
		}
		if !tc.errExpected && got != tc.want {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}

}

func TestAddRandom(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		a := rand.Float64()
		b := rand.Float64()
		sum := a + b
		if got := calculator.Add(a, b); got != sum {
			t.Errorf("want=%f, got=%f", sum, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	type testCase struct {
		x           float64
		want        float64
		errExpected bool
	}

	tests := []testCase{
		{
			x:           -1,
			want:        0,
			errExpected: true,
		},
		{
			x:           16,
			want:        4,
			errExpected: false,
		},
	}

	for _, tc := range tests {
		got, err := calculator.Sqrt(tc.x)
		if errReceived := err != nil; tc.errExpected != errReceived {
			t.Fatalf("Sqrt(%f): unexpected error status: %v", tc.x, errReceived)
		}
		if !tc.errExpected && got != tc.want {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

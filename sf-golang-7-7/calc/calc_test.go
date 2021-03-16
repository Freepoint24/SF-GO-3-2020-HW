package calc

import (
	"math"
	"testing"
)

func TestCalculate(t *testing.T) {
	var result, wanted float64
	var err error
	c := NewCalculator()

	result, _ = c.Calculate(2, 2, "+")
	wanted = 4
	if result != wanted {
		t.Fatalf(`Calculate(2, 2, "+") => %f; want => %f`, result, wanted)
	}

	result, _ = c.Calculate(4, 1, "-")
	wanted = 3
	if result != wanted {
		t.Fatalf(`Calculate(4, 2, "-") => %f; want => %f`, result, wanted)
	}

	result, _ = c.Calculate(6, 6, "*")
	wanted = 36
	if result != wanted {
		t.Fatalf(`Calculate(6, 6, "*") => %f; want => %f`, result, wanted)
	}

	result, _ = c.Calculate(24, 16, "/")
	wanted = 1.5
	if result != wanted {
		t.Fatalf(`Calculate(24, 16, "/") => %f; want => %f`, result, wanted)
	}

	result, err = c.Calculate(1, 0, "/")
	if !math.IsNaN(result) || err != ErrDivZero {
		t.Fatalf(`Calculate(1, 0, "/") => %f, %s; want => %f, %s`, result, err, math.NaN(), ErrDivZero)
	}

	result, err = c.Calculate(2, 8, "^")
	wanted = math.NaN()
	if !math.IsNaN(result) || err != ErrUnknownOperator {
		t.Fatalf(`Calculate(2, 8, "^") => %f, %s; want => %f, %s`, result, err, math.NaN(), ErrUnknownOperator)
	}
}

package squareEq

import "testing"

// TestSolveRealZeroA - если a = 0
func TestSolveRealZeroA(t *testing.T) {
	_, _, err := SolveReal(0, 1, 2)
	if err != ErrZeroA {
		t.Fatalf("SolveReal(0, 1, 2) => %s, want => %s", err, ErrZeroA)
	}
}

// TestSolveRealNoRoots - если дискриминант отрицательный и нет вещественных корней
func TestSolveRealNoRoots(t *testing.T) {
	_, _, err := SolveReal(1, 2, 3)
	if err != ErrNoRealRoots {
		t.Fatalf("SolveReal(1, 2, 3) => %s, want => %s", err, ErrNoRealRoots)
	}
}

// TestSolveRealZeroD - если дискриминант равен нулю и вещественный корень один
func TestSolveRealZeroD(t *testing.T) {
	x1, x2, err := SolveReal(2, 4, 2)
	if err != nil {
		t.Fatalf("SolveReal(2, 4, 2) => %s, want => nil", err)
	}
	if (x1 != -1) && (x2 != x1) {
		t.Fatalf("SolveReal(2, 4, 2) => %f, %f, want => -1, -1", x1, x2)
	}
}

// TestSolveRealRoots - если дискриминант положительный и есть два вещественных корня
func TestSolveRealRoots(t *testing.T) {
	x1, x2, err := SolveReal(2, 3, 1)
	if err != nil {
		t.Fatalf("SolveReal(2, 3, 1) => %s, want => nil", err)
	}
	if (x1 != -1) && (x2 != -0.5) {
		t.Fatalf("SolveReal(2, 3, 1) => %f, %f, want => -1, -0.5", x1, x2)
	}
}

// TestSolveComplexZeroA - если a = 0
func TestSolveComplexZeroA(t *testing.T) {
	_, _, err := SolveComplex(0, 1, 2)
	if err != ErrZeroA {
		t.Fatalf("SolveComplex(0, 1, 2) => %s, want => %s", err, ErrZeroA)
	}
}

// TestSolveComplexRoots - комплексные корни
func TestSolveComplexRoots(t *testing.T) {
	x1, x2, err := SolveComplex(1+0i, -5+6i, -1-9i)
	if err != nil {
		t.Fatalf("SolveComplex(1+0i, -5+6i, -1-9i) => %s, want => nil", err)
	}
	if (x1 != 1-1i) && (x2 != 4-5i) {
		t.Fatalf("SolveComplex(1+0i, -5+6i, -1-9i) => %f, %f, want => 1-1i, 4-5i", x1, x2)
	}
}

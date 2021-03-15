package main

import (
	"fmt"
	"sf-golang-7-2/fmtnum"
	"sf-golang-7-2/square-eq"
)

func main() {
	var (
		a float64 = 2
		b float64 = 3
		c float64 = 1
	)

	fmt.Printf("Уравнение %sx²+%sx+%s=0", fmtnum.F64(a), fmtnum.F64(b), fmtnum.F64(c))
	x1, x2, err := squareEq.SolveReal(a, b, c)

	switch err {
	case nil:
		fmt.Printf(" имеет решения: x₁=%s, x₂=%s\n", fmtnum.F64(x1), fmtnum.F64(x2))
	case squareEq.ErrNoRealRoots:
		fmt.Println(" не имеет вещественных решений")
	case squareEq.ErrZeroA:
		fmt.Println(" не является квадратным")
	}

	var (
		ai = 1 + 0i
		bi = -5 + 6i
		ci = -1 - 9i
	)

	fmt.Printf("Уравнение %sx²+%sx+%s=0", fmtnum.C128(ai), fmtnum.C128(bi), fmtnum.C128(ci))
	xi1, xi2, err := squareEq.SolveComplex(ai, bi, ci)

	switch err {
	case nil:
		fmt.Printf(" имеет решения: x₁=%s, x₂=%s\n", fmtnum.C128(xi1), fmtnum.C128(xi2))
	case squareEq.ErrZeroA:
		fmt.Println(" не является квадратным")
	}
}

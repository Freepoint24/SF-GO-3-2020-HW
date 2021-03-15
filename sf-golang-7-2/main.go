package main

import (
	"fmt"
	"sf-golang-7-2/square-eq"
)

func main() {
	var (
		a float64 = 2
		b float64 = 3
		c float64 = 1
	)

	fmt.Printf("\nУравнение %fx² + %fx + %f = 0\n", a, b, c)
	x1, x2, err := squareEq.SolveReal(a, b, c)

	switch err {
	case nil:
		fmt.Printf("имеет решения: x₁ = %f, x₂ = %f\n", x1, x2)
	case squareEq.ErrNoRealRoots:
		fmt.Println("не имеет вещественных решений")
	case squareEq.ErrZeroA:
		fmt.Println("не является квадратным")
	}

	var (
		ai = 1 + 0i
		bi = -5 + 6i
		ci = -1 - 9i
	)

	fmt.Printf("\nУравнение %fx² + %fx + %f = 0\n", ai, bi, ci)
	xi1, xi2, err := squareEq.SolveComplex(ai, bi, ci)

	switch err {
	case nil:
		fmt.Printf("имеет решения: x₁ = %f, x₂ = %f\n", xi1, xi2)
	case squareEq.ErrZeroA:
		fmt.Println("не является квадратным")
	}
}

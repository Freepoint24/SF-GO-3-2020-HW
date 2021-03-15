package squareEq

import (
	"fmt"
	"math"
	"math/cmplx"
)

// ErrZeroA - ошибка, в случае если a=0 и уравнение не является квадратным
var ErrZeroA = fmt.Errorf("not a quadratic equation")

// ErrNoRealRoots - ошибка, в случае если дискриминант отрицательный
// и уравнение не имеет вещественных решений
var ErrNoRealRoots = fmt.Errorf("equation has no real roots")

// SolveReal - находит вещественные корни квадратного уравнения
//     ax² + bx + c = 0
// или возвращает ошибку, если уравнение с заданными коэффициентами
// не имеет вещественных решений, либо не является квадратным.
// Коэффициенты a, b и c — вещественные числа.
func SolveReal(a, b, c float64) (x1, x2 float64, err error) {

	// Если a = 0, то уравнение не является квадратным
	if a == 0 {
		err = ErrZeroA
		return
	}

	// Вычисление дискриминанта
	D := b*b - 4*a*c

	// Если дискриминант отрицательный,
	// то уравнение не имеет вещественных корней
	if D < 0 {
		err = ErrNoRealRoots
		return
	}

	// Если дискриминант неотрицательный,
	// вычисляем оба корня квадратного уравнения
	DSqrt := math.Sqrt(D)
	x1 = (-b - DSqrt) / (2 * a)
	x2 = (-b + DSqrt) / (2 * a)
	return
}

// SolveComplex - находит комплексные корни квадратного уравнения
//     ax² + bx + c = 0
// или возвращает ошибку, если уравнение не является квадратным.
// Коэффициенты a, b и c — комплексные числа.
func SolveComplex(a, b, c complex128) (x1, x2 complex128, err error) {
	// Если a = 0, то уравнение не является квадратным
	if a == 0+0i {
		err = ErrZeroA
		return
	}

	DSqrt := cmplx.Sqrt(b*b - 4*a*c)
	x1 = (-b - DSqrt) / (2 * a)
	x2 = (-b + DSqrt) / (2 * a)

	return
}

// Package calc реализует структуру и методы калькулятора, который выполняет
// 4 арифметических действия: сложение, вычитание, умножение, деление
package calc

import (
	"fmt"
	"math"
)

// ErrDivZero - ошибка при попытке деления на ноль.
// В качестве результата вычислений Calculate вернет NaN
//
// ErrUnknownOperator - ошибка при попытке использовать неизвестный оператор.
// В качестве результата вычислений Calculate вернет NaN
var (
	ErrDivZero         = fmt.Errorf("попытка деления на ноль")
	ErrUnknownOperator = fmt.Errorf("неизвестный оператор")
)

// Операторы калькулятора
const (
	ADD = "+"
	SUB = "-"
	MUL = "*"
	DIV = "/"
)

/*
todo Вопрос ментору
Почему в этом задании (7.7), а также в задании 7.6.1 используется
не-экспортируемый тип и экспортируемый конструктор?
Какие есть практические применения у такого подхода?

todo Вопрос ментору
В "референсной" реализации калькулятора конструктор возвращает структуру, а не ссылку на структуру:

	type calculator struct{}
	func NewCalculator() calculator {
		return calculator{}
	}

Линтеры подобное объявление подсвечивают "Exported function with the unexported return type",
хотя код прекрасно компилируется и работает.

1. Есть ли практические применения у такого подхода?
2. В чем вообще практическая разница между тем, возвращать в конструкторе ссылку на структуру
   или саму структуру? Когда и в каких ситуациях нужно использовать тот и другой подход?

*/

type calculator struct{}

// NewCalculator - возвращает ссылку на экземпляр калькулятора.
func NewCalculator() *calculator {
	return &calculator{}
}

// Calculate - производит заданное арифметическое действие с операторами.
// Возвращает результат действий, либо NaN и ошибку.
func (c *calculator) Calculate(x, y float64, operator string) (float64, error) {
	switch operator {
	case ADD:
		return c.add(x, y), nil
	case SUB:
		return c.sub(x, y), nil
	case MUL:
		return c.mul(x, y), nil
	case DIV:
		return c.div(x, y)
	default:
		return math.NaN(), ErrUnknownOperator
	}
}

func (c *calculator) add(x, y float64) float64 {
	return x + y
}

func (c *calculator) sub(x, y float64) float64 {
	return x - y
}

func (c *calculator) mul(x, y float64) float64 {
	return x * y
}

func (c *calculator) div(x, y float64) (float64, error) {
	if y == 0 {
		return math.NaN(), ErrDivZero
	}
	return x / y, nil
}

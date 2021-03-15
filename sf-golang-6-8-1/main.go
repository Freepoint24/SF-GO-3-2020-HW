package main

import "fmt"

func main() {
	var a, b float32
	var operator string

	_, err := fmt.Scanln(&a)
	if err != nil {
		panic("Недопустимый операнд")
	}

	_, err = fmt.Scanln(&operator)

	_, err = fmt.Scanln(&b)
	if err != nil {
		panic("Недопустимый операнд")
	}

	switch operator {
	case "+":
		fmt.Printf("Результат сложения: %g\n", a+b)
	case "-":
		fmt.Printf("Результат вычитания: %g\n", a-b)
	case "*":
		fmt.Printf("Результат умножения: %g\n", a*b)
	case "/":
		if b != 0 {
			fmt.Printf("Результат деления: %g\n", a/b)
		} else {
			panic("Деление на ноль недопустимо")
		}
	default:
		panic("Неизвестный оператор")
	}
}

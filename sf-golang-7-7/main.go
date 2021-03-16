package main

import "fmt"
import "sf-golang-7-7/calc"

func userInput(title string, value interface{}) {
	fmt.Printf(title)
	if _, err := fmt.Scanln(value); err != nil {
		panic(err)
	}
}

func main() {
	var x, y float64
	var operator string

	userInput("Введите x: ", &x)
	userInput("Введите оператор: ", &operator)
	userInput("Введите y: ", &y)

	c := calc.NewCalculator()
	result, err := c.Calculate(x, y, operator)
	if err!=nil {
		panic(err)
	}

	fmt.Printf("Результат: x %s y = %f\n", operator, result)
}

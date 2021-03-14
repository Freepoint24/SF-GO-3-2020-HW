package main

import (
	"fmt"
	"os"
)

func main() {
	var str string
	_, err := fmt.Scanln(&str)
	if err != nil {
		panic(err)
	}

	balanced, pos := isBracketsBalanced(str)
	if !balanced {
		fmt.Printf("Строка не сбалансирована: символ %d\n", pos+1)
		os.Exit(1)
	} else {
		fmt.Println("Строка сбалансирована")
	}
}

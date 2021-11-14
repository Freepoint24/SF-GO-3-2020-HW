package main

import (
	"fmt"
	"sf-golang-14-6-1/intersect"
)

func main() {
	l1 := mustReadUint("Enter first array size:\n")
	l2 := mustReadUint("Enter second array size:\n")
	s1 := readStringSlice("Enter first array:\n", l1)
	s2 := readStringSlice("Enter second array:\n", l2)
	fmt.Printf("%v\n", intersect.StringSlices(s1, s2))
}

func mustReadUint(prompt string) uint {
	var value uint
	for {
		fmt.Print(prompt)
		if _, err := fmt.Scan(&value); err == nil {
			break
		}
		var flush string
		_, _ = fmt.Scanln(&flush)
	}
	return value
}

func readStringSlice(prompt string, size uint) []string {
	if size == 0 {
		return []string{}
	}
	s := make([]string, size)
	fmt.Print(prompt)
	for i := uint(0); i < size; i++ {
		_, _ = fmt.Scan(&s[i])
	}
	return s
}

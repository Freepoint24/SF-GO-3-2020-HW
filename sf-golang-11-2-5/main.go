package main

import "fmt"

// Size - размер списка
func Size(list []int) int {
	return len(list)
}

// Add - добавляет элемент в список
func Add(list []int, el int) []int {
	return append(list, el)
}

// Insert - добавляет элемент в список в позицию index
func Insert(list []int, el, index int) []int {
	list = append(list, el)
	for key := Size(list) - 1; key > index; key-- {
		list[key] = list[key-1]
	}
	list[index] = el
	return list
}

// Remove - удаляет элемент из списка в позиции index
func Remove(list []int, index int) []int {
	for key := index + 1; key < Size(list); key++ {
		list[key-1] = list[key]
	}
	return list[:Size(list)-1]
}

func main() {
	list := []int{1, 2, 3}
	fmt.Println(Size(list))

	list = Add(list, 4)
	fmt.Println(list)

	list = Insert(list, 0, 0)
	fmt.Println(list)

	list = Remove(list, 2)
	fmt.Println(list)
}

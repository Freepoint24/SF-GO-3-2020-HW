package sort

import (
	"math/rand"
)

// Quick - быстрая сортировка
//
// Сложность
//    В лучшем случае: O(n log(n))
//    В среднем случае: O(n log(n))
//    В худшем случае: O(n²)
//    Емкостная, в худшем: O(1)
func Quick(list []int) {
	length := len(list)

	if length < 2 {
		return
	}
	if length == 2 {
		if list[0] > list[1] {
			list[0], list[1] = list[1], list[0]
		}
		return
	}

	pivot := rand.Intn(length)
	left, right := 0, length-1
	list[pivot], list[right] = list[right], list[pivot]

	for i := range list {
		if list[i] < list[right] {
			list[i], list[left] = list[left], list[i]
			left++
		}
	}

	list[left], list[right] = list[right], list[left]
	Quick(list[:left])
	Quick(list[left+1:])
}

package sort

// Selection - сортировка выбором
//
// Сложность
//    В лучшем случае: O(n²)
//    В среднем случае: O(n²)
//    В худшем случае: O(n²)
//    Емкостная, в худшем: O(1)
func Selection(list []int) {
	length := len(list)
	for i := 0; i < length-1; i++ {
		foundIndex := i
		for j := i + 1; j < length; j++ {
			if list[foundIndex] > list[j] {
				foundIndex = j
			}
		}
		if foundIndex != i {
			list[i], list[foundIndex] = list[foundIndex], list[i]
		}
	}
}

// SelectionBidi - двунаправленная сортировка выбором
func SelectionBidi(list []int) {
	length := len(list)
	lastIndex := length - 1

	for i := 0; i < length/2; i++ {
		foundLeftIndex, foundRightIndex := i, lastIndex-i

		for j := i + 1; j < length; j++ {
			if list[foundLeftIndex] > list[j] {
				foundLeftIndex = j
			}
			if list[lastIndex-j] > list[foundRightIndex] {
				foundRightIndex = lastIndex - j
			}
		}

		if foundLeftIndex != i {
			list[i], list[foundLeftIndex] = list[foundLeftIndex], list[i]
			if foundRightIndex == i {
				foundRightIndex = foundLeftIndex
			}
		}
		if foundRightIndex != lastIndex-i {
			list[lastIndex-i], list[foundRightIndex] = list[foundRightIndex], list[lastIndex-i]
		}
	}
}

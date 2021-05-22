package sort

// Bubble - сортировка списка пузырьковым алгоритмом
//
// Сложность
//    В лучшем случае: O(n)
//    В среднем случае: O(n²)
//    В худшем случае: O(n²)Е
//    Емкостная, в худшем: O(1)
func Bubble(list []int) {
	length := len(list)
	for i := 1; i < length; i++ {
		isSorted := true
		for j := 0; j < length-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				isSorted = false
			}
		}
		if isSorted {
			return
		}
	}
}

// BubbleRecursive - рекурсивная сортировка списка пузырьковым алгоритмом
func BubbleRecursive(list []int) {
	length := len(list)
	var iteration func(int)
	iteration = func(lastIndex int) {
		if (lastIndex) < 1 {
			return
		}
		for i := 0; i < lastIndex; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
			}
		}
		iteration(lastIndex - 1)
	}
	iteration(length - 1)
}

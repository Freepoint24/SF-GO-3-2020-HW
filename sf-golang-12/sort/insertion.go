package sort

// Insertion - сортировка вставкой
//
// Сложность:
//    В лучшем случае: O(n).
//    В среднем случае: O(n²).
//    В худшем случае: O(n²).
//    Емкостная, в худшем: O(1).
func Insertion(list []int) {
	length := len(list)
	for i := 0; i < length-1; i++ {
		if list[i] > list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
			for j := i; j > 0; j-- {
				if list[j-1] > list[j] {
					list[j-1], list[j] = list[j], list[j-1]
				}
			}
		}
	}
}

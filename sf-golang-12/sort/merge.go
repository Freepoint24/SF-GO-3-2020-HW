package sort

// Merge - сортировка слиянием
//
// Сложность
//    В лучшем случае: O(n log(n))
//    В среднем случае: O(n log(n))
//    В худшем случае: O(n log(n))
//    Емкостная, в худшем: O(n)
func Merge(list []int) {
	length := len(list)
	switch {
	case length < 2:
		return
	case length == 2:
		if list[0] > list[1] {
			list[0], list[1] = list[1], list[0]
		}
		return
	case length > 2:
		middle := length / 2
		a := make([]int, middle)
		copy(a, list[:middle])
		b := make([]int, length-middle)
		copy(b, list[middle:])
		Merge(a)
		Merge(b)
		cursorA, cursorB := 0, 0
		lenA, lenB := len(a), len(b)
		for i := 0; !(cursorA >= lenA && cursorB >= lenB); i++ {
			if cursorB == lenB {
				list[i] = a[cursorA]
				cursorA++
				continue
			}
			if cursorA == lenA {
				list[i] = b[cursorB]
				cursorB++
				continue
			}
			if a[cursorA] <= b[cursorB] {
				list[i] = a[cursorA]
				cursorA++
			} else {
				list[i] = b[cursorB]
				cursorB++
			}
		}
	}
}

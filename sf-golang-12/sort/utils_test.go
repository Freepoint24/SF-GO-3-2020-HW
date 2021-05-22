package sort

import "fmt"

func Example_clusterSortedSlice()  {
	fmt.Printf("%v", clusterSortedSlice(5,5))
	// Output: [20 21 22 23 24 15 16 17 18 19 10 11 12 13 14 5 6 7 8 9 0 1 2 3 4]
}

func Example_sortedReversedSlice()  {
	fmt.Printf("%v", sortedReversedSlice(50,10))
	// Output: [59 58 57 56 55 54 53 52 51 50]
}
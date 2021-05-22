package sort

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomSlice(max, size int) []int {
	list := make([]int, size)
	for i := range list {
		list[i] = rand.Intn(max*2) - max
	}
	return list
}

func sortedSlice(start, size int) []int {
	list := make([]int, size)
	for i := range list {
		list[i] = i + start
	}
	return list
}

func sortedReversedSlice(start, size int) []int {
	list := make([]int, size)
	for i := range list {
		list[size-i-1] = i + start
	}
	return list
}

func clusterSortedSlice(clusters, clusterSize int) []int {
	size := clusters * clusterSize
	list := make([]int, size)
	i := 0
	for c := clusters-1; c >= 0; c-- {
		for s := 0; s < clusterSize; s++ {
			list[i] = c * clusterSize + s
			i++
		}
	}
	return list
}

func isSorted(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			return false
		}
	}
	return true
}

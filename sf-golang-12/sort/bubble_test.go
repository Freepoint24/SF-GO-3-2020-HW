package sort

import (
	"testing"
)

func TestBubble(t *testing.T) {
	var l = randomSlice(100, 100)
	Bubble(l)
	if !isSorted(l) {
		t.Error("список не отсортирован")
	}
}

func TestBubbleRecursive(t *testing.T) {
	var l = randomSlice(100, 100)
	BubbleRecursive(l)
	if !isSorted(l) {
		t.Error("список не отсортирован")
	}
}

func BenchmarkBubble(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := randomSlice(10, 10)
			b.StartTimer()
			Bubble(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := randomSlice(100, 1000)
			b.StartTimer()
			Bubble(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := randomSlice(10000, 100000)
			b.StartTimer()
			Bubble(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})
}

//func BenchmarkBubbleRecursive(b *testing.B) {
//	b.Run("small arrays", func(b *testing.B) {
//		b.StopTimer()
//		for i := 0; i < b.N; i++ {
//			l := randomSlice(10, 10)
//			b.StartTimer()
//			BubbleRecursive(l)
//			b.StopTimer()
//		}
//		b.ReportAllocs()
//	})
//
//	b.Run("middle arrays", func(b *testing.B) {
//		b.StopTimer()
//		for i := 0; i < b.N; i++ {
//			l := randomSlice(100, 1000)
//			b.StartTimer()
//			BubbleRecursive(l)
//			b.StopTimer()
//		}
//		b.ReportAllocs()
//	})
//
//	b.Run("big arrays", func(b *testing.B) {
//		b.StopTimer()
//		for i := 0; i < b.N; i++ {
//			l := randomSlice(10000, 100000)
//			b.StartTimer()
//			BubbleRecursive(l)
//			b.StopTimer()
//		}
//		b.ReportAllocs()
//	})
//}

func BenchmarkBubbleSorted(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedSlice(0, 10)
			b.StartTimer()
			Bubble(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedSlice(0, 1000)
			b.StartTimer()
			Bubble(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedSlice(0, 100000)
			b.StartTimer()
			Bubble(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})
}

func BenchmarkBubbleSortedReversed(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedReversedSlice(10, 10)
			b.StartTimer()
			Bubble(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedReversedSlice(1000, 1000)
			b.StartTimer()
			Bubble(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedReversedSlice(100000, 100000)
			b.StartTimer()
			Bubble(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})
}

func BenchmarkBubbleClusterSorted(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := clusterSortedSlice(2, 5)
			b.StartTimer()
			Bubble(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := clusterSortedSlice(12, 84)
			b.StartTimer()
			Bubble(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := clusterSortedSlice(117, 855)
			b.StartTimer()
			Bubble(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})
}

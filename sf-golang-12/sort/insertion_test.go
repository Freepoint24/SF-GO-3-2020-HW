package sort

import (
	"testing"
)

func TestInsertion(t *testing.T) {
	var l = randomSlice(100, 100)
	Insertion(l)
	if !isSorted(l) {
		t.Error("список не отсортирован")
	}
}

func BenchmarkInsertion(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := randomSlice(10, 10)
			b.StartTimer()
			Insertion(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := randomSlice(100, 1000)
			b.StartTimer()
			Insertion(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := randomSlice(10000, 100000)
			b.StartTimer()
			Insertion(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})
}

func BenchmarkInsertionSorted(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedSlice(0, 10)
			b.StartTimer()
			Insertion(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedSlice(0, 1000)
			b.StartTimer()
			Insertion(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedSlice(0, 100000)
			b.StartTimer()
			Insertion(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})
}

func BenchmarkInsertionSortedReversed(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedReversedSlice(10, 10)
			b.StartTimer()
			Insertion(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedReversedSlice(1000, 1000)
			b.StartTimer()
			Insertion(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedReversedSlice(100000, 100000)
			b.StartTimer()
			Insertion(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})
}

func BenchmarkInsertionClusterSorted(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := clusterSortedSlice(2, 5)
			b.StartTimer()
			Insertion(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := clusterSortedSlice(12, 84)
			b.StartTimer()
			Insertion(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := clusterSortedSlice(117, 855)
			b.StartTimer()
			Insertion(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})
}

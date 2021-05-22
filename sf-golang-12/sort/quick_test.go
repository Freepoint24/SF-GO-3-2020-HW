package sort

import (
	"testing"
)

func TestQuick(t *testing.T) {
	l := randomSlice(50, 100)
	Quick(l)
	if !isSorted(l) {
		t.Error("список не отсортирован")
	}
}

func BenchmarkQuick(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := randomSlice(10, 10)
			b.StartTimer()
			Quick(l)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := randomSlice(100, 1000)
			b.StartTimer()
			Quick(l)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := randomSlice(10000, 100000)
			b.StartTimer()
			Quick(l)
			b.StopTimer()
		}
	})
}

func BenchmarkQuickSorted(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedSlice(0, 10)
			b.StartTimer()
			Quick(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedSlice(0, 1000)
			b.StartTimer()
			Quick(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedSlice(0, 100000)
			b.StartTimer()
			Quick(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})
}

func BenchmarkQuickSortedReversed(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedReversedSlice(10, 10)
			b.StartTimer()
			Quick(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedReversedSlice(1000, 1000)
			b.StartTimer()
			Quick(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedReversedSlice(100000, 100000)
			b.StartTimer()
			Quick(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})
}

func BenchmarkQuickClusterSorted(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := clusterSortedSlice(2, 5)
			b.StartTimer()
			Quick(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := clusterSortedSlice(12, 84)
			b.StartTimer()
			Quick(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := clusterSortedSlice(117, 855)
			b.StartTimer()
			Quick(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})
}

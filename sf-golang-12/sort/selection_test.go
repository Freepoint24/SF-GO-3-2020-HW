package sort

import (
	"testing"
)

func TestSelection(t *testing.T) {
	var l = randomSlice(50, 100)
	Selection(l)
	if !isSorted(l) {
		t.Error("список не отсортирован")
	}
}

func TestSelectionBidi(t *testing.T) {
	var l = randomSlice(50, 100)
	SelectionBidi(l)
	if !isSorted(l) {
		t.Errorf("список не отсортирован по возрастанию")
	}
}

func BenchmarkSelection(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := randomSlice(10, 10)
			b.StartTimer()
			Selection(l)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := randomSlice(100, 1000)
			b.StartTimer()
			Selection(l)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := randomSlice(10000, 100000)
			b.StartTimer()
			Selection(l)
			b.StopTimer()
		}
	})

}

func BenchmarkSelectionBidi(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := randomSlice(10, 10)
			b.StartTimer()
			SelectionBidi(l)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := randomSlice(100, 1000)
			b.StartTimer()
			SelectionBidi(l)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := randomSlice(10000, 100000)
			b.StartTimer()
			SelectionBidi(l)
			b.StopTimer()
		}
	})
}

func BenchmarkSelectionSorted(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedSlice(0, 10)
			b.StartTimer()
			Selection(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedSlice(0, 1000)
			b.StartTimer()
			Selection(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedSlice(0, 100000)
			b.StartTimer()
			Selection(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})
}

func BenchmarkSelectionSortedReversed(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedReversedSlice(10, 10)
			b.StartTimer()
			Selection(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedReversedSlice(1000, 1000)
			b.StartTimer()
			Selection(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedReversedSlice(100000, 100000)
			b.StartTimer()
			Selection(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})
}

func BenchmarkSelectionClusterSorted(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := clusterSortedSlice(2, 5)
			b.StartTimer()
			Selection(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := clusterSortedSlice(12, 84)
			b.StartTimer()
			Selection(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := clusterSortedSlice(117, 855)
			b.StartTimer()
			Selection(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})
}

func BenchmarkSelectionBidiSorted(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedSlice(0, 10)
			b.StartTimer()
			SelectionBidi(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedSlice(0, 1000)
			b.StartTimer()
			SelectionBidi(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedSlice(0, 100000)
			b.StartTimer()
			SelectionBidi(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})
}

func BenchmarkSelectionBidiSortedReversed(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedReversedSlice(10, 10)
			b.StartTimer()
			SelectionBidi(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedReversedSlice(1000, 1000)
			b.StartTimer()
			SelectionBidi(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := sortedReversedSlice(100000, 100000)
			b.StartTimer()
			SelectionBidi(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})
}

func BenchmarkSelectionBidiClusterSorted(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := clusterSortedSlice(2, 5)
			b.StartTimer()
			SelectionBidi(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := clusterSortedSlice(12, 84)
			b.StartTimer()
			SelectionBidi(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			l := clusterSortedSlice(117, 855)
			b.StartTimer()
			SelectionBidi(l)
			b.StopTimer()
		}
		b.ReportAllocs()
	})
}

package pipe

import (
	"testing"
	"time"
)

func TestRingBuffer(t *testing.T) {
	out := make(chan int)
	start, finish := make(chan struct{}), make(chan struct{})
	v := 0
	p := NewPipe(
		RingBuffer(1, 500*time.Millisecond),
		ToChan(out),
	)

	p.Emit(1, 2)
	close(start)

	go func() {
		defer close(finish)
		<-start
		for {
			select {
			case v = <-out:
				if v != 2 {
					t.Errorf("got %v, want %v", v, 2)
				}
				return
			case <-time.After(550 * time.Millisecond):
				t.Errorf("timeout")
				return
			}
		}
	}()

	<-finish
}

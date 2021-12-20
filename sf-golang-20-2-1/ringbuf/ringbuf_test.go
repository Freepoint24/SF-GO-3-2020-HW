package ringbuf

import (
	"testing"
)

func TestRingBuf_trivial(t *testing.T) {
	r := NewRingBuf(3)
	r.Write(1)
	r.Write(2)
	r.Write(3)

	got, ok := r.Read()
	if got != 1 || !ok {
		t.Errorf("Read() = %v, %v, want %v, %v", got, ok, 1, true)
	}

	got, ok = r.Read()
	if got != 2 || !ok {
		t.Errorf("Read() = %v, %v, want %v, %v", got, ok, 2, true)
	}

	got, ok = r.Read()
	if got != 3 || !ok {
		t.Errorf("Read() = %v, %v, want %v, %v", got, ok, 3, true)
	}
}

func TestRingBuf_overflow(t *testing.T) {
	r := NewRingBuf(2)
	r.Write(1)
	r.Write(2)
	r.Write(3)

	got, ok := r.Read()
	if got != 2 || !ok {
		t.Errorf("Read() = %v, %v, want %v, %v", got, ok, 2, true)
	}

	got, ok = r.Read()
	if got != 3 || !ok {
		t.Errorf("Read() = %v, %v, want %v, %v", got, ok, 3, true)
	}

	got, ok = r.Read()
	if ok {
		t.Errorf("Read() = %v, %v, want %v, %v", got, ok, 0, false)
	}

	r.Write(10)
	r.Write(20)

	got, ok = r.Read()
	if got != 10 || !ok {
		t.Errorf("Read() = %v, %v, want %v, %v", got, ok, 10, true)
	}

	got, ok = r.Read()
	if got != 20 || !ok {

		got, ok = r.Read()
		t.Errorf("Read() = %v, %v, want %v, %v", got, ok, 20, true)
	}

}

func TestRingBuf_single(t *testing.T) {
	r := NewRingBuf(1)
	r.Write(1)
	r.Write(2)

	got, ok := r.Read()
	if got != 2 || !ok {
		t.Errorf("Read() = %v, %v, want %v, %v", got, ok, 2, true)
	}

	got, ok = r.Read()
	if ok {
		t.Errorf("Read() = %v, %v, want %v, %v", got, ok, 0, false)
	}

	r.Write(1)

	got, ok = r.Read()
	if got != 1 || !ok {
		t.Errorf("Read() = %v, %v, want %v, %v", got, ok, 1, true)
	}
}

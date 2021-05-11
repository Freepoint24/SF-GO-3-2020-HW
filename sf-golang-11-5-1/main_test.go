package main

import (
	"testing"
)

func TestIntRing_Read(t *testing.T) {
	var length = 5
	ring, err := NewIntRing(length, 0)
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < length+1; i++ {
		// записываем в буфер
		if err = ring.Write(42); err != nil {
			t.Error(err)
		}
		// считываем из буфера
		if _, err = ring.Read(); err != nil {
			t.Error(err)
		}
	}
}

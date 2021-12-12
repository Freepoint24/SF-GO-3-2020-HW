package semaphore

import (
	"errors"
	"time"
)

var (
	ErrAcquire = errors.New("не удалось захватить семафор")
	ErrRelease = errors.New("не удалось освободить семафор")
)

// Semaphore - реализация семафора подсчета,
// с начальным количеством занятых ресурсов и таймаутом
type Semaphore struct {
	c       chan struct{}
	timeout time.Duration
}

// NewSemaphore - инициализирует новый экземпляр Semaphore,
// с начальным количеством занятых ресурсов initialCount,
// максимальным количеством ресурсов maxCount и таймаутом.
// Вызывает панику при недопустимых initialCount и maxCount.
func NewSemaphore(initialCount, maxCount int, timeout time.Duration) *Semaphore {
	// Проверяем initialCount и maxCount на граничные значения
	switch {
	case initialCount < 0:
		panic("initialCount не может быть отрицательным")
	case maxCount < 1:
		panic("maxCount должен быть положительным")
	case maxCount < initialCount:
		panic("maxCount не может быть меньше initialCount")
	}

	s := Semaphore{
		c:       make(chan struct{}, maxCount),
		timeout: timeout,
	}

	// Занимаем начальное кол-во ресурсов
	for i := 0; i < initialCount; i++ {
		s.c <- struct{}{}
	}

	return &s
}

// Acquire - захват семафора. Возвращает ошибку ErrAcquire, если превышен таймаут
func (s *Semaphore) Acquire() error {
	select {
	case s.c <- struct{}{}:
		return nil
	case <-time.After(s.timeout):
		return ErrAcquire
	}
}

// Release - освобождение семафора. Возвращает ошибку ErrRelease, если превышен таймаут
func (s *Semaphore) Release() error {
	select {
	case <-s.c:
		return nil
	case <-time.After(s.timeout):
		return ErrRelease
	}
}

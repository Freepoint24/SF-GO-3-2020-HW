package main

import (
	"fmt"
	"sync"
)

const (
	step            int64 = 1    // Шаг наращивания счётчика
	endCounterValue int64 = 1000 // Конечное значение счетчика
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	var counter int64 = 0

	// Горутина-воркер
	increment := func() {
		c.L.Lock()
		defer c.L.Unlock()
		counter += step
		c.Signal()
	}

	// Счетчик итераций. Предполагаем, что endCounterValue кратно step
	var iterationCount = endCounterValue / step

	// Запуск воркеров
	for i := int64(0); i < iterationCount; i++ {
		go increment()
	}

	// Ожидаем завершения воркеров
	c.L.Lock()
	defer c.L.Unlock()
	for counter < endCounterValue {
		c.Wait()
	}

	// Печатаем результат
	fmt.Println(counter)
}

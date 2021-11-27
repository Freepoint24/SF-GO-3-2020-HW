package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	step            int64 = 1    // Шаг наращивания счётчика
	endCounterValue int64 = 1000 // Конечное значение счетчика
	workersNum            = 10   // Кол-во горутин-воркеров
)

func main() {
	var wg sync.WaitGroup

	// Результат
	var counter int64 = 0

	// Счетчик необходимых итераций. Предполагаем, что endCounterValue кратно step
	var iterationCount = endCounterValue / step

	// Горутина-воркер
	increment := func() {
		defer wg.Done()
		for {
			// Уменьшаем счетчик итераций на 1
			iterationRemain := atomic.AddInt64(&iterationCount, -1)
			// Если итераций не осталось, то завершаем воркер
			if iterationRemain < 0 {
				return
			}
			// Иначе увеличиваем значение счетчика и продолжаем цикл
			atomic.AddInt64(&counter, step)
		}
	}

	// Запускаем необходимое кол-во воркеров
	wg.Add(workersNum)
	for i := 0; i < workersNum; i++ {
		go increment()
	}

	// Ожидаем завершения работы воркеров
	wg.Wait()
	// Печатаем результат
	fmt.Println(counter)
}

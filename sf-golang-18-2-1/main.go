package main

import (
	"fmt"
	"sync"
)

const (
	workersN    = 5  // Количество запускаемых горутин
	iterationsN = 10 // Количество итераций каждой горутины
)

// worker - Горутина печатает свой идентификатор заданное кол-во раз
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= iterationsN; i++ {
		fmt.Printf("Worker #%d, iteration %d\n", id, i)
	}
}

func main() {
	var wg sync.WaitGroup
	// Добавляем заданное кол-во горутин в вейтгруппу
	wg.Add(workersN)
	// Запускаем горутины
	for id := 0; id < workersN; id++ {
		go worker(id, &wg)
	}
	// Ожидаем завершения всех запущенных горутин
	wg.Wait()
}

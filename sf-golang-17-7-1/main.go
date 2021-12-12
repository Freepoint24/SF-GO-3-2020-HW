package main

import (
	"fmt"
	"sf-golang-17-7-1/counter"
	"sync"
)

const (
	limit   = 1234567 // Лимит счетчика
	workers = 123     // Кол-во запускаемых горутин-воркеров
)

// worker - горутина-воркер. Итеративно увеличивает значение счетчика на 1,
// до тех пор, пока счетчик не превысит заданный лимит.
// После этого выводит на экран отчет о кол-ве совершенных итераций и завершается.
func worker(id int, c *counter.Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; ; i++ {
		if ok := c.Add(1); !ok {
			fmt.Println("Worker:", id, "Iterations:", i)
			break
		}
	}
}

func main() {
	var wg sync.WaitGroup
	// Создаем счетчик с заданным лимитом
	c := counter.NewCounter(limit)

	// Запускаем требуемое кол-во горутин-воркеров и ожидаем их завершения
	wg.Add(workers)
	for id := 0; id < workers; id++ {
		go worker(id, c, &wg)
	}
	wg.Wait()

	// Печатаем значение счетчика
	fmt.Println("Counter:", c.Value())
}

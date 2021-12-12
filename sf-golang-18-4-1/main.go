package main

import (
	"fmt"
	"sf-golang-18-4-1/semaphore"
	"sync"
	"time"
)

// 5 воркеров должны отработать 10 запросов,
// но одновременно могут обрабатывать запрос только 3 воркера
const (
	maxWorkers     = 5  // Максимальное кол-во горутин-воркеров
	requestsNumber = 10 // Кол-во запросов, которое необходимо обработать воркерам
	maxResources   = 3  // Максимальное кол-во ресурсов, которое могут использовать воркеры
)

// Данные запроса
type data struct{}

// worker - горутина принимает запросы из канала и обрабатывает их.
// Семафор ограничивает кол-во одновременно обрабатываемых запросов
func worker(id int, c <-chan data, s *semaphore.Semaphore, wg *sync.WaitGroup) {
	defer wg.Done()

	log := func(msg string) {
		fmt.Printf("#%d %s\n", id, msg)
	}

	log("Запущен")

	// Получаем запросы
	for range c {
		log("Получил запрос")

		// Захватываем ресурс
		if err := s.Acquire(); err != nil {
			log(err.Error())
		} else {
			log("Захватил ресурс")
		}

		// Выполняем работу
		// NB: одновременно в этом блоке кода может находиться
		// не более maxResources воркеров
		time.Sleep(time.Second)
		log("Обработал запрос")

		// Освобождаем ресурс
		if err := s.Release(); err != nil {
			log(err.Error())
		} else {
			log("Освободил ресурс")
		}
	}

	log("Остановлен")
}

func main() {
	// Канал, по которому будут поступать запросы к воркерам
	c := make(chan data, requestsNumber)

	// Семафор, ограничивающий кол-во одновременного использования ресурсов воркерами
	var s = semaphore.NewSemaphore(0, maxResources, time.Second*10)

	// Запускаем пул воркеров
	var wg sync.WaitGroup
	wg.Add(maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		go worker(i, c, s, &wg)
	}

	// Заполняем канал запросами
	for i := 0; i < requestsNumber; i++ {
		c <- data{}
	}
	close(c)

	// Ожидаем завершения воркеров
	wg.Wait()
}

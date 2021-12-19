package main

import (
	"fmt"
	"sync"
)

// demultiplexingFunc - функция разуплотнения каналов.
// Принимает канал-источник и необходимое кол-во каналов-потребителей.
// Возвращает слайс каналов-потребителей и канал закрытия done.
func demultiplexingFunc(dataSourceChan <-chan int, amount int) ([]chan int, <-chan struct{}) {
	// Инициализируем слайс каналов потребителей
	var output = make([]chan int, amount)
	for i := range output {
		output[i] = make(chan int)
	}

	// Инициализируем канал закрытия
	done := make(chan struct{})

	// Запускаем горутину, которая при получении сообщения из канала источника,
	// пересылает его во все каналы-потребители
	go func() {
		// При поступлении сообщения в канал-источник
		// отправляем его в каждый из каналов-потребителей.
		for v := range dataSourceChan {
			for _, c := range output {
				c <- v
			}
		}
		// После завершения посылки сообщений в основной
		// канал-источник данных, закрываем канал закрытия
		close(done)
	}()

	// Возвращаем инициализированные слайс каналов-потребителей
	// и канал закрытия
	return output, done
}

// multiplexingFunc - функция уплотнения каналов.
// Принимает канал закрытие done и произвольное кол-во каналов-источников.
// Возвращает канал-приемник.
func multiplexingFunc(done <-chan struct{}, channels ...chan int) <-chan int {
	// Инициализируем канал-приемник, в который будут попадать сообщения от всех источников.
	multiplexedChan := make(chan int)

	// Для каждого канала источника запускаем отдельную горутину,
	// которая будет принимать из него сообщения и отправлять их в общий канал.
	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, c := range channels {
		go func(c <-chan int) {
			defer wg.Done()
			for {
				select {
				// Если поступило сообщение из канала-источника, отправляем его в канал-приемник
				case v := <-c:
					multiplexedChan <- v
				// Если поступило сообщение в done, завершаем горутину
				case <-done:
					return
				}
			}
		}(c)
	}

	// Запускаем горутину, которая закроет канал после того,
	// как в закрывающий канал поступит сигнал о прекращении работы каналов-источников
	go func() {
		wg.Wait()
		close(multiplexedChan)
	}()
	return multiplexedChan
}

func main() {
	// Горутина-источник данных создаёт свой собственный канал
	// и посылает в него 5 сообщений
	startDataSource := func() <-chan int {
		c := make(chan int)
		go func() {
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 1; i <= 5; i++ {
					c <- i
				}
			}()
			wg.Wait()
			close(c)
		}()
		return c
	}

	// Запускаем источник данных и демультиплексируем его на 3 канала
	var consumers, done = demultiplexingFunc(startDataSource(), 3)
	// Мультиплексируем каналы в один
	c := multiplexingFunc(done, consumers...)

	// Получаем сообщения из мультиплексированного канала и печатаем их
	for data := range c {
		fmt.Println(data)
	}
}

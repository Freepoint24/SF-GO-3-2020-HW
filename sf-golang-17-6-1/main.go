package main

import (
	"fmt"
	"time"
)

func main() {
	// Продюсер сообщений с нормальным приоритетом
	normalC := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			normalC <- i
		}
	}()

	// Продюсер сообщений с высоким приоритетом
	priorityC := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			priorityC <- i
		}
	}()

	// Счетчики сообщений, прочитанных из каналов
	var normalCounter, priorityCounter = 0, 0

	// Читаем первые 1000 сообщений из каналов
	for normalCounter+priorityCounter < 1000 {
		select {
		case <-priorityC:
			priorityCounter++
		case <-normalC:
			// Увеличиваем вероятность получения из приоритетного канала в 2 раза
			select {
			case <-priorityC:
				priorityCounter++
			default:
				normalCounter++
			}
		default:
			// Эмулятор не блокирующейся при чтении из каналов "полезной нагрузки"
			time.Sleep(time.Millisecond)
		}
	}

	fmt.Println("Normal messages:", normalCounter)
	fmt.Println("Priority messages:", priorityCounter)
}

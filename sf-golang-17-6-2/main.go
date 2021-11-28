package main

import (
	"fmt"
	"time"
)

func main() {
	// Продюсер канала c1
	c1 := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c1 <- i
		}
		close(c1)
	}()

	// Продюсер канала c2
	c2 := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c2 <- i
		}
		close(c2)
	}()

	// Признаки закрытия каналов
	var c1Closed, c2Closed bool

	// Читаем сообщения из каналов
	for {
		select {
		// Если в канале есть сообщение и он не закрыт, печатаем id канала и сообщение из него
		case value, ok := <-c1:
			if ok {
				fmt.Println("c1:", value)
			} else {
				c1Closed = true
			}
		case value, ok := <-c2:
			if ok {
				fmt.Println("c2:", value)
			} else {
				c2Closed = true
			}
		// Если ни в одном канале нет сообщений, печатаем текущее время
		default:
			fmt.Println(time.Now().Format("15:04:05"))
		}
		// Если оба канала закрыты, завершаем цикл
		if c1Closed && c2Closed {
			break
		}
	}

}

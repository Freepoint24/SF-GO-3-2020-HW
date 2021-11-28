package main

import (
	"fmt"
	"sync"
)

// producer - отправляет в канал последовательные числа [1..n].
func producer(n int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 1; i <= n; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

// consumer - читает числа из канала и выводит их на экран
func consumer(c <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		n, ok := <-c
		if ok {
			fmt.Println(n)
		} else {
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	c := producer(100)
	wg.Add(1)
	go consumer(c, &wg)
	wg.Wait()
}

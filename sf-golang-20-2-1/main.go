package main

import (
	"os"
	"sf-golang-20-2-1/pipe"
	"time"
)

func main() {
	// Включаем вывод отладочных сообщений (по-умолчанию в stderr)
	pipe.Debug = true

	// Создаем пайплайн
	p := pipe.NewPipe(
		pipe.PassMin(0),                   // Фильтруем отрицательные значения
		pipe.PassDivBy(3),                 // Фильтруем значения не кратные 3
		pipe.RingBuffer(2, time.Second*5), // Кольцевой буфер на 2 элемента с таймаутом 5 сек
		pipe.ToWriter(os.Stdout),          // Выводим в stdout
	)

	// Читаем из stdin и ожидаем завершения чтения
	<-p.EmitFromReader(os.Stdin)

	// Тк мы больше ничего не собираемся передавать в пайплайн, то закрываем его
	p.Close()

	// Дожидаемся завершения всех обработчиков пайплайна
	<-p.Done()
}

package counter

// Counter - реализация потокобезопасного счетчика с лимитом
type Counter struct {
	value  int       // Значение счетчика
	limit  int       // Лимит счетчика
	txDone chan bool // Буферизованный канал синхронизации транзакций
}

// NewCounter - создает новый счетчик
func NewCounter(limit int) *Counter {
	c := Counter{
		limit: limit,
		value: 0,
		// Создаем буферизованный канал на 1 элемент
		txDone: make(chan bool, 1),
	}
	// Делаем счетчик доступным для транзакции
	c.txDone <- true
	return &c
}

// Add - потокобезопасное добавление значения к счетчику. Возвращает true в случае успеха.
// Если значение счетчика больше или равно лимиту,
// то счетчик не увеличивается, а функция возвращает false
func (c *Counter) Add(amount int) bool {
	// Ожидаем доступности транзакции
	<-c.txDone

	// Если значение больше или равно лимита,
	// завершаем транзакцию и возвращаем false
	if c.value >= c.limit {
		c.txDone <- true
		return false
	}

	// Иначе увеличиваем счетчик и завершаем транзакцию
	c.value += amount
	c.txDone <- true
	return true
}

// Value - потокобезопасно возвращает значение счетчика
func (c *Counter) Value() int {
	// Ожидаем доступности транзакции
	<-c.txDone
	// Читаем значение счетчика, завершаем транзакцию и возвращаем считанное значение
	v := c.value
	c.txDone <- true
	return v
}

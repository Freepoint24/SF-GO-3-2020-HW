package ringbuf

import "sync"

// RingBuf - потокобезопасная реализация кольцевого буфера.
// При переполнении буфера самое старое значение затирается новым,
// а чтение из буфера начинается со следующего элемента.
type RingBuf struct {
	data  []int      // Данные буфера
	size  int        // Размер буфера
	count int        // Счетчик элементов в буфере
	rPos  int        // Позиция чтения из буфера
	wPos  int        // Позиция записи в буфер
	m     sync.Mutex // Мютекс для потокобезопасной работы
}

// NewRingBuf - создает новый буфер заданного размера
func NewRingBuf(size int) *RingBuf {
	// Минимальный размер буфера - 1 элемент
	if size < 1 {
		size = 1
	}
	return &RingBuf{
		data:  make([]int, size),
		size:  size,
		count: 0,
		rPos:  0,
		wPos:  0,
		m:     sync.Mutex{},
	}
}

// Read - читает элемент из буфера.
// Если буфер пуст, возвращает вторым значением false.
func (r *RingBuf) Read() (int, bool) {
	r.m.Lock()
	defer r.m.Unlock()
	// Если в буфере нет элементов, возвращаем false
	if r.count == 0 {
		return 0, false
	}
	v := r.data[r.rPos]
	// Уменьшаем счетчик элементов в буфере
	r.count--
	// Перемещаем позицию чтения
	r.rPos = (r.rPos + 1) % r.size
	return v, true
}

// Write - записывает элемент в буфер.
// При переполнении буфера самое старое значение затирается новым,
// а чтение из буфера начинается со следующего элемента.
func (r *RingBuf) Write(v int) {
	r.m.Lock()
	defer r.m.Unlock()
	r.data[r.wPos] = v
	// Перемещаем позицию записи вперед
	r.wPos = (r.wPos + 1) % r.size

	if r.count < r.size {
		// Если буфер не заполнен, увеличиваем счетчик элементов
		r.count++
	} else {
		// Иначе перемещаем позицию чтения
		r.rPos = (r.rPos + 1) % r.size
	}
}

// Count - возвращает кол-во элементов в буфере
func (r *RingBuf) Count() int {
	r.m.Lock()
	defer r.m.Unlock()
	return r.count
}

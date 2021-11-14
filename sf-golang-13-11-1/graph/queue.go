package graph

// queue - односвязная очередь вершин для поиска по графу без весов
type queue struct {
	head   *queueItem      // Первый элемент очереди
	tail   *queueItem      // Последний элемент
	queued map[string]bool // Маркеры ранее добавленных в очередь вершин
}

// queueItem - элемент очередеи queue
type queueItem struct {
	node *node
	next *queueItem
}

func newQueue() *queue {
	return &queue{queued: make(map[string]bool)}
}

// put - добавление вершины в конец очереди
func (q *queue) put(node *node) {
	if node == nil {
		return
	}
	// Если узел с такми ключем ранее добавлялся в очередь, то не добавляем его
	if q.queued[node.key] {
		return
	}

	q.queued[node.key] = true
	item := &queueItem{node: node}

	if q.head == nil {
		q.head = item
	}

	if q.tail != nil {
		q.tail.next = item
	}
	q.tail = item
}

// get - извлечение вершины из начала очереди
func (q *queue) get() *node {
	// Если очередь пустая
	if q.head == nil {
		return nil
	}

	result := q.head.node
	q.head = q.head.next

	// Если извлечен последний элемент
	if q.head == nil {
		q.tail = nil
	}

	return result
}

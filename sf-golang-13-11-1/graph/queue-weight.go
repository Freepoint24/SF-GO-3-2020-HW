package graph

// queue - односвязная очередь вершин для поиска по графу с весами
type queueWeight struct {
	head   *queueWeightItem // Первый элемент очереди
	tail   *queueWeightItem // Последний элемент
	queued map[string]bool  // Маркеры ранее добавленных в очередь вершин
}

// queueWeightItem - элемент очередеи queueWeight
type queueWeightItem struct {
	node *nodeWeight
	next *queueWeightItem
}

func newQueueWeight() *queueWeight {
	return &queueWeight{queued: make(map[string]bool)}
}

// put - добавление вершины в конец очереди
func (q *queueWeight) put(node *nodeWeight) {
	if node == nil {
		return
	}
	// Если узел с такми ключем ранее добавлялся в очередь, то не добавляем его
	if q.queued[node.key] {
		return
	}

	q.queued[node.key] = true
	item := &queueWeightItem{node: node}

	if q.head == nil {
		q.head = item
	}

	if q.tail != nil {
		q.tail.next = item
	}
	q.tail = item
}

// get - извлечение вершины из начала очереди
func (q *queueWeight) get() *nodeWeight {
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

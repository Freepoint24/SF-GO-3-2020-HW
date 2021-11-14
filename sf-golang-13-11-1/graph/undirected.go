package graph

import (
	"errors"
	"fmt"
)

// node - вершина графа без весов
type node struct {
	key   string
	edges map[string]*node // Ребра вершины
}

// Undirected - неориентированный граф без весов
type Undirected struct {
	nodes map[string]*node // Вершины графа
}

func NewUndirected() *Undirected {
	return &Undirected{nodes: make(map[string]*node)}
}

// Add - добавляет вершину в граф
func (u *Undirected) Add(key string) {
	// Если такая вершина уже существует, ничего не делаем
	if _, exists := u.nodes[key]; exists {
		return
	}
	u.nodes[key] = &node{key: key, edges: make(map[string]*node)}
}

// Connect - связывает 2 вершины графа
// Возвращает ошибку, если вершина отсуствует в графе
func (u *Undirected) Connect(key1, key2 string) error {
	node1 := u.nodes[key1]
	node2 := u.nodes[key2]

	if node1 == nil || node2 == nil {
		return errors.New("неизвестная вершина графа")
	}

	node1.edges[key2] = node2
	node2.edges[key1] = node1
	return nil
}

// BFS - поиск по графу в ширину начиная со заданной начальной вершины.
// Возвращает ошибку, если начальная вершина отсуствует в графе
func (u *Undirected) BFS(startKey, searchKey string) error {
	startNode := u.nodes[startKey]
	if startNode == nil {
		return errors.New("неизвестная вершина графа")
	}

	distance := 0
	q := newQueue()
	q.put(startNode)

	for {
		n := q.get()
		if n == nil {
			fmt.Printf("Путь от `%s` до `%s` не найден\n", startKey, searchKey)
			return nil
		}
		if n.key == searchKey {
			fmt.Println("=>", n.key)
			return nil
		}
		fmt.Println("...", n.key)

		distance++
		for _, near := range n.edges {
			q.put(near)
		}
	}
}
